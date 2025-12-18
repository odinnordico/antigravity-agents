---
trigger: model_decision
description: When it is about TypeScript Node.js development, apply these rules.
---
You are an expert in TypeScript Node.js backend development.

Key Principles:
- Use strict TypeScript configuration
- Type all Express middleware and routes
- Implement proper error handling with types
- Use async/await with proper typing
- Follow Node.js best practices

Project Setup:
- Use ts-node-dev for development
- Configure tsconfig.json for Node.js
- Use @types/node for Node.js types
- Set up path aliases
- Configure build output properly

Express Typing:
- Type Request and Response objects
- Use typed middleware functions
- Type route handlers properly
- Implement typed error handlers
- Type Express application

Middleware:
- Type custom middleware
- Use NextFunction with types
- Type middleware parameters
- Implement typed error middleware
- Type third-party middleware

Request/Response:
- Type request body with generics
- Use typed request params
- Type query parameters
- Implement typed response methods
- Type request headers

Routing:
- Type Express Router
- Use typed route parameters
- Type route handlers
- Implement typed route guards
- Type route middleware

Validation:
- Use Zod for runtime validation
- Type validation schemas
- Implement typed validators
- Type validation errors
- Use class-validator with types

Database (TypeORM):
- Type entity classes
- Use typed repositories
- Type query builders
- Implement typed relations
- Type database migrations

Database (Prisma):
- Use Prisma Client types
- Type database queries
- Use generated types
- Implement typed relations
- Type Prisma middleware

Authentication:
- Type JWT payloads
- Use typed auth middleware
- Type user sessions
- Implement typed password hashing
- Type authentication strategies

Authorization:
- Type permission checks
- Use typed role-based access
- Type authorization middleware
- Implement typed policy checks
- Type user permissions

Error Handling:
- Type custom error classes
- Use typed error middleware
- Type error responses
- Implement typed error logging
- Type error codes

Async Operations:
- Type async route handlers
- Use typed Promise returns
- Type async middleware
- Implement typed error handling
- Type async utilities

API Documentation:
- Use Swagger with TypeScript
- Type OpenAPI schemas
- Generate types from OpenAPI
- Implement typed API specs
- Type documentation decorators

Testing:
- Type Supertest requests
- Use typed test utilities
- Type API mocks
- Implement typed integration tests
- Type test fixtures

Logging:
- Type logger instances
- Use typed log levels
- Type log metadata
- Implement typed structured logging
- Type logger configuration

Environment Variables:
- Type process.env properly
- Use typed environment validation
- Type configuration objects
- Implement typed env loaders
- Type environment schemas

Best Practices:
- Enable strict TypeScript mode
- Type all Express handlers
- Use discriminated unions for responses
- Avoid 'any' in backend code
- Type all database operations
- Implement proper error typing
- Type all middleware
- Use const assertions for constants
- Type all async operations
- Use typed dependency injection