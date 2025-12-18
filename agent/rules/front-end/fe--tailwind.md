---
trigger: model_decision
description: When it is about Tailwind CSS, apply these rules.
---
You are an expert in Tailwind CSS utility-first styling.

Key Principles:
- Utility-first workflow
- Design directly in markup
- Consistency via configuration
- Mobile-first responsive design
- Performance (purge unused styles)

Core Concepts:
- Utility Classes: flex, pt-4, text-center, etc.
- Responsive Prefixes: sm:, md:, lg:, xl:, 2xl:
- State Variants: hover:, focus:, active:, disabled:
- Dark Mode: dark: prefix
- Arbitrary Values: w-[32rem], bg-[#123456]

Configuration (tailwind.config.js):
- theme: Customize colors, spacing, fonts
- plugins: Add official or custom plugins
- content: Paths to files for tree-shaking
- darkMode: 'class' or 'media'

Directives (@layer):
- @tailwind base/components/utilities
- @apply: Extract utilities to custom classes (use sparingly)
- @layer: Organize custom CSS

Plugins:
- @tailwindcss/typography (prose)
- @tailwindcss/forms
- @tailwindcss/aspect-ratio
- @tailwindcss/container-queries

Best Practices:
- Don't use @apply just to make things look 'cleaner'
- Use components/partials for reusability instead of CSS classes
- Use the 'clsx' or 'tailwind-merge' libraries for dynamic classes
- Order classes consistently (prettier-plugin-tailwindcss)
- Customize the theme, don't fight it