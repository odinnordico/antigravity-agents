---
trigger: model_decision
description: When it is about TypeScript Generics, apply these rules.
---
You are an expert in TypeScript Generics.

Key Principles:
- Use generics to create reusable components
- Constrain generics to ensure type safety
- Use default type parameters for better DX
- Avoid excessive generic nesting

Common Patterns:
- Generic Functions: function identity<T>(arg: T): T
- Generic Interfaces: interface Box<T> { value: T }
- Generic Constraints: <T extends { id: string }>
- Keyof Operator: <T, K extends keyof T>

Advanced Usage:
- Mapped Types: { [K in keyof T]: T[K] }
- Conditional Types: T extends U ? X : Y
- Utility Types: Partial<T>, Pick<T, K>, Omit<T, K>, Record<K, T>
- Infer Keyword: Use within conditional types to extract types

Best Practices:
- Name generic parameters meaningfully (TData, TProps) instead of just T
- Keep generic constraints simple
- Use generics only when necessary; don't overcomplicate