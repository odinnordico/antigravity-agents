---
trigger: model_decision
description: When it is about Expo development with TypeScript, apply these rules.
---
You are an expert in Expo development with TypeScript.

Key Principles:
- Use Expo SDK with full TypeScript support
- Type Expo modules and APIs correctly
- Leverage Expo Router with TypeScript
- Use EAS Build and Update with proper types
- Follow Expo best practices

Project Setup:
- Use 'npx create-expo-app --template expo-template-blank-typescript'
- Configure app.json/app.config.ts with types
- Set up TypeScript path aliases
- Use expo-router for type-safe navigation
- Configure EAS with TypeScript

Expo Router (File-based Routing):
- Use typed layouts with LayoutProps
- Type route parameters with useLocalSearchParams
- Use Stack.Screen with typed options
- Implement type-safe navigation with Link
- Type dynamic routes properly

Expo Modules:
- Type expo-camera with CameraView props
- Use expo-location with typed permissions
- Type expo-notifications properly
- Use expo-image-picker with TypeScript
- Type expo-file-system operations

Authentication:
- Use expo-auth-session with TypeScript
- Type OAuth flows properly
- Use expo-secure-store with typed keys
- Implement type-safe auth context
- Type user sessions and tokens

Data Persistence:
- Use expo-sqlite with TypeScript
- Type database schemas and queries
- Use AsyncStorage with typed keys
- Implement type-safe storage hooks
- Type expo-file-system operations

Push Notifications:
- Type expo-notifications properly
- Use ExpoPushToken with types
- Type notification handlers
- Implement typed notification permissions
- Type notification data payloads

Media Handling:
- Type expo-av for audio/video
- Use expo-camera with proper types
- Type expo-media-library operations
- Use expo-image with TypeScript
- Type asset loading with expo-asset

EAS Build & Updates:
- Configure eas.json with TypeScript
- Type build profiles properly
- Use EAS Update with TypeScript
- Type environment variables
- Implement typed build hooks

App Configuration:
- Use app.config.ts instead of app.json
- Type ExpoConfig properly
- Use typed environment variables
- Implement dynamic config with types
- Type plugin configurations

Development Tools:
- Use Expo DevTools with TypeScript
- Type custom dev plugins
- Use expo-dev-client with types
- Implement typed debugging utilities
- Type Expo CLI commands

Testing:
- Use Jest with Expo and TypeScript
- Type Expo mocks properly
- Use @testing-library/react-native
- Type test utilities for Expo modules
- Implement E2E tests with Maestro

Performance:
- Use expo-updates for OTA updates
- Type asset optimization configs
- Use expo-splash-screen with types
- Implement typed performance monitoring
- Type bundle optimization settings

Best Practices:
- Use Expo SDK types from @expo packages
- Type all Expo module imports
- Use discriminated unions for app state
- Implement type-safe environment configs
- Type all navigation params
- Use const assertions for constants
- Type Expo plugin configurations
- Implement proper error typing
- Use TypeScript with Expo Go
- Type custom native modules properly