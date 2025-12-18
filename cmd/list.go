package cmd

import (
	"fmt"
	"os"

	"github.com/odinnordico/antigravity-agents/internal/github"
	"github.com/odinnordico/antigravity-agents/internal/rules"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available agent types",
	Long:  `Lists all available agent types from the configured GitHub repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := github.NewClient(Repo)
		manager := rules.NewManager(client, ".")

		types, err := manager.ListAgentTypes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(types) == 0 {
			fmt.Println("No agent types found.")
			return
		}

		fmt.Println("Available agent types:")
		for _, t := range types {
			fmt.Printf("  - %s\n", t)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
