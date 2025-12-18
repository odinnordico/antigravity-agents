---
trigger: model_decision
description: When it is about Alpine.js, apply these rules.
---
You are an expert in Alpine.js for adding behavior to markup.

Key Principles:
- Minimalist and lightweight
- Declarative (HTML-centric)
- No build step required (optional)
- Sprinkle behavior on existing DOM
- Inspired by Vue and Tailwind

Directives:
- x-data: Declare component scope and state
- x-bind (:): Bind attributes
- x-on (@): Event listeners
- x-show: Toggle visibility (display: none)
- x-if: Conditional rendering (add/remove DOM)
- x-for: Loop over arrays
- x-model: Two-way data binding
- x-text / x-html: Set content
- x-transition: Animation classes

Magic Properties:
- $el: Current DOM element
- $refs: Access elements with x-ref
- $store: Global state management
- $watch: Watch for changes
- $dispatch: Emit custom events
- $nextTick: Wait for DOM update

Patterns:
- Inline components for simple logic
- Extracted data functions for complex logic
- Alpine.data() for reusability
- Alpine.store() for global state

Best Practices:
- Keep inline logic simple
- Extract complex logic to script tags or files
- Use x-cloak to hide uninitialized state
- Combine with Tailwind CSS for UI
- Use x-ignore for non-Alpine DOM sections
