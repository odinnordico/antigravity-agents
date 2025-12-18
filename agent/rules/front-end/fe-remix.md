---
trigger: model_decision
description: When it is about Remix, apply these rules.
---
You are an expert in Remix, the full-stack web framework.

Key Principles:
- Embrace Web Standards (HTTP, HTML, Forms)
- Progressive Enhancement
- Nested Routing matches UI hierarchy
- Server/Client bridge is seamless
- No 'loading' spinners (Optimistic UI)

Data Loading (Server):
- loader function (GET requests)
- Access URL params and request object
- Return JSON or Response
- useLoaderData hook in component
- Parallel data fetching by default

Data Mutation (Server):
- action function (POST/PUT/DELETE)
- Handle form submissions
- Redirect or return validation errors
- Automatic revalidation of loaders

Components & UI:
- <Form> component (prevents default, uses fetch)
- useNavigation for pending states
- useActionData for form errors
- <Outlet> for nested routes
- ErrorBoundary and HydraFallback

Routing:
- File-system routing (app/routes)
- Dynamic segments ($id)
- Layout routes
- Resource routes (no UI)

Best Practices:
- Use standard HTML forms where possible
- Validate data with Zod on the server
- Handle errors granularly with ErrorBoundaries
- Use cache-control headers
- Implement Optimistic UI for instant feedback