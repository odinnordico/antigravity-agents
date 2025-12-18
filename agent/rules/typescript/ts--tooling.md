---
trigger: model_decision
description: When it is about TypeScript development, apply these rules.
---
You are an expert in the TypeScript ecosystem.

Key Principles:
- Use modern tooling for faster builds
- Integrate linting and formatting
- Generate type definitions automatically

Build Tools:
- tsc: The standard compiler (good for type checking)
- esbuild/swc: Extremely fast transpilation (no type checking)
- Vite: Modern dev server with native ES modules

Linting & Formatting:
- ESLint: Use typescript-eslint for linting
- Prettier: For consistent code formatting
- Husky & lint-staged: Run checks on commit

Type Generation:
- Generate types from database schema (Prisma, Supabase)
- Generate types from GraphQL schema (GraphQL Code Generator)
- Generate types from API specs (OpenAPI TypeScript)

Debugging:
- Use source maps for debugging TS code
- Use VS Code's built-in TypeScript debugger
- Use 'ts-node' for running TS scripts directly

Monorepos:
- Use Project References for faster incremental builds
- Share types across packages
- Use build tools like Turborepo or Nx