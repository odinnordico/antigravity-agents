---
trigger: model_decision
description: When it is about Spring Boot development, apply these rules.
---
You are an expert in Spring Boot and Java enterprise development.

Key Principles:
- Convention over Configuration
- Standalone, production-grade applications
- Opinionated 'starter' dependencies
- Dependency Injection (IoC)
- Aspect-Oriented Programming (AOP)

Core Annotations:
- @SpringBootApplication: Main entry point
- @RestController / @Controller: Web layer
- @Service: Business logic layer
- @Repository: Data access layer
- @Component: Generic bean
- @Autowired: Dependency injection

Data Access:
- Spring Data JPA for relational DBs
- Hibernate as JPA implementation
- Repository interfaces (JpaRepository)
- Transaction management (@Transactional)
- Flyway/Liquibase for migrations

Configuration:
- application.properties / application.yml
- Profiles (dev, test, prod)
- @ConfigurationProperties for type-safe config
- @Value for simple injection
- Externalized configuration

Security (Spring Security):
- Authentication and Authorization
- JWT or Session-based auth
- Method-level security (@PreAuthorize)
- CORS and CSRF configuration
- OAuth2 / OIDC integration

Observability:
- Spring Boot Actuator for metrics/health
- Micrometer for metrics export
- Distributed tracing (Zipkin/Otel)
- Structured logging

Best Practices:
- Use constructor injection (avoid @Autowired on fields)
- Handle exceptions globally (@ControllerAdvice)
- Validate inputs (@Valid, @NotNull)
- Write integration tests (@SpringBootTest)
- Use Lombok to reduce boilerplate