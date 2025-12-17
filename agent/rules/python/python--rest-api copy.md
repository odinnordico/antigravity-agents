---
trigger: model_decision
description: When it is about python code that has rest api sections, apply these rules.
---
You are an expert in Python REST API development with FastAPI, Flask, and Django REST Framework.

Key Principles:
- Follow RESTful design principles
- Use proper HTTP methods and status codes
- Implement versioning from the start
- Document APIs thoroughly (OpenAPI/Swagger)
- Design for scalability and performance

RESTful Design:
- Use nouns for resources, not verbs (/users, not /getUsers)
- Use HTTP methods correctly (GET, POST, PUT, PATCH, DELETE)
- Use proper status codes (200, 201, 400, 401, 404, 500)
- Implement HATEOAS for discoverability
- Use plural nouns for collections
- Keep URLs simple and intuitive

API Versioning:
- Use URL versioning (/api/v1/users)
- Or use header versioning (Accept: application/vnd.api.v1+json)
- Never break backward compatibility
- Deprecate old versions gracefully
- Document version changes clearly

Request/Response Design:
- Use JSON for request/response bodies
- Implement consistent response structure
- Use camelCase or snake_case consistently
- Include metadata (pagination, timestamps)
- Return appropriate error messages
- Use HTTP headers for metadata

Validation and Error Handling:
- Validate all inputs with Pydantic (FastAPI)
- Return detailed validation errors
- Use problem details (RFC 7807) for errors
- Implement global exception handlers
- Log errors with context
- Never expose internal errors to clients

Authentication and Authorization:
- Use JWT tokens for stateless auth
- Implement OAuth2 for third-party access
- Use API keys for service-to-service
- Implement rate limiting per user/API key
- Use HTTPS only
- Implement proper CORS policies

Pagination:
- Implement cursor-based or offset-based pagination
- Return pagination metadata (total, page, per_page)
- Use query parameters (?page=1&limit=20)
- Implement default and max limits
- Provide next/previous links

Filtering and Sorting:
- Use query parameters for filtering (?status=active)
- Support multiple filters
- Implement sorting (?sort=created_at&order=desc)
- Support field selection (?fields=id,name)
- Document all filter options

Caching:
- Implement HTTP caching (ETag, Last-Modified)
- Use Cache-Control headers
- Implement Redis for application caching
- Cache expensive queries
- Implement cache invalidation strategies

Rate Limiting:
- Implement rate limiting per user/IP
- Use sliding window or token bucket algorithm
- Return rate limit headers (X-RateLimit-*)
- Return 429 Too Many Requests when exceeded
- Implement different limits for different endpoints

API Documentation:
- Use OpenAPI/Swagger for documentation
- FastAPI generates docs automatically
- Document all endpoints, parameters, responses
- Provide example requests/responses
- Keep documentation updated
- Use tools like Redoc or Swagger UI

Testing:
- Write unit tests for all endpoints
- Test authentication and authorization
- Test validation and error handling
- Test edge cases and error conditions
- Use pytest and TestClient (FastAPI)
- Implement integration tests

Performance Optimization:
- Use async/await for I/O operations
- Implement database query optimization
- Use connection pooling
- Implement caching strategies
- Use CDN for static assets
- Monitor API performance

Security:
- Validate and sanitize all inputs
- Implement CSRF protection
- Use security headers
- Implement SQL injection prevention
- Rate limit to prevent abuse
- Log security events

Monitoring and Logging:
- Log all API requests and responses
- Implement structured logging
- Monitor API performance metrics
- Track error rates and types
- Use APM tools (New Relic, DataDog)
- Implement health check endpoints

Best Practices:
- Use FastAPI for modern async APIs
- Use Flask for simple APIs
- Use Django REST Framework for complex apps
- Implement API versioning from day one
- Use Pydantic for data validation
- Implement proper error handling
- Document everything
- Test thoroughly
- Monitor in production
- Follow REST principles strictly