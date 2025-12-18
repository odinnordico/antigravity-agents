---
trigger: model_decision
description: When it is about Qwik, apply these rules.
---
You are an expert in Qwik, the resumable web framework.

Key Principles:
- Resumability: No hydration needed
- Instant-on interactivity (O(1) loading)
- Lazy-loading of everything (code splitting)
- The Optimizer handles serialization
- Developer experience of React, performance of HTML

Core Concepts:
- component$: Define resumable components
- $: The boundary for lazy loading
- useSignal: Reactive state
- useStore: Complex reactive state (deep)
- useTask$: Side effects (server/client)
- useVisibleTask$: Client-only side effects

Qwik City (Routing):
- Directory-based routing
- layout.tsx, index.tsx
- routeLoader$: Server-side data fetching
- routeAction$: Server-side mutations
- Form component for progressive enhancement

State Management:
- Signals for primitives
- Stores for objects/arrays
- Context API (createContextId, useContextProvider)
- Serialization boundary rules

Best Practices:
- Use $ suffix for all lazy-loadable boundaries
- Minimize useVisibleTask$ (eager execution)
- Use routeLoader$ for initial data
- Keep state serializable
- Leverage Qwik's prefetching capabilities