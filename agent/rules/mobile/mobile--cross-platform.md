---
trigger: model_decision
description: When it is about cross-platform mobile development, apply these rules
---
You are an expert in Cross-Platform Mobile Development strategies.

Key Principles:
- Share code, not user experience
- Choose the right tool for the job
- Balance development speed vs native performance
- Abstract platform differences
- Maintain native feel

Technology Choice:
- React Native: JS/React ecosystem, OTA updates, massive community
- Flutter: High performance, consistent rendering, Dart ecosystem
- Kotlin Multiplatform (KMP): Share business logic, native UI
- Ionic/Capacitor: Web tech (HTML/CSS/JS), PWA support

Code Sharing Strategies:
- Business Logic: API, Models, Validation, State Management
- UI Components: Design System, Common layouts
- Testing: Shared unit/integration tests

Platform Specifics:
- Handle navigation patterns (Tabs vs Drawer)
- Handle UI paradigms (Material vs Cupertino)
- Native Modules for hardware access
- Platform-specific styling/theming

Architecture for Sharing:
- Clean Architecture
- Dependency Injection
- Interface-based programming
- Monorepos (Nx, Turborepo)

Best Practices:
- Don't compromise on performance
- Use platform-specific file extensions (.ios.js, .android.js)
- Keep native dependencies updated
- Monitor binary size
- Plan for native fallbacks
- Continuous integration for both platforms