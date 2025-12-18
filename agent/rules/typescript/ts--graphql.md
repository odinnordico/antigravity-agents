---
trigger: model_decision
description: When it is about GraphQL development with TypeScript, apply these rules.
---
You are an expert in GraphQL development with TypeScript.

Key Principles:
- Use GraphQL Code Generator for type safety
- Type all resolvers and queries
- Implement schema-first or code-first approach
- Use typed GraphQL clients
- Follow GraphQL best practices

Code Generation:
- Use @graphql-codegen/cli for type generation
- Generate types from GraphQL schema
- Type queries, mutations, and subscriptions
- Generate typed React hooks
- Use typed fragments

Schema Definition (Code-First):
- Use TypeGraphQL for code-first schemas
- Type @ObjectType() classes
- Use @Field() with proper types
- Type @Resolver() classes
- Implement typed @Arg() parameters

Schema Definition (Schema-First):
- Define schema in .graphql files
- Generate TypeScript types from schema
- Type resolver implementations
- Use typed schema directives
- Implement typed custom scalars

Resolvers:
- Type resolver functions properly
- Use typed context objects
- Type parent/root arguments
- Implement typed field resolvers
- Type resolver return values

Queries and Mutations:
- Type query arguments
- Use typed input types
- Type mutation payloads
- Implement typed error handling
- Type query responses

Apollo Server:
- Type ApolloServer configuration
- Use typed context creation
- Type dataloaders properly
- Implement typed plugins
- Type server middleware

Apollo Client:
- Type useQuery hook with generics
- Use typed useMutation hook
- Type cache operations
- Implement typed Apollo Link
- Type client configuration

Subscriptions:
- Type subscription resolvers
- Use typed PubSub
- Type subscription payloads
- Implement typed WebSocket connections
- Type subscription filters

DataLoader:
- Type DataLoader instances
- Use typed batch functions
- Type cache key functions
- Implement typed loader context
- Type loader return values

Fragments:
- Generate typed fragments
- Use fragment composition with types
- Type fragment variables
- Implement typed fragment matchers
- Type inline fragments

Error Handling:
- Type GraphQL errors
- Use typed error extensions
- Implement typed error codes
- Type error formatting
- Use typed error handling middleware

Validation:
- Type input validation
- Use class-validator with GraphQL
- Type custom validators
- Implement typed validation rules
- Type validation errors

Testing:
- Type GraphQL test utilities
- Use typed mock resolvers
- Type test queries
- Implement typed integration tests
- Type Apollo Server testing

Performance:
- Type query complexity calculations
- Use typed persisted queries
- Type batching configurations
- Implement typed caching strategies
- Type query optimization

Best Practices:
- Always use GraphQL Code Generator
- Type all resolver arguments
- Use strict TypeScript mode
- Avoid 'any' in GraphQL code
- Type all context objects
- Use discriminated unions for responses
- Implement proper error typing
- Type all dataloaders
- Use const assertions for constants
- Type all GraphQL operations