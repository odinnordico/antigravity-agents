---
trigger: model_decision
description: When it is about TypeScript monorepo architecture, apply these rules.
---
You are an expert in TypeScript monorepo architecture and management.

Key Principles:
- Use workspace protocols for package management
- Share types across packages efficiently
- Implement proper build orchestration
- Use TypeScript project references
- Maintain consistent tooling across packages

Monorepo Tools:
- Use Turborepo for build caching and orchestration
- Use pnpm workspaces for package management
- Use Nx for advanced monorepo features
- Use Lerna for publishing workflows
- Use Changesets for versioning

Project Structure:
```
monorepo/
├── apps/
│   ├── web/          # Next.js app
│   ├── mobile/       # React Native app
│   └── api/          # NestJS API
├── packages/
│   ├── ui/           # Shared UI components
│   ├── utils/        # Shared utilities
│   ├── types/        # Shared TypeScript types
│   └── config/       # Shared configs
├── turbo.json
├── pnpm-workspace.yaml
└── tsconfig.base.json
```

TypeScript Configuration:
- Use tsconfig.base.json for shared config
- Implement TypeScript project references
- Use path aliases across packages
- Configure composite: true for references
- Use incremental builds

Shared Types Package:
- Create dedicated @repo/types package
- Export common interfaces and types
- Use declaration files (.d.ts)
- Implement type versioning
- Type API contracts

Package Dependencies:
- Use workspace protocol (workspace:*)
- Type internal package imports
- Implement proper peer dependencies
- Use typed package exports
- Configure package.json types field

Build Orchestration:
- Configure turbo.json for task dependencies
- Use typed build scripts
- Implement incremental builds
- Type build outputs
- Use remote caching

Code Sharing:
- Share UI components with proper types
- Export typed utility functions
- Share configuration with types
- Implement typed design tokens
- Share validation schemas

Linting and Formatting:
- Use shared ESLint config
- Implement typed lint rules
- Share Prettier configuration
- Use typed lint-staged configs
- Implement pre-commit hooks

Testing:
- Share test utilities with types
- Use typed test configurations
- Implement shared test setup
- Type test fixtures
- Share mock factories

Versioning and Publishing:
- Use Changesets for versioning
- Type changelog generation
- Implement typed release scripts
- Use semantic versioning
- Type package publishing configs

Development Workflow:
- Use typed dev scripts
- Implement hot module reloading
- Type watch mode configurations
- Use typed environment variables
- Implement development tools

CI/CD:
- Type CI pipeline configurations
- Use affected commands (Nx/Turborepo)
- Implement typed deployment scripts
- Type build artifacts
- Use matrix builds for packages

Performance:
- Use build caching effectively
- Implement parallel builds
- Type build optimization configs
- Use incremental type checking
- Optimize package dependencies

Best Practices:
- Use TypeScript project references
- Implement strict tsconfig across all packages
- Share types through dedicated package
- Use consistent naming conventions
- Type all internal package APIs
- Implement proper dependency management
- Use workspace protocols
- Type all build configurations
- Document package relationships
- Use typed scripts in package.json