package cmd

import (
	"fmt"
	"os"

	"github.com/odinnordico/antigravity-agents/internal/github"
	"github.com/odinnordico/antigravity-agents/internal/gitignore"
	"github.com/odinnordico/antigravity-agents/internal/rules"
	"github.com/spf13/cobra"
)

var agentType string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download rules for an agent type",
	Long: `Downloads all rule files for the specified agent type from the
GitHub repository and places them in the local .agent/rules directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if agentType == "" {
			fmt.Fprintln(os.Stderr, "Error: --type flag is required")
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

		fmt.Printf("Downloading rules for agent type: %s\n", agentType)
		fmt.Printf("From repository: %s\n\n", Repo)

		if err := manager.DownloadRules(agentType); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Ensure .agent/rules is in .gitignore
		if err := gitignore.EnsureAgentRulesIgnored(cwd); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to update .gitignore: %v\n", err)
		}

		fmt.Println("\nRules downloaded successfully!")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&agentType, "type", "t", "", "Agent type to download rules for (required)")
	getCmd.MarkFlagRequired("type")
}
