---
trigger: model_decision
description: When it is about Next.js state management, apply these rules.
---
You are an expert in Next.js state management patterns and libraries.

Key Principles:
- Use Server Components for server state
- Use React Query for async state
- Use Zustand for client state
- Minimize client-side state
- Leverage URL for shareable state

Server State:
- Fetch data in Server Components
- Use Next.js cache for data fetching
- Implement ISR for dynamic data
- Use Server Actions for mutations
- Revalidate cache after updates

React Query (TanStack Query):
- Install @tanstack/react-query
- Set up QueryClientProvider
- Use useQuery for data fetching
- Use useMutation for updates
- Implement optimistic updates
- Configure stale time and cache time

Zustand for Client State:
- Install zustand
- Create stores with create()
- Use hooks in Client Components
- Implement persist middleware
- Use devtools for debugging
- Keep stores small and focused

URL State:
- Use searchParams for filters
- Use useRouter for navigation
- Implement shareable URLs
- Sync URL with component state
- Use nuqs for type-safe params

Form State:
- Use react-hook-form
- Implement Zod validation
- Use Server Actions for submission
- Implement optimistic updates
- Handle form errors properly

Context API:
- Use for theme, locale, auth
- Avoid for frequently changing data
- Use with Server Components carefully
- Implement proper memoization
- Split contexts by concern

Local Storage:
- Use for user preferences
- Implement with Zustand persist
- Handle SSR hydration issues
- Use localStorage only in Client Components
- Implement fallbacks for SSR

Optimistic Updates:
- Update UI immediately
- Rollback on error
- Use React Query mutations
- Implement with Server Actions
- Provide user feedback

Cache Management:
- Use Next.js cache for server data
- Use React Query cache for client data
- Implement cache invalidation
- Use revalidateTag and revalidatePath
- Configure cache strategies per query

Global State Patterns:
- Minimize global state
- Use composition over global state
- Lift state only when necessary
- Use URL for shareable state
- Keep state close to where it's used

State Synchronization:
- Sync server and client state
- Use React Query for server state
- Implement real-time updates
- Handle stale data properly
- Use optimistic updates

Performance:
- Use React.memo for expensive components
- Implement proper memoization
- Use useCallback and useMemo
- Avoid unnecessary re-renders
- Use Zustand's shallow comparison

Testing State:
- Test state logic separately
- Mock state in component tests
- Test state updates
- Test optimistic updates
- Test error states

Best Practices:
- Prefer Server Components for data
- Use React Query for async state
- Keep client state minimal
- Use URL for shareable state
- Implement proper error handling
- Use TypeScript for type safety
- Test state management logic
- Document state structure
- Use devtools for debugging
- Monitor state performance