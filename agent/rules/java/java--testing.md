---
trigger: model_decision
description: When it is about Java development, apply these rules.
---
You are an expert in Java Testing using JUnit 5, Mockito, and Testcontainers.

General Testing Philosophy:
- **Test Pyramid**: Prioritize Unit Tests > Integration Tests > E2E Tests.
- **Naming**: Use descriptive names: `shouldThrowExceptionWhenUserNotFound()` or `Given_UserExists_When_Login_Then_ReturnToken`.

Frameworks & Tools:
- **JUnit 5 (Jupiter)**: The standard test runner.
- **AssertJ**: Use fluent assertions (`assertThat(actual).isEqualTo(expected)`) instead of JUnit's `assertEquals`.
- **Mockito**: For isolating unit tests.
- **Testcontainers**: For integration tests involving databases/brokers. NEVER mock the database in integration tests; use a real containerized instance.

Unit Testing Best Practices:
- One assertion per test concept (logical assertion, not necessarily one line).
- **Mocking**:
  - Only mock dependencies that you own or interact with directly.
  - Use `@ExtendWith(MockitoExtension.class)` and `@Mock`/`@InjectMocks`.
  - Avoid `Mockito.any()` when precise matching is possible.

Integration Testing:
- Annotate integration tests with `@SpringBootTest` (for full context) or `@DataJpaTest` (for slice testing).
- **Testcontainers**:
  - Use `@Container` and `@Testcontainers`.
  - Prefer Singleton Container pattern (static container) to speed up test suites by reusing the container across methods.
  ```java
  @Testcontainers
  class UserRepositoryTest {
      @Container
      static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:15");
      // ...
  }
  ```

Parameterized Tests:
- Use `@ParameterizedTest` for testing multiple data scenarios without duplicating logic.
- Use `@CsvSource` or `@MethodSource` for input data.

Properties:
- Avoid `@DirtyContext` as it kills performance by reloading the Spring Context.
