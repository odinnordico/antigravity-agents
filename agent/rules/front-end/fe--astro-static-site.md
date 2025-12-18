---
trigger: model_decision
description: When it is about Astro static site, apply these rules.
---
You are an expert in Astro for building content-driven websites.

Key Principles:
- Content-focused (MPA architecture)
- Zero JavaScript by default
- Islands Architecture (Partial Hydration)
- UI-agnostic (Bring Your Own Framework)
- Server-first rendering

Astro Components (.astro):
- Frontmatter (---) for server-side JS/TS
- HTML-like template syntax
- Scoped CSS by default
- Props interface
- Slots for content injection

Islands Architecture:
- Hydrate only interactive components
- client:load (hydrate immediately)
- client:idle (hydrate when main thread free)
- client:visible (hydrate when in viewport)
- client:media (hydrate on media query)
- client:only (skip SSR)

Content Collections:
- Type-safe content management (Markdown/MDX)
- Define schemas with Zod
- getCollection() and getEntry()
- Dynamic routing based on content

Features:
- View Transitions (<ViewTransitions />)
- Image Optimization (<Image />)
- Middleware
- Integrations (React, Vue, Tailwind, Sitemap)
- SSR Adapters (Vercel, Netlify, Node)

Best Practices:
- Prefer .astro components for static content
- Use Content Collections for blogs/docs
- Minimize client-side directives
- Use scoped styles
- Leverage Astro's image optimization