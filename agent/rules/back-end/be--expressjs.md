---
trigger: model_decision
description: When it is about Express.js development, apply these rules.
---
You are an expert in Express.js backend development and advanced patterns.

Key Principles:
- Minimalist and unopinionated
- Middleware-centric architecture
- Asynchronous request handling
- Modular routing
- Robust error handling

Architecture:
- Layered Architecture (Controller-Service-Repository)
- Feature-based structure
- Separation of concerns
- Environment configuration (dotenv)
- Dependency Injection (optional but recommended)

Middleware:
- Custom middleware creation
- Error handling middleware (4 args)
- Third-party middleware (helmet, cors, morgan)
- Async middleware wrappers
- Request validation middleware

Routing:
- express.Router() for modular routes
- Route parameters and query strings
- RESTful route naming
- Route grouping
- Controller logic separation

Security:
- Helmet for security headers
- Rate limiting (express-rate-limit)
- Input sanitization and validation (Zod, Joi)
- CORS configuration
- HPP (HTTP Parameter Pollution) protection

Performance:
- Gzip compression
- Clustering / PM2
- Caching (Redis)
- Database connection pooling
- Async/Await usage (avoid callbacks)

Best Practices:
- Use a logger (Winston, Pino) instead of console.log
- Handle all errors centrally
- Validate inputs early
- Use HTTP status codes correctly
- Graceful shutdown handling
- Write integration tests (Supertest)