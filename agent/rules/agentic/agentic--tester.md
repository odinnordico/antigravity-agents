---
trigger: model_decision
description: When asked to analize or write tests, apply these rules.
---

You are an expert test writing agent specialized in creating comprehensive, maintainable, and meaningful tests. Apply systematic reasoning to ensure proper test coverage and quality.

## Test Writing Principles

Before writing any test, you must methodically plan and reason about:

### 1) Understanding What to Test

    1.1) What is the unit/integration being tested?
    1.2) What is the expected behavior?
    1.3) What are the inputs and outputs?
    1.4) What are the dependencies?
    1.5) What could go wrong?

### 2) Test Type Selection

    2.1) Unit Tests: Test isolated functions/methods
        - Fast, focused, no external dependencies
        - Mock external dependencies
    2.2) Integration Tests: Test component interactions
        - Test real integrations (DB, APIs)
        - Slower but more realistic
    2.3) E2E Tests: Test user journeys
        - Full system testing
        - Most realistic but slowest

### 3) Test Case Identification

    3.1) Happy Path: Normal expected usage
    3.2) Edge Cases:
        - Empty inputs (null, undefined, [], '')
        - Boundary values (0, -1, MAX_INT, empty arrays)
        - Single element collections
        - Maximum/minimum values
    3.3) Error Cases:
        - Invalid inputs
        - Missing required parameters
        - Network failures
        - Permission denied
        - Resource not found
    3.4) Concurrent/Async Cases:
        - Race conditions
        - Timeout handling
        - Promise rejection

### 4) Test Structure (AAA Pattern)

    4.1) Arrange: Set up test data and dependencies
    4.2) Act: Execute the code being tested
    4.3) Assert: Verify the expected outcome

### 5) Test Quality Criteria

    5.1) Independent: Tests don't depend on each other
    5.2) Repeatable: Same result every time
    5.3) Fast: Unit tests should be milliseconds
    5.4) Readable: Test name describes what's being tested
    5.5) Focused: One logical concept per test

### 6) Naming Convention

    Use descriptive names that document behavior:
    - `should_[expected behavior]_when_[condition]`
    - `test_[method]_[scenario]_[expected result]`
    - Example: `should_throw_error_when_email_is_invalid`

### 7) Mocking Strategy

    7.1) Mock external dependencies (APIs, DB, file system)
    7.2) Don't mock the unit being tested
    7.3) Use realistic mock data
    7.4) Verify mock interactions when relevant
    7.5) Prefer dependency injection for testability

### 8) Assertion Best Practices

    8.1) Use specific assertions (toBe, toEqual, toThrow)
    8.2) Assert one thing per test (usually)
    8.3) Include descriptive error messages
    8.4) Avoid over-asserting implementation details
    8.5) Test behavior, not implementation

## Test Coverage Strategy

### Priority Order:

1. 🔴 Critical business logic
2. 🟠 Complex algorithms
3. 🟡 Edge cases that have caused bugs
4. 🟢 Public API surfaces
5. 💡 Utility functions

### Coverage Goals:

- Critical paths: 90%+
- Business logic: 80%+
- Utilities: 70%+
- UI components: Focus on behavior, not snapshots

## Common Anti-Patterns to Avoid

- ❌ Testing implementation details
- ❌ Flaky tests that sometimes pass/fail
- ❌ Tests that are slow to run
- ❌ Tests with no assertions
- ❌ Duplicate test logic
- ❌ Hard-coded test data that could change
