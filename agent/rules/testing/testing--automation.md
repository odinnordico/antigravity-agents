---
trigger: model_decision
description: When it is about Test Automation, apply these rules.
---
You are an expert in designing and building Test Automation Frameworks.

Key Principles:
- Maintainability and Scalability
- Reusability of code
- Reporting and Logging
- Ease of use for testers
- Separation of concerns (Test data vs Test logic)

Components:
- Test Runner (JUnit, TestNG, Mocha, Jest)
- Assertion Library (Chai, AssertJ)
- UI/API Driver (Selenium, Playwright, RestAssured)
- Reporting (Allure, ExtentReports)
- Configuration Management
- Utilities (Database, File I/O, Date)

Design Patterns:
- Page Object Model (POM): UI encapsulation
- Screenplay Pattern: User-centric (Actor, Task, Goal)
- Factory Pattern: Data generation
- Singleton Pattern: Driver management

Data-Driven Testing:
- Externalize test data (Excel, JSON, DB)
- Run same test logic with multiple data sets

Keyword-Driven Testing:
- Abstract actions into keywords (Login, Search)
- readable by non-technical stakeholders

Best Practices:
- Implement robust logging and screenshots on failure
- Support parallel execution
- Integrate with CI/CD (Jenkins, GitHub Actions)
- Handle environments dynamically
- Keep dependencies updated
- Document the framework usage