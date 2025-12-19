package cmd

import (
	"fmt"
	"os"

	"github.com/odinnordico/antigravity-agents/internal/github"
	"github.com/odinnordico/antigravity-agents/internal/rules"
	"github.com/odinnordico/antigravity-agents/internal/workflows"
	"github.com/spf13/cobra"
)

var (
	listRules     bool
	listWorkflows bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available types",
	Long: `Lists available agent types (rules) and/or workflow types from the GitHub repository.

By default, lists both rules and workflows.
Use --rules or --workflows to list only one category.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := github.NewClient(Repo)

		// If neither flag is set, show both
		showBoth := !listRules && !listWorkflows

		if listRules || showBoth {
			rulesManager := rules.NewManager(client, ".")
			agentTypes, err := rulesManager.ListTypes()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error listing agent types: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Available agent types (rules):")
			if len(agentTypes) == 0 {
				fmt.Println("  (none found)")
			} else {
				for _, t := range agentTypes {
					fmt.Printf("  - %s\n", t)
				}
			}

			if showBoth {
				fmt.Println()
			}
		}

		if listWorkflows || showBoth {
			workflowsManager := workflows.NewManager(client, ".")
			workflowTypes, err := workflowsManager.ListTypes()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error listing workflow types: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Available workflow types:")
			if len(workflowTypes) == 0 {
				fmt.Println("  (none found)")
			} else {
				for _, t := range workflowTypes {
					fmt.Printf("  - %s\n", t)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listRules, "rules", "u", false, "List only agent types (rules)")
	listCmd.Flags().BoolVarP(&listWorkflows, "workflows", "w", false, "List only workflow types")
}
