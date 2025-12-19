# Antigravity Agent CLI 🚀

`agent-cli` is a powerful, extensible command-line tool designed to manage agentic rules and workflows. It allows developers to quickly download curated rule sets and automation workflows directly from the [antigravity-agents](https://github.com/odinnordico/antigravity-agents) ecosystem into their local projects.

> [!NOTE]
> This project is inspired by **Google's Antigravity Agent-First IDE** ([antigravity.codes](https://antigravity.codes/)) and its [predefined rules](https://antigravity.codes/rules). We aim to bring similar agentic capabilities to CLI workflows.

## Features

- 📂 **Multi-Category Downloads**: Separately manage **Rules** (best practices, coding standards) and **Workflows** (automation steps, CI/CD).
- 🔗 **Fork-Friendly**: Easily specify a custom repository (fork) using the `--repo` flag.
- ✅ **Built-in Validation**: Ensures requested types exist in the repository before downloading.
- 🚀 **Built with Go & Cobra**: Fast, efficient, and provides a polished CLI experience with shorthands and bash completion.
- 🛡️ **Auto-Gitignore**: Automatically ensures downloaded rules/workflows are ignored by git to keep your project clean.

---

## Installation

### From Source
Requires Go 1.21 or higher.

```bash
git clone https://github.com/odinnordico/antigravity-agents.git
cd antigravity-agents
go build -o agent-cli .
```

---

## Usage Guide

The CLI uses a flexible flag-based system. You must specify which category you want to fetch (`--rules` or `--workflows`) and which specific types (`--type`).

### 1. Listing Available Content
Explore what's available in the repository:

```bash
# List everything (rules and workflows)
./agent-cli list

# List only agent rules
./agent-cli list --rules

# List only workflows
./agent-cli list --workflows
```

### 2. Downloading Content
Use the `get` command to pull files into your project.

```bash
# Download agent rules for Go
./agent-cli get --rules --type go

# Download workflows for Git
./agent-cli get --workflows --type git

# Download multiple types at once (comma-separated)
./agent-cli get --rules --type go,python,javascript

# Download both rules and workflows for a type
./agent-cli get --rules --workflows --type testing

# Download to a custom directory
./agent-cli get --rules --type go --output my-rules
```

### 3. Advanced Usage
Specify a custom repository (e.g., your own fork):

```bash
./agent-cli get --rules --type go --repo my-username/antigravity-agents-fork
```

---

## Architecture Overview

The project is structured following clean Go patterns:

- `cmd/`: CLI command definitions using **Cobra**.
- `internal/github/`: API client for interacting with GitHub's REST API.
- `internal/rules/`: Logic for managing agent rule downloads.
- `internal/workflows/`: Logic for managing workflow downloads.
- `internal/gitignore/`: Utility to manage `.gitignore` file state.

Downloaded files are placed in:
- `.agent/rules/`
- `.agent/workflows/`

---

## Contributing

We welcome contributions! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to submit issues and pull requests.

## License

This project is licensed under the **GNU General Public License v3.0**. See the [LICENSE](LICENSE) file for details.
