---
trigger: model_decision
description: When it is about Java development, apply these rules.
---
You are an expert in Modern Java (Java 21+) programming.

Key Principles:
- **Adopt Java 21 Features**: Use Records, Pattern Matching, and Virtual Threads where applicable.
- **Immutability**: Prefer immutable data structures. Use `final` keyword and Java Records (`record User(String name) {}`).
- **Readability**: Follow "Clean Code" principles. Code is read more often than written.
- **Functional Style**: Use Streams API and Lambda expressions for data processing, but don't overcomplicate debuggability.

Modern Language Features:
- **Records**: Use `record` for data carriers (DTOs, configuration objects) instead of Lombok `@Data` where possible.
  ```java
  public record Customer(String id, String name) {}
  ```
- **Pattern Matching**: Use pattern matching for `switch` and `instanceof`.
  ```java
  if (obj instanceof String s) {
      System.out.println(s.toLowerCase());
  }
  ```
- **Sequenced Collections**: Use `SequencedCollection` interfaces (e.g., `list.getFirst()`, `list.getLast()`) available in Java 21.
- **Var**: Use local variable type inference (`var`) when the type is obvious from the right-hand side.

Code Style & Formatting:
- **Style Guide**: Follow Google Java Style Guide.
- **Indentation**: Use **4 spaces** for indentation.
- **Braces**: Opening braces on the same line (K&R style).
- **Naming**:
  - Classes: `PascalCase`
  - Methods/Variables: `camelCase`
  - Constants: `UPPER_SNAKE_CASE`

Error Handling:
- Avoid returning `null`. Return `Optional<T>` for methods that may not produce a value.
- Use explicit, custom exceptions (e.g., `UserNotFoundException`) rather than generic `RuntimeException`.
- Don't swallow exceptions. Log them or rethrow them wrapped in a business exception.

Null Safety:
- Assume non-null by default.
- Use `Objects.requireNonNull()` or proper `@NonNull` annotations if working with external libraries.
