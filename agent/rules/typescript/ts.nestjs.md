---
trigger: model_decision
description: When it is about NestJS development with TypeScript, apply these rules.
---
You are an expert in NestJS backend development with TypeScript.

Key Principles:
- Use decorators with proper TypeScript types
- Leverage dependency injection with types
- Type all DTOs and entities
- Use class-validator for runtime validation
- Follow NestJS best practices

Project Structure:
- Use NestJS CLI for scaffolding
- Organize by feature modules
- Type all module imports/exports
- Use barrel exports with types
- Implement typed configuration

Controllers:
- Type @Controller() routes
- Use typed @Param(), @Query(), @Body()
- Type response objects
- Implement typed exception filters
- Type request/response interceptors

Services and Providers:
- Type @Injectable() services
- Use constructor injection with types
- Type service methods properly
- Implement typed business logic
- Type async operations

DTOs and Validation:
- Use class-validator decorators
- Type DTO classes properly
- Implement validation pipes with types
- Use class-transformer for typing
- Type validation error responses

Entities and Database:
- Type TypeORM entities properly
- Use @Entity() with typed columns
- Type repository methods
- Implement typed query builders
- Type database relations

Middleware and Guards:
- Type custom middleware
- Implement typed auth guards
- Type CanActivate interface
- Use typed ExecutionContext
- Type guard return values

Interceptors and Pipes:
- Type NestInterceptor interface
- Implement typed transform pipes
- Type CallHandler and Observable
- Use typed validation pipes
- Type interceptor responses

Exception Filters:
- Type custom exception filters
- Implement typed error responses
- Use HttpException with types
- Type ArgumentsHost properly
- Type exception handling logic

GraphQL:
- Type @Resolver() classes
- Use @Query() and @Mutation() with types
- Type GraphQL schemas
- Implement typed resolvers
- Type GraphQL context

WebSockets:
- Type @WebSocketGateway()
- Use typed @SubscribeMessage()
- Type WebSocket events
- Implement typed socket handlers
- Type WebSocket middleware

Microservices:
- Type microservice patterns
- Use typed message patterns
- Type client proxies
- Implement typed event handlers
- Type microservice responses

Testing:
- Type Test module configurations
- Use typed mocks and stubs
- Type test fixtures
- Implement typed E2E tests
- Type supertest requests

Configuration:
- Use @nestjs/config with types
- Type ConfigService properly
- Implement typed environment variables
- Type configuration namespaces
- Use validation schemas with types

Best Practices:
- Enable strict TypeScript mode
- Type all decorators properly
- Use DTOs for all inputs/outputs
- Avoid 'any' in NestJS code
- Type all database operations
- Use discriminated unions for responses
- Implement proper error typing
- Type all async operations
- Use const assertions for constants
- Type all dependency injections