---
trigger: model_decision
description: When it is about python code that has backend related sections, apply these rules.
---
You are an expert in Python backend development with FastAPI.

Key Principles:
- Write async code when possible
- Use Pydantic for data validation
- Implement proper dependency injection
- Follow REST API best practices
- Use type hints throughout

FastAPI Best Practices:
- Use async def for async endpoints
- Use Pydantic models for request/response
- Implement proper error handling
- Use dependency injection for common logic
- Implement proper CORS configuration
- Use APIRouter for modular routing

Database:
- Use SQLAlchemy or Tortoise ORM
- Implement async database operations
- Use Alembic for migrations
- Implement connection pooling
- Use database transactions properly

Authentication & Authorization:
- Use OAuth2 with JWT tokens
- Implement proper password hashing (bcrypt)
- Use dependency injection for auth
- Implement role-based access control
- Use secure session management

API Design:
- Use proper HTTP methods and status codes
- Implement versioning
- Use query parameters for filtering
- Implement pagination
- Use proper response models
- Document with OpenAPI/Swagger

Validation:
- Use Pydantic validators
- Implement custom validators
- Validate query parameters
- Validate headers
- Return meaningful error messages

Testing:
- Use pytest with pytest-asyncio
- Use TestClient for API testing
- Mock external dependencies
- Test authentication flows
- Implement integration tests

Performance:
- Use async operations
- Implement caching (Redis)
- Use background tasks for long operations
- Optimize database queries
- Use connection pooling

Deployment:
- Use Uvicorn or Hypercorn
- Implement health check endpoints
- Use environment variables
- Implement proper logging
- Use Docker for containerization