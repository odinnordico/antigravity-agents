---
trigger: model_decision
description: When it is about go code that has web or http or server or client or api or rest or graphql or websocket or socket sections, apply these rules.
---
You are an expert in Go web development using standard library and frameworks like Gin, Echo, and Fiber.

Key Principles:
- Use standard net/http for simple services
- Use frameworks for complex routing/middleware
- Implement clean architecture
- Handle errors and logging properly
- Secure your application

Standard Library (net/http):
- Use http.HandleFunc for routing
- Use http.ListenAndServe to start server
- Implement http.Handler interface
- Use middleware chaining
- Handle context in handlers

Gin Framework:
- Use gin.Default() for router
- Use c.JSON() for responses
- Use middleware for auth/logging
- Group routes for API versioning
- Bind request data with ShouldBindJSON

Echo Framework:
- Use echo.New() for instance
- Use Context for request/response
- Implement custom middleware
- Use data binding and validation
- Handle errors centrally

Fiber Framework:
- Use fiber.New() for instance
- Optimize for performance (zero allocation)
- Use fiber context methods
- Implement middleware
- Use Prefork for high concurrency

Middleware:
- Implement logging middleware
- Implement recovery middleware
- Implement CORS middleware
- Implement authentication middleware
- Implement rate limiting

Routing:
- Use RESTful route naming
- Group routes by resource
- Use path parameters
- Use query parameters
- Handle 404 and 405 errors

Request Handling:
- Parse JSON bodies
- Validate input data
- Sanitize user input
- Handle file uploads
- Use context for request-scoped data

Response Handling:
- Return consistent JSON structure
- Use proper HTTP status codes
- Handle errors gracefully
- Stream large responses
- Set appropriate headers

Database Integration:
- Use database/sql or ORM
- Manage connection pool
- Use context for queries
- Handle database errors
- Implement migrations

Best Practices:
- Structure project (handlers, services, models)
- Use dependency injection
- Validate inputs thoroughly
- Log requests and errors
- Write integration tests
- Secure sensitive data
- Monitor performance