// Package workflows provides functionality to manage agent workflows downloaded from GitHub.
package workflows

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/odinnordico/antigravity-agents/internal/github"
)

const workflowsPath = "agent/workflows"

// Manager handles downloading and managing agent workflows.
type Manager struct {
	client    *github.Client
	targetDir string
}

// NewManager creates a new workflows manager.
func NewManager(client *github.Client, targetDir string) *Manager {
	return &Manager{
		client:    client,
		targetDir: targetDir,
	}
}

// ListTypes returns the available workflow types from the repository.
func (m *Manager) ListTypes() ([]string, error) {
	contents, err := m.client.ListContents(workflowsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list workflow types: %w", err)
	}

	var types []string
	for _, item := range contents {
		if item.Type == "dir" {
			types = append(types, item.Name)
		}
	}

	return types, nil
}

// Download downloads all workflow files for the specified workflow type.
func (m *Manager) Download(workflowType string) error {
	sourcePath := filepath.Join(workflowsPath, workflowType)
	contents, err := m.client.ListContents(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to list workflows for type '%s': %w", workflowType, err)
	}

	// Ensure target directory exists
	targetPath := filepath.Join(m.targetDir, ".agent", "workflows")
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

		fmt.Printf("  Downloaded: %s\n", item.Name)
	}

	return nil
}
