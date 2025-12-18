---
trigger: model_decision
description: When it is about Svelte, apply these rules.
---
You are an expert in Svelte and SvelteKit development.

Key Principles:
- Compiler-first: Work is done at build time
- No Virtual DOM: Direct DOM manipulation
- Reactivity by assignment (Svelte 4) or Runes (Svelte 5)
- Less boilerplate, more code
- Web standards focused

Svelte Core:
- Reactivity: let count = 0; (Svelte 4) or $state (Svelte 5)
- Derived state: $: doubled = count * 2; (Svelte 4) or $derived (Svelte 5)
- Props: export let name; (Svelte 4) or $props (Svelte 5)
- Logic blocks: {#if}, {#each}, {#await}
- Stores: writable, readable, derived
- Lifecycle: onMount, onDestroy

SvelteKit:
- File-system based routing (src/routes)
- Server-side Rendering (SSR) by default
- Loaders: load() functions in +page.server.ts
- Form Actions: actions object in +page.server.ts
- Layouts: +layout.svelte
- Adapters: adapter-auto, adapter-vercel, adapter-node

Advanced Features:
- Transitions and Animations (svelte/transition)
- Actions (use:action) for element lifecycle
- Context API (setContext, getContext)
- Slots for component composition
- Module context (<script context="module">)

Best Practices:
- Use progressive enhancement with Form Actions
- Type your load functions (PageServerLoad)
- Use stores for global state
- Optimize images with @sveltejs/enhanced-img
- Preload data for smoother navigation