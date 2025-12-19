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
	agentComment = "# Agent rules/workflows downloaded by agent-cli (do not commit)"
)

// EnsureAgentIgnored ensures that the specified directory is in the .gitignore file.
// If the .gitignore file doesn't exist, it creates one.
// If the entry already exists, it does nothing.
func EnsureAgentIgnored(targetDir, ignoredDir string) error {
	gitignorePath := filepath.Join(targetDir, ".gitignore")

	// Check if .gitignore exists
	_, err := os.Stat(gitignorePath)
	if os.IsNotExist(err) {
		return createGitignoreWithEntry(gitignorePath, ignoredDir)
	}
	if err != nil {
		return fmt.Errorf("failed to check .gitignore: %w", err)
	}

	// Check if entry already exists
	exists, err := entryExists(gitignorePath, ignoredDir)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	return appendEntry(gitignorePath, ignoredDir)
}

// createGitignoreWithEntry creates a new .gitignore file with the agent entry.
func createGitignoreWithEntry(path, ignoredDir string) error {
	entry := ignoredDir
	if !strings.HasSuffix(entry, "/") {
		entry += "/"
	}
	content := fmt.Sprintf("%s\n%s\n", agentComment, entry)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	fmt.Printf("Created .gitignore with %s entry\n", entry)
	return nil
}

// entryExists checks if the agent entry already exists in .gitignore.
func entryExists(path, ignoredDir string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("failed to open .gitignore: %w", err)
	}
	defer file.Close()

	entryWithSlash := ignoredDir
	if !strings.HasSuffix(entryWithSlash, "/") {
		entryWithSlash += "/"
	}
	entryWithoutSlash := strings.TrimSuffix(ignoredDir, "/")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == entryWithSlash || line == entryWithoutSlash {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// appendEntry appends the agent entry to an existing .gitignore file.
func appendEntry(path, ignoredDir string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore for appending: %w", err)
	}
	defer file.Close()

	entry := ignoredDir
	if !strings.HasSuffix(entry, "/") {
		entry += "/"
	}
	content := fmt.Sprintf("\n%s\n%s\n", agentComment, entry)
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to append to .gitignore: %w", err)
	}

	fmt.Printf("Added %s entry to .gitignore\n", entry)
	return nil
}
