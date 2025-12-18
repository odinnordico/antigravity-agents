package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/odinnordico/antigravity-agents/internal/github"
	"github.com/odinnordico/antigravity-agents/internal/gitignore"
	"github.com/odinnordico/antigravity-agents/internal/rules"
	"github.com/spf13/cobra"
)

var agentTypes string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download rules for one or more agent types",
	Long: `Downloads all rule files for the specified agent type(s) from the
GitHub repository and places them in the local .agent/rules directory.

Multiple types can be specified as a comma-separated list.
Example: agent-cli get --type go,python,agentic`,
	Run: func(cmd *cobra.Command, args []string) {
		if agentTypes == "" {
			fmt.Fprintln(os.Stderr, "Error: --type flag is required")
			os.Exit(1)
		}

		// Parse comma-separated types
		requestedTypes := parseTypes(agentTypes)
		if len(requestedTypes) == 0 {
			fmt.Fprintln(os.Stderr, "Error: no valid types specified")
			os.Exit(1)
		}

		// Get current working directory as target
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		client := github.NewClient(Repo)
		manager := rules.NewManager(client, cwd)

		// Validate all types before downloading
		availableTypes, err := manager.ListAgentTypes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching available types: %v\n", err)
			os.Exit(1)
		}

		invalidTypes := validateTypes(requestedTypes, availableTypes)
		if len(invalidTypes) > 0 {
			fmt.Fprintf(os.Stderr, "Error: invalid agent type(s): %s\n", strings.Join(invalidTypes, ", "))
			fmt.Fprintf(os.Stderr, "Available types: %s\n", strings.Join(availableTypes, ", "))
			os.Exit(1)
		}

		fmt.Printf("Downloading rules for agent type(s): %s\n", strings.Join(requestedTypes, ", "))
		fmt.Printf("From repository: %s\n\n", Repo)

		// Download rules for each type
		for _, agentType := range requestedTypes {
			fmt.Printf("--- Downloading: %s ---\n", agentType)
			if err := manager.DownloadRules(agentType); err != nil {
				fmt.Fprintf(os.Stderr, "Error downloading '%s': %v\n", agentType, err)
				os.Exit(1)
			}
		}

		// Ensure .agent/rules is in .gitignore
		if err := gitignore.EnsureAgentRulesIgnored(cwd); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to update .gitignore: %v\n", err)
		}

		fmt.Println("\nRules downloaded successfully!")
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
	getCmd.Flags().StringVarP(&agentTypes, "type", "t", "", "Agent type(s) to download rules for, comma-separated (required)")
	getCmd.MarkFlagRequired("type")
}
