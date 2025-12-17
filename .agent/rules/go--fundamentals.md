---
trigger: model_decision
description: When it is about go code
---
You are an expert in Go (Golang) programming and best practices.

Key Principles:
- Follow idiomatic Go (Effective Go)
- Keep it simple and readable
- Handle errors explicitly
- Prefer composition over inheritance
- Use goroutines for concurrency

Code Organization:
- Use standard project layout (cmd/, pkg/, internal/)
- Group related code in packages
- Keep packages small and focused
- Use meaningful package names
- Avoid circular dependencies

Naming Conventions:
- Use CamelCase for exported names
- Use camelCase for unexported names
- Keep names short and concise
- Use single-letter names for short loops/scopes
- Avoid stuttering (e.g., user.UserInfo -> user.Info)

Error Handling:
- Check errors immediately after function calls
- Return errors as the last return value
- Use custom error types for specific cases
- Wrap errors with context (fmt.Errorf("%w"))
- Don't panic unless truly unrecoverable

Functions and Methods:
- Keep functions short and focused
- Use named return values sparingly
- Use defer for cleanup
- Use interfaces for flexibility
- Accept interfaces, return structs

Data Structures:
- Use slices over arrays
- Use maps for key-value storage
- Use structs for grouping data
- Use pointers for large structs or mutability
- Initialize structs with field names

Concurrency:
- Use goroutines for concurrent tasks
- Use channels for communication
- Use sync.WaitGroup to wait for goroutines
- Use sync.Mutex for shared state
- Avoid sharing memory by communicating

Testing:
- Write unit tests in _test.go files
- Use the testing package
- Use table-driven tests
- Run tests with go test
- Use go test -race to check for race conditions

Dependency Management:
- Use Go Modules (go.mod)
- Keep dependencies minimal
- Vendor dependencies if necessary
- Use semantic versioning
- Audit dependencies regularly

Formatting and Linting:
- Always run gofmt
- Use go vet to catch common errors
- Use golangci-lint for comprehensive linting
- Follow community style guides
- Document exported names with comments

Best Practices:
- Handle all errors
- Avoid global state
- Use context for cancellation and timeouts
- Write benchmarks for performance-critical code
- Keep main() simple
- Use standard library when possible