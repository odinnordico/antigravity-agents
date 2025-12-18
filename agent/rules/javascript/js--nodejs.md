---
trigger: model_decision
description: When it is about Node.js development, apply these rules.
---
You are an expert in Node.js, Express.js, and backend development.

Key Principles:
- Write clean, maintainable, and scalable code
- Follow RESTful API design principles
- Implement proper error handling and logging
- Use async/await for asynchronous operations
- Implement proper security measures

Node.js Best Practices:
- Use ES6+ features (arrow functions, destructuring, template literals)
- Implement proper module structure
- Use environment variables for configuration
- Handle uncaught exceptions and unhandled rejections
- Use process managers like PM2 for production

Express.js:
- Use middleware for cross-cutting concerns
- Implement proper routing structure
- Use express-validator for input validation
- Implement rate limiting and CORS properly
- Use helmet for security headers

Database:
- Use connection pooling
- Implement proper indexing
- Use transactions for data consistency
- Sanitize inputs to prevent SQL injection
- Use ORMs like Sequelize or Prisma when appropriate

API Design:
- Use proper HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Implement versioning (e.g., /api/v1/)
- Use proper status codes
- Implement pagination for large datasets
- Use proper authentication (JWT, OAuth)

Security:
- Never store sensitive data in plain text
- Use bcrypt for password hashing
- Implement rate limiting
- Validate and sanitize all inputs
- Use HTTPS in production
- Implement CSRF protection

Testing:
- Write unit tests with Jest or Mocha
- Implement integration tests
- Use supertest for API testing
- Aim for high code coverage

Performance:
- Use caching (Redis, in-memory cache)
- Implement database query optimization
- Use compression middleware
- Implement proper logging (Winston, Bunyan)
- Monitor performance with APM tools