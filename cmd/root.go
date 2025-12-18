// Package cmd contains the CLI commands for agent-cli.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// Repo is the GitHub repository to fetch rules from (format: owner/repo)
	Repo string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agent-cli",
	Short: "A CLI tool to manage agent rules",
	Long: `agent-cli downloads agent rule files from a GitHub repository
and places them in the local .agent/rules directory.

Use 'agent-cli list' to see available agent types.
Use 'agent-cli get --type <type>' to download rules for a specific agent type.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&Repo, "repo", "odinnordico/antigravity-agents",
		"GitHub repository to fetch rules from (format: owner/repo)")
}
