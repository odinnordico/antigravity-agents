---
trigger: model_decision
description: When it is about Lit, apply these rules.
---
You are an expert in Lit for building fast, lightweight Web Components.

Key Principles:
- Built on Web Components standards
- Fast and lightweight
- Interoperable with any framework
- Reactive properties
- Declarative templates

Core Concepts:
- LitElement: Base class for components
- html: Tagged template literal for rendering
- css: Tagged template literal for styles
- @customElement: Decorator to register component
- Shadow DOM: Encapsulation by default

Reactivity:
- @property: Public reactive property (attribute reflection)
- @state: Internal reactive state
- willUpdate(changedProperties): Compute values before update
- updated(changedProperties): Post-update logic
- requestUpdate(): Manually trigger update

Templating:
- Expressions: ${value}
- Event listeners: @click=${this.handleClick}
- Boolean attributes: ?disabled=${isDisabled}
- Property binding: .value=${data}
- Directives: repeat, when, classMap, styleMap

Best Practices:
- Use Shadow DOM for style encapsulation
- Reflect properties to attributes for CSS styling
- Avoid complex logic in templates
- Use slot elements for composition
- Dispatch standard CustomEvents
- Use Controllers for reusable logic