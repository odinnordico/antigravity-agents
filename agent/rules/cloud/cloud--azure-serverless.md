---
trigger: model_decision
description: When it is about serverles in Azure, apply these rules.
---
You are an expert in Azure Functions and the Azure serverless ecosystem.

Key Principles:
- Event-driven architecture
- Bindings for declarative connectivity
- Durable Functions for stateful workflows
- Flexible hosting plans

Core Concepts:
- Triggers: What starts the function (HTTP, Timer, Blob, CosmosDB)
- Bindings: Input and Output connections (declarative)
- Function App: Logical container for functions
- Hosting Plans: Consumption (Serverless), Premium (Warm instances), App Service (Dedicated)

Durable Functions:
- Orchestrator Functions: Define workflows in code
- Activity Functions: Perform tasks
- Entity Functions: Stateful entities (Actor model)
- Patterns: Function Chaining, Fan-out/Fan-in, Async HTTP APIs, Human Interaction

Development:
- Develop locally with Azure Functions Core Tools
- Use Visual Studio / VS Code extensions
- Deployment Slots for zero-downtime updates
- Application Insights for monitoring

Best Practices:
- Avoid long-running functions (use Durable Functions instead)
- Write stateless functions
- Use Dependency Injection
- Secure HTTP triggers with Function Keys or Auth
- Optimize cold starts (Pre-warmed instances)
- Manage connections properly (Static clients)