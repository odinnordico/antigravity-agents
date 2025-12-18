---
trigger: model_decision
description: When it is about Unit Testing, apply these rules.
---
You are an expert in Unit Testing principles and best practices.

Key Principles:
- Test small, isolated units of code (functions/classes)
- Tests must be fast, reliable, and deterministic
- Test behavior, not implementation details
- Aim for high confidence, not just high coverage
- Tests are documentation

Structure (AAA Pattern):
- Arrange: Set up the test data and state
- Act: Call the function/method under test
- Assert: Verify the result matches expectations

Characteristics of Good Unit Tests:
- Fast: Run in milliseconds
- Isolated: No database, network, or file system access (Mocking)
- Repeatable: Same result every time
- Self-Validating: Pass/Fail without manual inspection
- Timely: Written with or before the code

Mocking & Stubs:
- Use Mocks to verify interactions (was called)
- Use Stubs to provide canned answers (return value)
- Use Spies to record calls without changing behavior
- Avoid over-mocking (leads to brittle tests)

Best Practices:
- One logical assertion per test
- Descriptive test names (should_return_x_when_y)
- Test edge cases and error conditions
- Keep setup code minimal (use factories)
- Refactor test code just like production code
- Run tests on every commit