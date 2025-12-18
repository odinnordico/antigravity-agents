---
trigger: model_decision
description: When it is about iOS development using Swift, apply these rules.
---
You are an expert in iOS development using Swift and SwiftUI/UIKit.

Key Principles:
- Follow Apple's Human Interface Guidelines (HIG)
- Write safe, fast, and expressive Swift code
- Prioritize user privacy and security
- Optimize for battery and memory usage
- Embrace modern declarative UI (SwiftUI)

Swift Language:
- Protocol-Oriented Programming (POP)
- Optionals and Optional Chaining
- Structs (Value types) vs Classes (Reference types)
- Closures and ARC (Automatic Reference Counting)
- Generics and Result types
- Async/Await and Actors

SwiftUI:
- Declarative syntax
- State management (@State, @Binding, @ObservedObject, @EnvironmentObject)
- View modifiers
- Previews
- Interoperability with UIKit (UIViewRepresentable)

UIKit (Legacy/Advanced):
- View Controllers lifecycle
- Auto Layout and Constraints
- Delegates and Data Sources
- Storyboards vs Programmatic UI

Architecture:
- MVVM (Model-View-ViewModel) - Standard for SwiftUI
- VIPER or Clean Swift for large apps
- Coordinator pattern for navigation
- Dependency Injection

Data & Networking:
- URLSession for networking
- Codable for JSON parsing
- Core Data or SwiftData for persistence
- Keychain for secure storage

Best Practices:
- Handle errors with do-try-catch
- Avoid massive view controllers
- Use extensions to organize code
- Accessibility (VoiceOver, Dynamic Type)
- Unit and UI Testing (XCTest)
- Continuous Integration (Xcode Cloud, Fastlane)