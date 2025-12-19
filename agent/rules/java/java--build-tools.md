---
trigger: model_decision
description: When it is about Java development, apply these rules.
---
You are an expert in Java Build Systems.

Gradle (Preferred):
- **DSL**: Use Kotlin DSL (`build.gradle.kts`) for Type-Safety and better IDE autocomplete.
- **Version Catalog**: Use `libs.versions.toml` to manage dependency versions centrally.
- **Wrapper**: Always commit the `gradle/wrapper` (jar and properties) to ensure reproducible builds.
- **Multi-Module**: Isolate features into modules (`:core`, `:api`, `:web`) to enforce separation of concerns and speed up incremental compilation.
- **Tasks**: Do not use `doLast` for everything. Create custom Task classes for complex logic.

Maven (Enterprise Standard):
- **POM Organization**: Use a `pom.xml` with properly defined `<dependencyManagement>` section to control versions in multi-module projects.
- **Wrapper**: Commit `.mvn/wrapper` and `mvnw` script.
- **Plugins**: Lock down plugin versions in `<pluginManagement>` to ensure build stability over time.
- **Profiles**: Use profiles sparingly, mostly for environment-specific build behavior (e.g., skip tests in specific CI stages).

General Dependency Management:
- **Minimalism**: Don't add a library for a one-line utility function.
- **Security**: Regularly scan dependencies for CVEs (e.g., using OWASP Dependency Check or Snyk).
- **Exclusions**: Explicitly exclude transitive dependencies that conflict or are unused to keep the artifact size small.
