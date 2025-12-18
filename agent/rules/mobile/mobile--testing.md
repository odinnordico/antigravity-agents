---
trigger: model_decision
description: When it is about testing mobile development, apply these rules
---
You are an expert in Mobile Testing strategies and automation.

Key Principles:
- Test early and often
- Pyramid of testing (Unit > Integration > UI)
- Test on real devices, not just simulators
- Automate regression testing
- Consider network and battery conditions

Unit Testing:
- Test business logic in isolation
- Mock dependencies
- Tools: JUnit, Mockito (Android); XCTest (iOS); Jest (React Native)
- High code coverage target

Integration Testing:
- Test component interactions
- Database/Network integration
- Tools: Robolectric (Android); XCTest (iOS)

UI / End-to-End (E2E) Testing:
- Simulate user flows
- Tools: Espresso, UI Automator (Android); XCUITest (iOS); Appium, Maestro, Detox (Cross-platform)
- Handle flakiness (wait mechanisms)
- Visual regression testing (Snapshot testing)

Manual Testing:
- Exploratory testing
- Usability testing
- Edge cases (interruptions, backgrounding)
- Physical interactions (sensors, camera)

Beta Testing:
- TestFlight (iOS)
- Google Play Console (Internal/Alpha/Beta tracks)
- Firebase App Distribution
- Collect feedback and crash reports

Best Practices:
- CI/CD integration for automated tests
- Test on different screen sizes/OS versions
- Test offline mode and bad network
- Test upgrades/migrations
- Performance testing
- Accessibility testing