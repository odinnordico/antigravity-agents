---
trigger: model_decision
description: When it is about Vue 3 development with TypeScript, apply these rules.
---
You are an expert in Vue 3 development with TypeScript and Composition API.

Key Principles:
- Use Composition API with <script setup lang="ts">
- Type all props, emits, and refs properly
- Leverage Vue 3's built-in TypeScript support
- Use defineComponent for complex components
- Follow Vue 3 best practices

Component Setup:
- Use <script setup lang="ts"> for all components
- Define props with defineProps<Props>()
- Define emits with defineEmits<Emits>()
- Use ref<Type>() and reactive<Type>() with types
- Type computed properties properly

Props and Emits:
- Define prop interfaces with proper types
- Use withDefaults for default prop values
- Type emits with event payload types
- Use PropType for complex prop types
- Implement prop validation with types

Reactivity:
- Use ref<T>() for primitive values
- Use reactive<T>() for objects
- Type computed<T>() return values
- Use readonly<T>() for immutable state
- Type watch and watchEffect properly

Composables:
- Create typed composable functions
- Return typed objects from composables
- Use generic composables when appropriate
- Type composable parameters
- Export composable types

Vue Router:
- Type route params and query
- Use typed useRouter and useRoute
- Define RouteRecordRaw with types
- Type navigation guards
- Use typed router.push() and router.replace()

Pinia State Management:
- Define typed stores with defineStore
- Type state, getters, and actions
- Use typed store composition
- Type store plugins
- Use $patch with proper types

Template Refs:
- Type template refs with Ref<HTMLElement | null>
- Use component refs with proper types
- Type expose() for component methods
- Use defineExpose<T>() for typed exposure
- Type $refs properly

Provide/Inject:
- Use InjectionKey<T> for type-safe injection
- Type provide() and inject() properly
- Create typed injection keys
- Use default values with types
- Type app-level provides

Directives:
- Type custom directive hooks
- Use DirectiveBinding<T> for typed bindings
- Type directive modifiers
- Implement typed directive arguments
- Type directive lifecycle hooks

Components:
- Type component props with interfaces
- Use generic components when needed
- Type slots with defineSlots()
- Use typed component imports
- Type dynamic components

Vue 3 APIs:
- Type onMounted, onUnmounted, etc.
- Use typed lifecycle hooks
- Type nextTick properly
- Use typed teleport
- Type suspense components

Vite + Vue:
- Configure vite.config.ts for Vue
- Use @vitejs/plugin-vue with types
- Type Vite environment variables
- Use typed import.meta.env
- Configure path aliases with types

Testing:
- Use Vitest with Vue 3 and TypeScript
- Type component test utilities
- Use @vue/test-utils with types
- Type mock composables
- Implement typed test fixtures

Best Practices:
- Always use <script setup lang="ts">
- Type all component APIs
- Use discriminated unions for state
- Avoid 'any' type in Vue code
- Type all composables
- Use const assertions for constants
- Type all event handlers
- Implement proper error typing
- Use TypeScript with Vue DevTools
- Type all Pinia stores