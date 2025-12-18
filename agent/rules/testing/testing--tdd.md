---
trigger: model_decision
description: When it is about Test-Driven Development (TDD), apply these rules.
---
You are an expert in Test-Driven Development (TDD).

Key Principles:
- Write the test BEFORE the code
- Red-Green-Refactor cycle
- Design emerges from tests
- Small, incremental steps
- Confidence in refactoring

The Cycle:
1. Red: Write a failing test for a small feature.
2. Green: Write just enough code to pass the test.
3. Refactor: Improve the code quality without changing behavior.

Benefits:
- Better code design (loosely coupled, highly cohesive)
- Built-in regression suite
- Documentation of behavior
- Reduced debugging time
- Fearless refactoring

Types of TDD:
- Unit Level TDD: Focus on classes/functions
- Acceptance TDD (ATDD): Focus on user requirements

Common Pitfalls:
- Writing too many tests at once
- Writing complex tests
- Skipping the Refactor step
- Testing implementation details instead of behavior
- Giving up when it gets hard

Best Practices:
- Start with the simplest case
- Treat test code as first-class citizen
- Run all tests frequently
- Listen to your tests (if it's hard to test, design might be wrong)
- Pair programming enhances TDD