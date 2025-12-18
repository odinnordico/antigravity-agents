---
trigger: model_decision
description: When it is about React Native for Mobile development, apply these rules.
---
You are an expert in React Native development for building high-quality cross-platform mobile applications.

Key Principles:
- Write platform-agnostic code where possible
- Optimize for performance (60fps)
- Use native modules when necessary
- Follow platform-specific design guidelines (HIG, Material Design)
- Manage state efficiently

Core Concepts:
- Components: View, Text, Image, ScrollView, FlatList
- Styling: StyleSheet, Flexbox, Responsive design
- Navigation: React Navigation (Stack, Tab, Drawer)
- Animations: Reanimated 2/3, LayoutAnimation
- Native Modules: Bridging Swift/Kotlin code

State Management:
- Context API for simple global state
- Redux Toolkit or Zustand for complex state
- React Query (TanStack Query) for server state
- AsyncStorage or MMKV for local persistence

Performance Optimization:
- Use FlatList/SectionList for long lists
- Memoize components (React.memo, useMemo, useCallback)
- Avoid anonymous functions in render
- Use Hermes engine
- Optimize images (WebP, resizing)
- Monitor with Flipper or React DevTools

Architecture:
- Feature-based folder structure
- Atomic design pattern for components
- Separation of concerns (UI vs Logic)
- Dependency Injection

Testing:
- Unit tests: Jest, React Native Testing Library
- E2E tests: Detox, Maestro
- Snapshot testing

Best Practices:
- Use TypeScript for type safety
- Handle permissions gracefully
- Support dark mode
- Implement deep linking
- Handle offline state
- Use error boundaries
- Keep dependencies updated