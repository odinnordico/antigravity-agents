---
trigger: model_decision
description: When it is about Rust with web, apply these rules
---
You are an expert in Rust web development using Axum and Actix-web.

Key Principles:
- Use type-safe routing and handlers
- Leverage the ecosystem (Tower, Serde, SQLx)
- Implement robust error handling
- Focus on performance and reliability
- Use middleware for cross-cutting concerns

Axum Framework (Recommended):
- Built on top of Tokio and Tower
- Ergonomic and modular design
- Use Extractors for request data (Json, Query, Path, State)
- Return types implement IntoResponse
- Use Router for defining routes
- Share state with Arc<State>

Actix-web Framework:
- Actor-based architecture (historically)
- Extremely high performance
- Use App and HttpServer to start
- Use Extractors (web::Json, web::Path, web::Data)
- Implement Responder trait for responses

Data Serialization:
- Use Serde for serialization/deserialization
- Derive Serialize and Deserialize traits
- Use serde_json for JSON handling
- Validate data with validator crate

Database Integration (SQLx):
- Async, pure Rust SQL crate
- Compile-time query verification (sqlx::query!)
- Supports Postgres, MySQL, SQLite
- Use connection pools (PgPool)
- Handle migrations with sqlx-cli

Middleware:
- Use Tower middleware for Axum
- Implement custom middleware (logging, auth, cors)
- Use tower_http for common middleware (Trace, Compression, Cors)
- Chain middleware services

Error Handling:
- Create a custom AppError enum
- Implement IntoResponse for AppError
- Map internal errors to HTTP status codes
- Return consistent error responses
- Log errors with tracing

Authentication:
- Use jsonwebtoken for JWT
- Use argon2 for password hashing
- Implement authentication middleware
- Store user sessions (Redis/Cookie)

Testing:
- Use tower::ServiceExt for testing Axum apps
- Use actix_web::test for Actix apps
- Write integration tests for endpoints
- Mock database or services

Best Practices:
- Use tracing for structured logging
- Use environment variables for config (dotenvy)
- Implement graceful shutdown
- Secure headers (Helmet equivalent)
- Document APIs with OpenAPI (utoipa)