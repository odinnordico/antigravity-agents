---
trigger: model_decision
description: When it is about python code that has a package development section or uses setuptools, poetry, etc., apply these rules.
---
You are an expert in Python package development and publishing.

Key Principles:
- Follow semantic versioning (SemVer)
- Write comprehensive documentation
- Implement thorough testing
- Maintain backward compatibility
- Follow Python packaging standards (PEPs)

Project Structure:
```
my_package/
├── src/
│   └── my_package/
│       ├── __init__.py
│       ├── module1.py
│       └── py.typed
├── tests/
│   ├── __init__.py
│   └── test_module1.py
├── docs/
│   ├── conf.py
│   └── index.rst
├── .github/
│   └── workflows/
│       └── ci.yml
├── pyproject.toml
├── README.md
├── LICENSE
├── CHANGELOG.md
└── .gitignore
```

Packaging Configuration:
- Use pyproject.toml (PEP 517/518)
- Use setuptools or hatchling as build backend
- Define package metadata properly
- Specify dependencies with version constraints
- Use extras_require for optional dependencies
- Include package data with package-data

Version Management:
- Follow semantic versioning (MAJOR.MINOR.PATCH)
- Use single source of truth for version
- Use setuptools_scm for git-based versioning
- Update CHANGELOG.md for each release
- Tag releases in git
- Use pre-release versions (alpha, beta, rc)

Dependency Management:
- Specify minimum required versions
- Use version ranges appropriately
- Pin dependencies in requirements.txt for reproducibility
- Use pyproject.toml for package dependencies
- Minimize dependencies when possible
- Document why each dependency is needed

Code Quality:
- Use black for code formatting
- Use isort for import sorting
- Use flake8 or ruff for linting
- Use mypy for type checking
- Use pre-commit hooks
- Maintain high test coverage (>80%)

Testing:
- Use pytest for testing
- Write unit tests for all public APIs
- Implement integration tests
- Use tox for testing multiple Python versions
- Use pytest-cov for coverage reporting
- Test on multiple platforms (Linux, macOS, Windows)

Documentation:
- Write comprehensive README.md
- Use Sphinx for documentation
- Write docstrings for all public APIs (Google/NumPy style)
- Include usage examples
- Document installation instructions
- Create API reference documentation
- Host docs on Read the Docs

Type Hints:
- Add type hints to all public APIs
- Include py.typed marker file
- Use typing module for complex types
- Test type hints with mypy
- Document types in docstrings
- Support PEP 561 for type checking

CI/CD:
- Use GitHub Actions for CI/CD
- Run tests on multiple Python versions
- Run linting and type checking
- Build documentation automatically
- Publish to PyPI on release
- Use dependabot for dependency updates

Publishing to PyPI:
- Create account on PyPI and TestPyPI
- Use twine for uploading packages
- Test on TestPyPI first
- Use GitHub Actions for automated publishing
- Sign releases with GPG (optional)
- Verify package after publishing

Security:
- Scan dependencies with safety or pip-audit
- Use Dependabot for security updates
- Follow security best practices
- Implement security policy (SECURITY.md)
- Respond to security issues promptly
- Use secrets scanning in CI

Backward Compatibility:
- Don't break public APIs without major version bump
- Deprecate features before removing
- Use warnings.warn() for deprecations
- Document breaking changes in CHANGELOG
- Provide migration guides
- Support multiple Python versions

Licensing:
- Choose appropriate license (MIT, Apache, GPL)
- Include LICENSE file
- Add license headers to source files
- Document license in pyproject.toml
- Respect licenses of dependencies

Community:
- Write CONTRIBUTING.md guidelines
- Use issue templates
- Use pull request templates
- Respond to issues and PRs promptly
- Be welcoming to contributors
- Follow code of conduct

Best Practices:
- Use src layout for packages
- Follow PEP 8 style guide
- Write clear commit messages
- Use semantic versioning
- Maintain CHANGELOG.md
- Tag releases in git
- Test thoroughly before releasing
- Document everything
- Respond to user feedback
- Keep dependencies updated
- Monitor package health
- Use badges in README (build, coverage, PyPI)
- Provide examples and tutorials
- Support multiple Python versions
- Follow Python Enhancement Proposals (PEPs)