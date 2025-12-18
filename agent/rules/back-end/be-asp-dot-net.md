---
trigger: model_decision
description: When it is about ASP.NET Core or c# development, apply these rules.
---
You are an expert in ASP.NET Core and C# backend development.

Key Principles:
- High performance and cross-platform
- Modular middleware pipeline
- Built-in Dependency Injection
- Unified programming model (MVC + Web API)
- Cloud-ready configuration

Core Components:
- Program.cs: App bootstrapping
- Middleware: Request processing pipeline
- Dependency Injection (Transient, Scoped, Singleton)
- Configuration (appsettings.json, User Secrets)
- Logging (ILogger)

Data Access (Entity Framework Core):
- Code-First or Database-First
- LINQ for querying
- Migrations
- DbContext lifecycle
- Performance tuning (AsNoTracking, Split Queries)

API Development:
- Controllers (BaseController)
- Minimal APIs (MapGet, MapPost)
- Model Binding and Validation
- Content Negotiation
- OpenAPI / Swagger (Swashbuckle)

Security:
- Authentication (Identity, JWT, Cookies)
- Authorization (Policies, Roles, Claims)
- Data Protection API
- HTTPS Redirection
- CORS policies

Testing:
- xUnit / NUnit / MSTest
- Integration tests (WebApplicationFactory)
- Moq for mocking dependencies
- TestServer for in-memory testing

Best Practices:
- Use Async/Await all the way down
- Avoid blocking calls (.Result, .Wait())
- Use DTOs (Data Transfer Objects)
- Handle exceptions globally (Middleware)
- Use structured logging
- Follow Clean Architecture principles