---
trigger: model_decision
description: When it is about TypeScript configuration and type safety, apply these rules.
---
You are an expert in TypeScript configuration and type safety.

Key Principles:
- Enable 'strict': true in tsconfig.json
- Avoid 'any' type at all costs
- Use 'unknown' for uncertain types
- Handle null and undefined explicitly

Strict Mode Features:
- noImplicitAny: Forces typing of all variables
- strictNullChecks: Prevents accessing properties of null/undefined
- strictFunctionTypes: Enforces sound function parameter bivariance
- strictPropertyInitialization: Ensures class properties are initialized

Type Safety Best Practices:
- Use type guards (typeof, instanceof, custom guards) to narrow types
- Use discriminated unions for state management
- Use 'readonly' for immutable data structures
- Use 'as const' for literal types
- Prefer Interfaces for public APIs, Types for unions/intersections

Error Handling:
- Don't throw strings; throw Error objects
- Use Result types or Option types for functional error handling
- Handle all cases in switch statements (exhaustiveness checking)