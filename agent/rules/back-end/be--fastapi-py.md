---
trigger: model_decision
description: When it is about FastAPI development with Python, apply these rules.
---
You are an expert in FastAPI for building high-performance Python APIs.

Key Principles:
- High performance (on par with NodeJS/Go)
- Easy to learn and fast to code
- Type hints are first-class citizens
- Automatic interactive documentation
- Standards-based (OpenAPI, JSON Schema)

Core Concepts:
- Path Operations (decorators like @app.get)
- Path Parameters and Query Parameters
- Request Body with Pydantic Models
- Dependency Injection System
- Background Tasks

Pydantic Models:
- Define data schemas with Python types
- Automatic data validation and parsing
- Use Field() for metadata and validation
- Separate Input and Output models (Response Model)
- Use Config for model settings

Dependency Injection:
- Create reusable dependencies (Depends)
- Handle authentication and authorization
- Manage database sessions
- Share logic across endpoints
- Support sub-dependencies

Async/Await:
- Use async def for I/O bound operations
- Use def for CPU bound operations (run in threadpool)
- Integrate with async libraries (SQLAlchemy 2.0, Tortoise ORM)
- Handle concurrency effectively

Security:
- OAuth2 with Password Flow and JWT
- API Key authentication
- CORS middleware configuration
- TrustedHost middleware
- HTTPS enforcement

Best Practices:
- Structure app with APIRouter
- Use lifespan events for startup/shutdown
- Handle errors with exception handlers
- Write tests using TestClient
- Use environment variables (pydantic-settings)
- Pin dependencies