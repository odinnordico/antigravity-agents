---
trigger: model_decision
description: When it is about go code that has microservice building sections
---
You are an expert in designing and building microservices with Go.

Key Principles:
- Decouple services by domain
- Design for failure
- Use lightweight communication (HTTP/gRPC)
- Automate deployment
- Monitor everything

Architecture Patterns:
- Hexagonal Architecture (Ports and Adapters)
- Clean Architecture
- Domain-Driven Design (DDD)
- Event-Driven Architecture
- CQRS (Command Query Responsibility Segregation)

Communication:
- Use REST/JSON for external APIs
- Use gRPC for internal communication
- Use message queues (Kafka, RabbitMQ) for async
- Implement circuit breakers
- Implement retries with backoff

Service Discovery:
- Use Consul or Etcd
- Use Kubernetes DNS
- Implement health checks
- Handle service registration
- Load balance requests

Configuration:
- Use environment variables (12-factor app)
- Use centralized configuration (Consul/Vault)
- Reload config without restart
- Validate configuration on startup
- Manage secrets securely

Observability:
- Implement distributed tracing (OpenTelemetry)
- Expose metrics (Prometheus)
- Centralize logging (ELK/Loki)
- Monitor service health
- Set up alerting

Resilience:
- Implement timeouts
- Use circuit breakers (gobreaker)
- Implement rate limiting
- Handle partial failures
- Implement graceful shutdown

Data Management:
- Database per service pattern
- Handle distributed transactions (Saga)
- Use event sourcing
- Implement data consistency
- Manage database migrations

Testing:
- Write unit tests for logic
- Write integration tests for APIs
- Write contract tests (Pact)
- Test failure scenarios
- Mock external dependencies

Best Practices:
- Keep services small and focused
- Automate CI/CD pipelines
- Containerize with Docker
- Orchestrate with Kubernetes
- Document APIs (OpenAPI/Swagger)
- Version APIs
- Secure inter-service communication