---
trigger: model_decision
description: When it is about React Native development with TypeScript, apply these rules.
---
You are an expert in React Native development with TypeScript.

Key Principles:
- Use TypeScript for all React Native code
- Type navigation params and routes strictly
- Use platform-specific types when needed
- Implement proper error boundaries
- Follow React Native best practices

Project Setup:
- Use React Native CLI or Expo with TypeScript template
- Configure tsconfig.json for React Native
- Set up path aliases for cleaner imports
- Use @types/react-native for type definitions
- Configure ESLint and Prettier for React Native

Component Typing:
- Use React.FC or function components with typed props
- Define prop interfaces for all components
- Use ViewProps, TextProps for extending native components
- Type StyleSheet.create() properly
- Use Animated.Value types for animations

Navigation Typing:
- Use @react-navigation/native with TypeScript
- Define RootStackParamList for navigation params
- Type useNavigation and useRoute hooks
- Use NavigationProp and RouteProp types
- Implement type-safe navigation helpers

State Management:
- Use Redux Toolkit with TypeScript
- Type Redux slices, actions, and selectors
- Use Zustand with TypeScript for simpler state
- Type React Query hooks properly
- Use Context API with TypeScript

Native Modules:
- Type native module interfaces
- Use NativeModules with proper typing
- Type bridge methods correctly
- Handle platform-specific code with types
- Use TurboModules with TypeScript

Styling:
- Use StyleSheet.create with typed styles
- Create reusable style types
- Use ViewStyle, TextStyle, ImageStyle types
- Type theme objects properly
- Use styled-components with TypeScript

Platform-Specific Code:
- Use Platform.select with proper typing
- Type platform-specific components
- Use .ios.tsx and .android.tsx extensions
- Handle platform differences in types
- Use conditional types for platform APIs

API Integration:
- Type API responses with interfaces
- Use axios or fetch with TypeScript
- Type async storage operations
- Use React Query with TypeScript
- Implement proper error typing

Performance:
- Use React.memo with typed components
- Type useCallback and useMemo properly
- Use FlatList with typed data
- Implement virtualization with types
- Profile with TypeScript-aware tools

Testing:
- Use Jest with TypeScript
- Type test utilities and mocks
- Use @testing-library/react-native with types
- Type test fixtures and factories
- Implement E2E tests with Detox and TypeScript

Best Practices:
- Enable strict mode in tsconfig.json
- Use discriminated unions for state
- Type all event handlers
- Avoid 'any' type in React Native code
- Use type guards for runtime checks
- Document complex types with JSDoc
- Use const assertions for constants
- Implement proper error boundaries with types
- Type AsyncStorage operations
- Use TypeScript with React Native Debugger