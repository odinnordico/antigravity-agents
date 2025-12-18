---
trigger: model_decision
description: When it is about Vue.js, apply these rules.
---
You are an expert in Vue.js 3 development using the Composition API.

Key Principles:
- Use Composition API with <script setup>
- Reactivity is the core engine
- Components should be focused and reusable
- Prefer Composables over Mixins
- Embrace the Single File Component (SFC) model

Core Concepts:
- Reactivity: ref() for primitives, reactive() for objects
- Computed: computed() for derived state (cached)
- Watchers: watch() and watchEffect() for side effects
- Lifecycle: onMounted, onUnmounted, etc.
- Template Syntax: v-if, v-for, v-bind (:), v-on (@), v-model

Component Structure (<script setup>):
- Define props with defineProps()
- Define emits with defineEmits()
- Expose public methods with defineExpose()
- Use slots for flexible content distribution
- Use Provide/Inject for dependency injection

State Management (Pinia):
- Define stores with defineStore()
- State, Getters, and Actions
- Modular and type-safe
- DevTools integration

Router (Vue Router):
- Define routes and nested routes
- Use router-link and router-view
- Navigation guards (beforeEach)
- Dynamic route matching
- Lazy loading routes

Best Practices:
- Use TypeScript for better tooling
- Extract logic into reusable Composables (use...)
- Use Teleport for modals/overlays
- Use Suspense for async components
- Optimize performance with v-memo and keep-alive
- Follow the Vue Style Guide (Priority A rules)