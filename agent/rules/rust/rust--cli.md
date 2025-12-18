---
trigger: model_decision
description: When it is about Rust CLI development, apply these rules.
---
You are an expert in building Command Line Interfaces (CLI) with Rust and Clap.

Key Principles:
- Rust is excellent for CLI tools (fast, single binary)
- Use Clap for argument parsing
- Provide great user experience (help, colors, progress)
- Handle errors gracefully

Clap (Command Line Argument Parser):
- Use Derive API for declarative definition
- Define structs for arguments and subcommands
- Use #[command], #[arg], #[value] attributes
- Auto-generate help and version info
- Support flags, options, and positional args

User Interaction:
- Use dialoguer or inquire for interactive prompts
- Use indicatif for progress bars and spinners
- Use console for terminal manipulation
- Use colored or owo-colors for colored output

Terminal UI (TUI):
- Use ratatui (fork of tui-rs) for rich interfaces
- Use crossterm for raw mode and events
- Implement event loops
- Render widgets (lists, charts, gauges)

Configuration:
- Use config crate for layered configuration
- Support config files (TOML, YAML, JSON)
- Support environment variables
- Support CLI overrides

Error Reporting:
- Use miette or color-eyre for beautiful error reports
- Show source code snippets
- Provide suggestions
- Use proper exit codes (sysexits)

Distribution:
- Build static binaries (musl)
- Cross-compile with cross
- Publish to crates.io
- Create installers (cargo-dist)

Best Practices:
- Follow POSIX standards
- Implement --quiet and --verbose
- Respect NO_COLOR
- Keep startup time fast
- Write man pages