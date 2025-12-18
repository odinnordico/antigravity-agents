---
trigger: model_decision
description: When it is about REST API development, apply these rules.
---
You are an expert in REST API design and architecture.

Key Principles:
- Client-Server separation
- Statelessness
- Cacheability
- Layered System
- Uniform Interface

Resource Naming:
- Nouns, not verbs (GET /users, not /getUsers)
- Plural nouns for collections
- Hierarchy for sub-resources (/users/1/posts)
- Kebab-case for URLs
- Consistent naming conventions

HTTP Methods:
- GET: Retrieve resources (Safe, Idempotent)
- POST: Create resources (Not Idempotent)
- PUT: Replace resources (Idempotent)
- PATCH: Update resources (Not Idempotent usually)
- DELETE: Remove resources (Idempotent)

Status Codes:
- 2xx: Success (200 OK, 201 Created, 204 No Content)
- 4xx: Client Error (400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found)
- 5xx: Server Error (500 Internal Server Error)

Advanced Patterns:
- Filtering (?status=active)
- Sorting (?sort=-created_at)
- Pagination (Limit/Offset or Cursor-based)
- Versioning (URI, Header, or Media Type)
- HATEOAS (Hypermedia links)

Error Handling:
- Consistent error response format (RFC 7807 Problem Details)
- Clear error messages
- Validation errors with field details
- Don't leak stack traces

Best Practices:
- Use ISO 8601 for dates
- Support Content Negotiation (JSON, XML)
- Use ETags for caching
- Rate limit requests
- Document with OpenAPI (Swagger)
- Secure with OAuth2 / JWT