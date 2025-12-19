package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/odinnordico/antigravity-agents/internal/github"
	"github.com/odinnordico/antigravity-agents/internal/gitignore"
	"github.com/odinnordico/antigravity-agents/internal/rules"
	"github.com/odinnordico/antigravity-agents/internal/workflows"
	"github.com/spf13/cobra"
)

var (
	selectedTypes string
	getRules      bool
	getWorkflows  bool
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download rules and/or workflows for specified types",
	Long: `Downloads rule and/or workflow files for the specified type(s) from the
GitHub repository and places them in the local .agent/rules and .agent/workflows directories.

Multiple types can be specified as comma-separated lists.

Examples:
  agent-cli get --type go,python
  agent-cli get --workflow git,testing
  agent-cli get --type go --workflow git`,
	Run: func(cmd *cobra.Command, args []string) {
		if !getRules && !getWorkflows {
			fmt.Fprintln(os.Stderr, "Error: at least one of --rules or --workflow is required")
			os.Exit(1)
		}

		if selectedTypes == "" {
			fmt.Fprintln(os.Stderr, "Error: --type flag is required to specify which items to download")
			os.Exit(1)
		}

		// Get current working directory as target
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		client := github.NewClient(Repo)
		rulesManager := rules.NewManager(client, cwd)
		workflowsManager := workflows.NewManager(client, cwd)

		requestedTypes := parseTypes(selectedTypes)

		// Parse and validate agent types (rules)
		var requestedAgentTypes []string
		if getRules {
			requestedAgentTypes = requestedTypes
			availableAgentTypes, err := rulesManager.ListTypes()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching available agent types: %v\n", err)
				os.Exit(1)
			}
			invalidTypes := validateTypes(requestedAgentTypes, availableAgentTypes)
			if len(invalidTypes) > 0 {
				fmt.Fprintf(os.Stderr, "Error: invalid agent type(s): %s\n", strings.Join(invalidTypes, ", "))
				fmt.Fprintf(os.Stderr, "Available types: %s\n", strings.Join(availableAgentTypes, ", "))
				os.Exit(1)
			}
		}

		// Parse and validate workflow types
		var requestedWorkflowTypes []string
		if getWorkflows {
			requestedWorkflowTypes = requestedTypes
			availableWorkflowTypes, err := workflowsManager.ListTypes()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching available workflow types: %v\n", err)
				os.Exit(1)
			}
			invalidTypes := validateTypes(requestedWorkflowTypes, availableWorkflowTypes)
			if len(invalidTypes) > 0 {
				fmt.Fprintf(os.Stderr, "Error: invalid workflow type(s): %s\n", strings.Join(invalidTypes, ", "))
				fmt.Fprintf(os.Stderr, "Available workflows: %s\n", strings.Join(availableWorkflowTypes, ", "))
				os.Exit(1)
			}
		}

		fmt.Printf("From repository: %s\n\n", Repo)

		// Download rules for each agent type
		if len(requestedAgentTypes) > 0 {
			fmt.Printf("Downloading rules for: %s\n", strings.Join(requestedAgentTypes, ", "))
			for _, agentType := range requestedAgentTypes {
				fmt.Printf("\n[Rules: %s]\n", agentType)
				if err := rulesManager.Download(agentType); err != nil {
					fmt.Fprintf(os.Stderr, "Error downloading rules '%s': %v\n", agentType, err)
					os.Exit(1)
				}
			}
		}

		// Download workflows for each workflow type
		if len(requestedWorkflowTypes) > 0 {
			fmt.Printf("\nDownloading workflows for: %s\n", strings.Join(requestedWorkflowTypes, ", "))
			for _, workflowType := range requestedWorkflowTypes {
				fmt.Printf("\n[Workflows: %s]\n", workflowType)
				if err := workflowsManager.Download(workflowType); err != nil {
					fmt.Fprintf(os.Stderr, "Error downloading workflows '%s': %v\n", workflowType, err)
					os.Exit(1)
				}
			}
		}

		// Ensure .agent/ is in .gitignore
		if err := gitignore.EnsureAgentIgnored(cwd); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to update .gitignore: %v\n", err)
		}

		fmt.Println("\nDownload completed successfully!")
	},
}

// parseTypes splits a comma-separated string into a slice of trimmed, non-empty strings.
func parseTypes(input string) []string {
	var result []string
	for _, t := range strings.Split(input, ",") {
		trimmed := strings.TrimSpace(t)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// validateTypes checks if all requested types exist in the available types.
// Returns a slice of invalid types.
func validateTypes(requested, available []string) []string {
	availableSet := make(map[string]bool)
	for _, t := range available {
		availableSet[t] = true
	}

	var invalid []string
	for _, t := range requested {
		if !availableSet[t] {
			invalid = append(invalid, t)
		}
	}
	return invalid
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&selectedTypes, "type", "t", "", "Type(s) to download, comma-separated")
	getCmd.Flags().BoolVarP(&getWorkflows, "workflows", "w", false, "Download workflows")
	getCmd.Flags().BoolVarP(&getRules, "rules", "u", false, "Download agent rules")
}
