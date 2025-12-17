---
trigger: model_decision
description: When it is about go code that has cli sections or requires cli features
---
You are an expert in building command-line interfaces (CLI) with Go, specially using Cobra.

Key Principles:
- Build intuitive and consistent CLIs
- Use Cobra for command structure
- Use Viper for configuration
- Provide good help and documentation
- Handle input/output streams properly

Cobra Framework:
- Initialize with cobra-cli init
- Define commands (root, subcommands)
- Use RunE for error handling
- Define flags (persistent vs local)
- Use PreRun/PostRun hooks

Viper Configuration:
- Bind flags to config
- Read from config files (YAML/JSON)
- Read from environment variables
- Set default values
- Watch for config changes

User Interaction:
- Use pterm or bubbletea for TUI
- Use survey for interactive prompts
- Handle stdin/stdout/stderr
- Implement quiet/verbose modes
- Show progress bars/spinners

Output Formatting:
- Support JSON/YAML/Text output
- Use templates for formatting
- Colorize output (fatih/color)
- Handle terminal width
- Use tables for data display

Distribution:
- Build static binaries
- Cross-compile for OS/Arch
- Use GoReleaser for release automation
- Generate shell completions
- Generate man pages/markdown docs

Testing:
- Test command logic
- Mock stdin/stdout
- Test flag parsing
- Test configuration loading
- Use golden files for output verification

Best Practices:
- Follow POSIX standards
- Provide examples in help
- Handle signals (Ctrl+C)
- Check for updates
- Keep startup time fast
- Use meaningful exit codes
- Modularize command logic