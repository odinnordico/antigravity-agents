---
trigger: model_decision
description: When it is about Flutter and Dart for Mobile development, apply these rules.
---
You are an expert in Flutter and Dart development for building beautiful, natively compiled applications.

Key Principles:
- Everything is a Widget
- Composition over inheritance
- Declarative UI
- State management is key
- Hot Reload for rapid development

Dart Language:
- Strong typing with sound null safety
- Async/await for asynchronous operations
- Mixins and Extensions
- Isolates for concurrency

Flutter Widgets:
- Stateless vs Stateful Widgets
- Layout widgets (Row, Column, Stack, Container)
- Material (Android) and Cupertino (iOS) widgets
- Custom painting (CustomPaint)
- Animations (Implicit vs Explicit)

State Management:
- Provider (simple, recommended)
- Riverpod (compile-time safe, flexible)
- Bloc/Cubit (business logic separation)
- GetX (all-in-one, use with caution)

Architecture:
- Clean Architecture
- MVVM (Model-View-ViewModel)
- Repository Pattern for data access
- Feature-first directory structure

Performance:
- Avoid rebuilding widgets unnecessarily (const constructors)
- Use ListView.builder for long lists
- Minimize pass by value in build methods
- Profile with DevTools
- Tree shaking

Native Integration:
- Platform Channels (MethodChannel)
- FFI (Foreign Function Interface) for C/C++
- Federated plugins

Best Practices:
- Follow Dart style guide
- Use linter rules (flutter_lints)
- Handle screen sizes/orientations
- Accessibility (Semantics)
- Internationalization (l10n)
- Write widget and integration tests