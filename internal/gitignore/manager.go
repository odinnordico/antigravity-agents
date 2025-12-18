// Package gitignore provides functionality to manage .gitignore entries.
package gitignore

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	agentEntry        = ".agent"
	agentRulesComment = "# Agent rules/workflows downloaded by agent-cli (do not commit)"
)

// EnsureAgentRulesIgnored ensures that .agent/rules/ is in the .gitignore file.
// If the .gitignore file doesn't exist, it creates one.
// If the entry already exists, it does nothing.
func EnsureAgentRulesIgnored(targetDir string) error {
	gitignorePath := filepath.Join(targetDir, ".gitignore")

	// Check if .gitignore exists
	_, err := os.Stat(gitignorePath)
	if os.IsNotExist(err) {
		return createGitignoreWithEntry(gitignorePath)
	}
	if err != nil {
		return fmt.Errorf("failed to check .gitignore: %w", err)
	}

	// Check if entry already exists
	exists, err := entryExists(gitignorePath)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	return appendEntry(gitignorePath)
}

// createGitignoreWithEntry creates a new .gitignore file with the agent rules entry.
func createGitignoreWithEntry(path string) error {
	content := fmt.Sprintf("%s\n%s\n", agentRulesComment, agentEntry)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	fmt.Println("Created .gitignore with agent rules entry")
	return nil
}

// entryExists checks if the agent rules entry already exists in .gitignore.
func entryExists(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("failed to open .gitignore: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == agentEntry || line == ".agent/rules" {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// appendEntry appends the agent rules entry to an existing .gitignore file.
func appendEntry(path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for appending: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf("\n%s\n%s\n", agentRulesComment, agentEntry)
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to append to .gitignore: %w", err)
	}

	fmt.Println("Added agent rules entry to .gitignore")
	return nil
}
