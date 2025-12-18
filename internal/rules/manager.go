// Package rules provides functionality to manage agent rules downloaded from GitHub.
package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/odinnordico/antigravity-agents/internal/github"
)

const rulesPath = "agent/rules"

// Manager handles downloading and managing agent rules.
type Manager struct {
	client    *github.Client
	targetDir string
}

// NewManager creates a new rules manager.
func NewManager(client *github.Client, targetDir string) *Manager {
	return &Manager{
		client:    client,
		targetDir: targetDir,
	}
}

// ListAgentTypes returns the available agent types from the repository.
func (m *Manager) ListAgentTypes() ([]string, error) {
	contents, err := m.client.ListContents(rulesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list agent types: %w", err)
	}

	var types []string
	for _, item := range contents {
		if item.Type == "dir" {
			types = append(types, item.Name)
		}
	}

	return types, nil
}

// DownloadRules downloads all rule files for the specified agent type.
func (m *Manager) DownloadRules(agentType string) error {
	path := filepath.Join(rulesPath, agentType)
	contents, err := m.client.ListContents(path)
	if err != nil {
		return fmt.Errorf("failed to list rules for type '%s': %w", agentType, err)
	}

	// Ensure target directory exists
	targetPath := filepath.Join(m.targetDir, ".agent", "rules")
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	for _, item := range contents {
		if item.Type != "file" {
			continue
		}

		data, err := m.client.DownloadFile(item.DownloadURL)
		if err != nil {
			return fmt.Errorf("failed to download '%s': %w", item.Name, err)
		}

		filePath := filepath.Join(targetPath, item.Name)
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return fmt.Errorf("failed to write '%s': %w", item.Name, err)
		}

		fmt.Printf("Downloaded: %s\n", item.Name)
	}

	return nil
}
