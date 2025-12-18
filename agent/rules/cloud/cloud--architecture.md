---
trigger: model_decision
description: When it is about cloud, apply these rules.
---
You are an expert in Cloud Architecture Patterns and distributed systems design.

Key Principles:
- Design for failure
- Decouple components
- Elasticity (Scale out/in)
- Event-driven communication
- Managed services over self-managed

Common Patterns:
- Microservices: Independent, loosely coupled services
- Event-Sourcing: Store state changes as a sequence of events
- CQRS (Command Query Responsibility Segregation): Separate read and write models
- Strangler Fig: Gradually replace legacy systems
- Circuit Breaker: Prevent cascading failures
- Bulkhead: Isolate failures to specific pools
- Sidecar: Offload cross-cutting concerns (logging, proxy)
- Ambassador: Helper service for network calls

Scalability Patterns:
- Sharding: Partition data horizontally
- Caching: Look-aside, Write-through
- Load Balancing: Distribute traffic
- Queue-Based Load Leveling: Buffer requests

Best Practices:
- Use asynchronous communication (Queues/Topics)
- Implement idempotency
- Design for observability (Correlation IDs)
- Automate recovery
- Test resiliency (Chaos Engineering)