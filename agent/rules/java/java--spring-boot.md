---
trigger: model_decision
description: When it is about Java development with Spring, apply these rules.
---
You are an expert in Spring Boot 3+ development.

Architecture:
- **Layered Architecture**: Strictly standard interactions: Controller → Service → Repository.
  - Controllers should handle HTTP mapping and DTO conversion.
  - Services should contain all business logic.
  - Repositories should only interact with the database.
- **DTO Pattern**: NEVER expose Entity classes (`@Entity`) in the API. Use Records as DTOs.
  ```java
  public record UserResponse(String id, String email) {}
  ```
- **Dependency Injection**: Use Constructor Injection. Avoid `@Autowired` on fields.
  ```java
  @Service
  public class UserService {
      private final UserRepository repo;

      public UserService(UserRepository repo) {
          this.repo = repo;
      }
  }
  ```

Configuration:
- **Properties**: Use `application.yml` over `application.properties`.
- **Type-Safe Configuration**: Use `@ConfigurationProperties` classes instead of `@Value` for grouped properties.
- **Profiles**: Use strictly defined profiles (`dev`, `staging`, `prod`) for environment-specific config.

Observability & Ops:
- **Actuator**: Always enable Spring Boot Actuator for health and metrics.
- **Logging**: Use SLF4J. Do not use `System.out.println`.
- **OpenTelemetry**: Integrate Micrometer Tracing for distributed tracing.

Data Access:
- Use **Spring Data JPA** repositories.
- For complex queries, prefer `@Query` or explicit Specifications/Criteria API over complex derived method names.
- Avoid N+1 problems by using `JOIN FETCH` or EntityGraphs.

Security:
- Use **Spring Security**.
- Prefer declarative method security (`@PreAuthorize`) over configuration-based matching where possible for improved granularity.
