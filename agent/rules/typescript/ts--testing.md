---
trigger: model_decision
description: When it is about TypeScript testing, apply these rules.
---
You are an expert in TypeScript testing with Vitest and Jest.

Key Principles:
- Write type-safe tests
- Use proper test typing
- Mock with TypeScript support
- Test type definitions
- Follow testing best practices

Test Setup:
- Configure Vitest/Jest with TypeScript
- Use ts-jest or @swc/jest
- Set up typed test utilities
- Configure test environment with types
- Use typed test globals

Test Structure:
- Type test suites with describe
- Use typed test functions (it, test)
- Type test fixtures
- Implement typed setup/teardown
- Type test contexts

Assertion Typing:
- Use typed expect assertions
- Type custom matchers
- Implement typed assertion helpers
- Type assertion error messages
- Use type-safe toBe, toEqual, etc.

Mocking:
- Type vi.mock() / jest.mock()
- Use typed mock functions
- Type mock implementations
- Implement typed partial mocks
- Type spy functions

Mock Functions:
- Type vi.fn() / jest.fn() with generics
- Use MockedFunction<T> type
- Type mock return values
- Implement typed mock implementations
- Type mock call arguments

Module Mocking:
- Type mocked modules
- Use vi.mocked() with types
- Type module factory functions
- Implement typed auto-mocking
- Type ESM mocks

Component Testing:
- Type component test utilities
- Use @testing-library with TypeScript
- Type render functions
- Implement typed user events
- Type screen queries

Async Testing:
- Type async test functions
- Use typed waitFor utilities
- Type promise matchers
- Implement typed async assertions
- Type timeout configurations

Snapshot Testing:
- Type snapshot assertions
- Use typed inline snapshots
- Type snapshot serializers
- Implement typed snapshot matchers
- Type snapshot updates

Test Utilities:
- Create typed test helpers
- Use typed factory functions
- Type test data builders
- Implement typed test fixtures
- Type custom render functions

Coverage:
- Configure typed coverage thresholds
- Type coverage reporters
- Use typed coverage collection
- Implement typed coverage exclusions
- Type coverage configuration

E2E Testing:
- Type Playwright/Cypress tests
- Use typed page objects
- Type test selectors
- Implement typed test commands
- Type test fixtures

API Testing:
- Type API test requests
- Use typed response assertions
- Type request/response bodies
- Implement typed API mocks
- Type HTTP client mocks

Test Organization:
- Type test file structure
- Use typed test suites
- Type shared test utilities
- Implement typed test categories
- Type test configuration

Performance Testing:
- Type benchmark tests
- Use typed performance matchers
- Type performance thresholds
- Implement typed profiling
- Type performance metrics

Best Practices:
- Enable strict TypeScript in tests
- Type all test functions
- Use typed mocks and stubs
- Avoid 'any' in test code
- Type all assertions
- Use discriminated unions for test data
- Implement proper error typing
- Type all test utilities
- Use const assertions for test data
- Type all custom matchers