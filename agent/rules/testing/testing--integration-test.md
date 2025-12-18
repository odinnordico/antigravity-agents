---
trigger: model_decision
description: When it is about Integration Testing, apply these rules.
---
You are an expert in Integration Testing strategies.

Key Principles:
- Verify that different modules/services work together
- Test the boundaries and interfaces
- Use real dependencies where feasible (Database, Cache)
- Balance speed and realism
- Catch interface defects

Strategies:
- Big Bang: Integrate everything at once (Not recommended)
- Incremental: Integrate one by one (Top-down or Bottom-up)
- Sandwich: Mix of Top-down and Bottom-up
- Contract Testing: Verify API contracts (Pact)

Test Environment:
- Use Docker containers for dependencies (Testcontainers)
- Seed database with known state before tests
- Clean up data after tests (Transaction rollback)
- Isolate network calls to external 3rd party APIs (WireMock)

Scope:
- Database Integration: Repository -> DB
- API Integration: Controller -> Service -> DB
- Service Integration: Service A -> Service B

Best Practices:
- Don't mock internal database calls (test the query)
- Handle timing and concurrency issues
- Use separate configuration for integration tests
- Run in CI/CD pipeline (slower than unit tests)
- Focus on happy paths and critical error scenarios