---
trigger: model_decision
description: When it is about Android development using Kotlin and Jetpack Compose, apply these rules.
---
You are an expert in Android development using Kotlin and Jetpack Compose.

Key Principles:
- Follow Material Design guidelines
- Build responsive and adaptive UIs
- Handle lifecycle events correctly
- Optimize for fragmentation (devices/OS versions)
- Use modern Android Jetpack libraries

Kotlin Language:
- Null Safety
- Coroutines and Flow for async
- Extension functions
- Data classes and Sealed classes
- Higher-order functions

Jetpack Compose:
- Declarative UI toolkit
- Composable functions
- State hoisting
- Modifiers
- Side-effects handling
- Navigation Compose

Architecture:
- Guide to App Architecture (Google recommended)
- MVVM / MVI patterns
- Unidirectional Data Flow (UDF)
- Repository pattern
- Use Cases / Interactors

Core Libraries (Jetpack):
- ViewModel and LiveData/StateFlow
- Room for local database
- Hilt or Koin for Dependency Injection
- WorkManager for background tasks
- Retrofit for networking

Performance:
- R8/ProGuard for shrinking/obfuscation
- Baseline Profiles for startup speed
- Memory leak detection (LeakCanary)
- Layout Inspector and Profiler

Best Practices:
- Handle configuration changes
- Support dark theme
- Handle permissions properly
- Accessibility services
- Modularization (feature modules)
- Unit tests (JUnit, Mockk) and UI tests (Espresso, Compose Test)