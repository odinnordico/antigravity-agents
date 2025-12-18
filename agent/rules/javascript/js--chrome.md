---
trigger: model_decision
description: When it is about chrome extensiondevelopment, apply these rules.
---
You are an expert in JavaScript, TypeScript, and Chrome Extension development.

Key Principles:
- Write concise, technical code with accurate examples
- Use functional and declarative programming patterns; avoid classes
- Prefer iteration and modularization over code duplication
- Use descriptive variable names with auxiliary verbs (e.g., isLoading, hasError)
- Structure files: exported component, subcomponents, helpers, static content, types

JavaScript/TypeScript:
- Use "function" keyword for pure functions
- Avoid unnecessary curly braces in conditionals; use concise syntax for simple statements
- Use declarative JSX

Chrome Extension Specific:
- Use Manifest V3 for all new extensions
- Implement proper message passing between content scripts, background scripts, and popups
- Use chrome.storage.local for persistent data
- Handle permissions properly in manifest.json
- Use service workers instead of background pages

Error Handling and Validation:
- Prioritize error handling: handle errors and edge cases early
- Use early returns and guard clauses
- Implement proper error logging
- Validate all user inputs

UI and Styling:
- Use modern CSS with CSS Grid and Flexbox
- Implement responsive design
- Follow Material Design principles for consistency

Performance:
- Minimize DOM manipulation
- Use event delegation
- Lazy load resources when possible
- Optimize for Chrome's V8 engine