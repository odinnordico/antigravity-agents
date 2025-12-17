---
trigger: model_decision
description: When it is about python code that has a CLI section or uses click, argparse, typer, etc., apply these rules.
---
You are an expert in Python command-line interface (CLI) development.

Key Principles:
- Make CLIs intuitive and user-friendly
- Follow Unix philosophy (do one thing well)
- Provide helpful error messages
- Support both interactive and non-interactive modes
- Follow standard CLI conventions

CLI Frameworks:
- Use Click for complex CLIs (recommended)
- Use argparse for simple CLIs (stdlib)
- Use Typer for modern type-hinted CLIs
- Use docopt for declarative CLIs
- Use Fire for automatic CLI generation

Click Fundamentals:
- Use @click.command() decorator
- Use @click.option() for options
- Use @click.argument() for positional arguments
- Implement command groups with @click.group()
- Use click.echo() instead of print()
- Implement proper help text

Argument and Option Design:
- Use short (-v) and long (--verbose) forms
- Provide sensible defaults
- Use type validation (int, float, Path)
- Implement required vs optional arguments
- Use click.Choice() for enumerated values
- Support multiple values with multiple=True

User Input:
- Use click.prompt() for interactive input
- Use click.confirm() for yes/no questions
- Use click.password_prompt() for sensitive input
- Implement input validation
- Provide default values
- Support stdin for piping

Output Formatting:
- Use click.echo() for output
- Implement --quiet and --verbose flags
- Support multiple output formats (JSON, CSV, table)
- Use click.style() for colored output
- Implement progress bars with click.progressbar()
- Use rich library for beautiful terminal output

Error Handling:
- Use click.ClickException for user errors
- Provide helpful error messages
- Exit with appropriate status codes (0 for success)
- Log errors to stderr
- Implement --debug flag for detailed errors

Configuration:
- Support config files (YAML, TOML, INI)
- Use environment variables for defaults
- Implement --config option
- Follow XDG Base Directory specification
- Support both global and local configs

Subcommands:
- Organize related commands in groups
- Use click.group() for command groups
- Implement help for each subcommand
- Support command aliases
- Implement nested command groups

File Handling:
- Use click.File() for file arguments
- Support stdin/stdout with '-'
- Use pathlib.Path for file paths
- Implement file validation
- Support glob patterns

Progress and Feedback:
- Use click.progressbar() for long operations
- Implement spinner for indeterminate progress
- Use rich.progress for advanced progress bars
- Provide status messages during execution
- Implement --quiet mode to suppress output

Testing:
- Use click.testing.CliRunner for testing
- Test all commands and options
- Test error conditions
- Test with different input combinations
- Mock external dependencies

Documentation:
- Write comprehensive help text
- Use docstrings for commands
- Generate man pages
- Create README with usage examples
- Use --help for inline documentation

Packaging and Distribution:
- Use setuptools with entry_points
- Create pyproject.toml for modern packaging
- Publish to PyPI
- Support pip install
- Include shell completion scripts

Shell Completion:
- Implement tab completion for commands
- Support bash, zsh, fish completions
- Use click.shell_completion
- Generate completion scripts
- Document completion installation

Interactive Mode:
- Implement REPL with click.prompt()
- Use cmd module for advanced REPL
- Support command history
- Implement auto-completion in REPL
- Provide interactive help

Best Practices:
- Follow POSIX conventions
- Use --version flag
- Implement --help for all commands
- Exit with 0 on success, non-zero on error
- Support --verbose and --quiet
- Use environment variables for configuration
- Implement proper signal handling
- Support piping and redirection
- Make CLIs composable (Unix philosophy)
- Provide clear, actionable error messages
- Use colors sparingly and make them optional
- Test on multiple platforms
- Document all options and arguments
- Version your CLI properly
- Maintain backward compatibility