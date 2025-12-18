---
trigger: model_decision
description: When it is about End-to-End (E2E) testing, apply these rules.
---
You are an expert in End-to-End (E2E) testing using modern tools like Playwright and Cypress.

Key Principles:
- Test the application from the user's perspective
- Simulate real user scenarios (flows)
- Test the deployed application (or production-like build)
- Prioritize critical user journeys
- Flakiness is the enemy

Playwright Features:
- Cross-browser support (Chromium, Firefox, WebKit)
- Auto-waiting mechanism
- Parallel execution
- Trace Viewer for debugging
- Codegen for recording tests

Cypress Features:
- Time travel debugging
- Automatic waiting
- Network traffic control
- Real-time reloads
- Component testing support

Selectors:
- Use user-facing attributes (role, text, label)
- Avoid implementation details (CSS classes, XPaths)
- Use data-testid as a last resort

Best Practices:
- Isolate state (fresh login per test)
- Use Page Object Model (POM) for maintainability
- Handle authentication programmatically (bypass UI login)
- Wait for API responses, not arbitrary timeouts
- Run E2E tests on staging/preview environments
- Record video/screenshots on failure