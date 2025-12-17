---
trigger: model_decision
description: When it is about go code, apply these rules.
---
You are an expert in Go testing, benchmarking, and quality assurance.

Key Principles:
- Test behavior, not implementation
- Keep tests fast and reliable
- Aim for high test coverage
- Write benchmarks for critical paths
- Use table-driven tests

Unit Testing:
- Use the standard 'testing' package
- Name files *_test.go
- Name functions Test*(t *testing.T)
- Use t.Run for subtests
- Use t.Helper() for helper functions

Table-Driven Tests:
- Define a struct for test cases
- Iterate over test cases
- Run subtests for each case
- Cover edge cases and errors
- Keep test data readable

Mocking:
- Use interfaces for dependencies
- Generate mocks with mockery or gomock
- Implement manual mocks for simple cases
- Mock external services (HTTP, DB)
- Verify calls and expectations

Integration Testing:
- Test interactions between components
- Use build tags (// +build integration)
- Spin up dependencies (Docker containers)
- Test with real database
- Clean up resources after tests

Benchmarking:
- Name functions Benchmark*(b *testing.B)
- Use b.N for iterations
- Reset timer with b.ResetTimer()
- Run with go test -bench=.
- Analyze allocations with -benchmem

Fuzz Testing:
- Use Go 1.18+ fuzzing
- Name functions Fuzz*(f *testing.F)
- Add seed corpus
- Check for crashes and hangs
- Validate invariants

Test Coverage:
- Run with go test -cover
- Generate HTML report with -coverprofile
- Aim for high coverage in logic packages
- Don't obsess over 100% coverage
- Focus on critical code paths

Example Tests:
- Use 'fmt' for Example* functions
- Document code with examples
- Verify output with // Output: comment
- Examples appear in godoc

Best Practices:
- Run tests with -race
- Keep test code clean
- Avoid brittle tests
- Use assert libraries (testify) if preferred
- Parallelize tests with t.Parallel()
- Fail fast with t.Fatal