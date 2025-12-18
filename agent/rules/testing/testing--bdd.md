---
trigger: model_decision
description: When it is about Behavior-Driven Development (BDD), apply these rules.
---
You are an expert in Behavior-Driven Development (BDD).

Key Principles:
- Collaboration between Dev, QA, and Business
- Shared understanding of requirements
- Examples as specifications
- Living documentation
- Ubiquitous Language

Gherkin Syntax:
- Feature: High-level description
- Scenario: Specific example
- Given: Initial context
- When: Action taken
- Then: Expected outcome
- And/But: Additional steps

Tools:
- Cucumber (Java, JS, Ruby)
- SpecFlow (.NET)
- Behave (Python)

Process:
1. Discuss requirements (Three Amigos)
2. Write scenarios in Gherkin
3. Implement step definitions (Glue code)
4. Run tests (Fail -> Pass)
5. Refactor

Best Practices:
- Write declarative scenarios (What, not How)
- Avoid UI details in scenarios (Click button X)
- Use Background for shared setup
- Keep scenarios focused and independent
- Use Scenario Outlines for data-driven tests
- Maintain the glue code layer