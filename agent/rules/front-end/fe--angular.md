---
trigger: model_decision
description: When it is about Angular, apply these rules.
---
You are an expert in modern Angular development (v17+).

Key Principles:
- TypeScript-first development
- Component-based architecture
- Dependency Injection is core
- Use Signals for fine-grained reactivity
- Prefer Standalone Components over NgModules

Core Concepts:
- Standalone Components: imports: [] in @Component
- Signals: signal(), computed(), effect()
- Control Flow: @if, @for, @switch (new syntax)
- Input/Output: input(), output() (signal-based)
- Dependency Injection: inject() function or constructor

RxJS & Observables:
- Use Observables for streams of data
- Pipeable operators (map, filter, switchMap)
- AsyncPipe for subscription management in templates
- DestroyRef for cleanup
- Interop with Signals (toSignal, toObservable)

Routing:
- Define routes in app.routes.ts
- Lazy loading components (loadComponent)
- Route guards (CanActivateFn)
- Resolvers for data fetching

Forms:
- Reactive Forms (FormControl, FormGroup)
- FormBuilder for cleaner syntax
- Typed Forms
- Custom Validators

Best Practices:
- Use OnPush change detection strategy
- Avoid logic in templates
- Use Smart/Dumb component pattern
- Use Services for business logic and state
- Strict mode in TypeScript
- Use Angular CLI for generation and building