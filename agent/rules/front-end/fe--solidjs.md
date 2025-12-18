---
trigger: model_decision
description: When it is about SolidJS, apply these rules.
---
You are an expert in Solid.js and fine-grained reactivity.

Key Principles:
- Fine-grained reactivity (Signals)
- No Virtual DOM
- Components run once (setup phase)
- JSX for templating
- High performance and small bundle size

Core Primitives:
- Signals: createSignal() -> [get, set]
- Effects: createEffect() for side effects
- Memos: createMemo() for derived cached values
- Resources: createResource() for async data
- Stores: createStore() for nested reactivity

Control Flow:
- Use components, not array maps/ternaries
- <Show when={...}> for conditionals
- <For each={...}> for keyed lists
- <Index each={...}> for non-keyed lists
- <Switch> and <Match>

Lifecycle & Context:
- onMount, onCleanup
- createContext, useContext
- batch() for multiple updates
- untrack() to read without subscribing

Solid Router:
- <Router>, <Route>, <A>
- Nested routing
- Data functions (Loaders)
- Route actions

Best Practices:
- Don't destructure props (loses reactivity)
- Access signals as functions (count())
- Use derived signals for computed values
- Use Stores for complex state
- Optimize with lazy loading (lazy())
- Understand the tracking scope