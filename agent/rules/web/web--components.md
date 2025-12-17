---
trigger: model_decision
description: When it is about web development for custom elements and web components, apply these rules.
---
You are an expert in Web Components and Custom Elements.

Key Principles:
- Use Web Components for reusable UI
- Implement proper encapsulation
- Follow web standards
- Use Shadow DOM for style isolation
- Implement progressive enhancement

Custom Elements:
- Define custom elements with customElements.define()
- Extend HTMLElement class
- Use lifecycle callbacks
- Implement proper naming (kebab-case)
- Use autonomous custom elements
- Extend built-in elements when appropriate

Lifecycle Callbacks:
- connectedCallback: element added to DOM
- disconnectedCallback: element removed from DOM
- attributeChangedCallback: attribute changed
- adoptedCallback: element moved to new document
- Use observedAttributes for watched attributes

Shadow DOM:
- Use attachShadow for encapsulation
- Implement open or closed shadow roots
- Use slots for content projection
- Style with :host and :host-context
- Use ::slotted for slotted content
- Implement proper CSS encapsulation

Templates:
- Use <template> for reusable markup
- Clone template content
- Use with Shadow DOM
- Implement efficient rendering
- Cache template references

Slots:
- Use <slot> for content projection
- Implement named slots
- Use default slot content
- Handle slotchange events
- Implement slot fallbacks
- Use assignedNodes() and assignedElements()

Attributes and Properties:
- Reflect properties to attributes
- Use getters and setters
- Implement proper type conversion
- Use attributeChangedCallback
- Follow HTML attribute conventions
- Implement boolean attributes properly

Events:
- Dispatch custom events
- Use composed: true for cross-boundary events
- Implement proper event naming
- Use detail for event data
- Handle event bubbling
- Implement event delegation

Styling:
- Use :host for component styles
- Use CSS custom properties for theming
- Implement :host-context for context styles
- Use ::part for styling internal elements
- Implement CSS Shadow Parts
- Use constructable stylesheets

Accessibility:
- Use ARIA attributes
- Implement keyboard navigation
- Use semantic HTML in shadow DOM
- Provide proper focus management
- Implement accessible custom controls
- Test with screen readers

Performance:
- Use lazy registration
- Implement efficient rendering
- Use requestAnimationFrame for updates
- Minimize DOM operations
- Use event delegation
- Implement proper cleanup

Libraries and Tools:
- Lit for declarative templates
- Stencil for component compilation
- Polymer for polyfills and utilities
- Use web component analyzers
- Implement proper build tooling

Testing:
- Test custom elements in isolation
- Test lifecycle callbacks
- Test attribute/property sync
- Test event dispatching
- Test accessibility
- Use web component testing libraries

Best Practices:
- Follow HTML naming conventions
- Implement proper encapsulation
- Use Shadow DOM for style isolation
- Provide clear API documentation
- Implement progressive enhancement
- Use semantic HTML
- Test across browsers
- Implement proper error handling
- Use TypeScript for type safety
- Publish to npm for reusability