# Contributing to Antigravity Agent CLI 🤝

Thank you for your interest in contributing to the Antigravity Agent CLI! We welcome all contributions, from bug fixes to new features and documentation improvements.

## Code of Conduct

By participating in this project, you are expected to uphold our [Code of Conduct](CODE_OF_CONDUCT.md).

## How Can I Contribute?

### Reporting Bugs
If you find a bug, please search the [Existing Issues](https://github.com/odinnordico/antigravity-agents/issues) to see if it has already been reported. If not, open a new issue and include:
- A clear, descriptive title.
- Steps to reproduce the bug.
- Actual vs. expected behavior.
- Your OS and Go version.

### Suggesting Enhancements
We Love new ideas! If you have a suggestion:
1. Check if the feature has already been requested.
2. Open an issue with the "enhancement" label.
3. Describe the problem you're trying to solve and how this feature addresses it.

### Pull Requests
Ready to contribute code?
1. Fork the repository.
2. Create a new branch (`feature/my-new-feature` or `fix/issue-id`).
3. Write your code and ensure it follows the project's style.
4. Run tests (if applicable) and verify your changes.
5. Submit a pull request with a clear description of your changes.

## Development Setup

### Prerequisites
- [Go](https://golang.org/dl/) (version 1.21+)

### Building Locally
```bash
git clone https://github.com/your-username/antigravity-agents.git
cd antigravity-agents
go build -o agent-cli .
```

### Project Structure
- `cmd/`: CLI commands (Cobra)
- `internal/`: Core logic and third-party integrations
- `agent/`: Local rule sets and workflows

## Style Guide
- Follow standard Go idioms and `gofmt`.
- Document new functions and exported types.
- Be concise and clear in commit messages.

## Questions?
If you have any questions, feel free to open a discussion or an issue!
