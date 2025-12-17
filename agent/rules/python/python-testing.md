---
trigger: model_decision
description: When it is about python code that has testing related sections, apply these rules.
---
You are an expert in Python testing with pytest and testing best practices.

Key Principles:
- Write tests before or alongside code (TDD/BDD)
- Aim for high test coverage (80%+ for critical code)
- Make tests independent and isolated
- Follow the AAA pattern (Arrange, Act, Assert)
- Keep tests simple and readable

Pytest Fundamentals:
- Use pytest instead of unittest for modern testing
- Use assert statements (pytest rewrites them)
- Use fixtures for setup and teardown
- Use parametrize for testing multiple inputs
- Use markers to categorize tests (@pytest.mark.slow)

Test Organization:
- Mirror source code structure in tests/
- Name test files test_*.py or *_test.py
- Name test functions test_*
- Group related tests in classes (TestClassName)
- Use conftest.py for shared fixtures

Fixtures:
- Use @pytest.fixture decorator
- Use scope parameter (function, class, module, session)
- Use yield for setup/teardown
- Use autouse=True for automatic fixtures
- Use fixture factories for dynamic fixtures

Mocking and Patching:
- Use unittest.mock or pytest-mock
- Mock external dependencies (APIs, databases)
- Use patch() decorator or context manager
- Verify mock calls with assert_called_with()
- Use MagicMock for complex objects

Parametrized Testing:
- Use @pytest.mark.parametrize for multiple test cases
- Test edge cases and boundary conditions
- Use pytest.param() for custom test IDs
- Combine multiple parametrize decorators
- Use indirect parametrization for fixtures

Assertion Best Practices:
- Use descriptive assertion messages
- Test one concept per test function
- Use pytest.raises() for exception testing
- Use pytest.approx() for floating-point comparisons
- Use pytest.warns() for warning testing

Test Coverage:
- Use pytest-cov for coverage reporting
- Run: pytest --cov=myproject --cov-report=html
- Aim for 80%+ coverage on critical code
- Don't obsess over 100% coverage
- Focus on testing behavior, not implementation

Integration Testing:
- Test interactions between components
- Use docker-compose for external dependencies
- Use pytest-docker for container management
- Test database migrations
- Test API endpoints end-to-end

Performance Testing:
- Use pytest-benchmark for benchmarking
- Set performance thresholds
- Test with realistic data volumes
- Profile slow tests
- Use pytest-timeout to catch hanging tests

Test Data Management:
- Use factories (factory_boy) for test data
- Use fixtures for reusable test data
- Don't use production data in tests
- Use faker for generating realistic data
- Clean up test data after tests

Continuous Integration:
- Run tests on every commit
- Use tox for testing multiple Python versions
- Use pre-commit hooks for fast feedback
- Fail builds on test failures
- Track test coverage trends

Best Practices:
- Write descriptive test names
- Keep tests fast (mock slow operations)
- Don't test third-party code
- Refactor tests like production code
- Use test-driven development when appropriate