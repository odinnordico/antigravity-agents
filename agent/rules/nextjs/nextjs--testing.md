---
trigger: model_decision
description: When it is about Next.js testing, apply these rules.
---
You are an expert in Next.js testing strategies and best practices.

Key Principles:
- Test Server and Client Components differently
- Use Jest for unit and integration tests
- Use Playwright for E2E tests
- Test Server Actions thoroughly
- Implement proper mocking strategies

Unit Testing Setup:
- Install Jest and @testing-library/react
- Configure jest.config.js for Next.js
- Use @testing-library/jest-dom
- Set up test environment
- Configure TypeScript for tests

Testing Server Components:
- Test as regular async functions
- Mock database calls
- Test data fetching logic
- Verify rendered output
- Test error boundaries

Testing Client Components:
- Use @testing-library/react
- Test user interactions
- Mock useRouter and usePathname
- Test state management
- Verify event handlers

Testing Server Actions:
- Test as async functions
- Mock database operations
- Test validation logic
- Verify revalidation calls
- Test error handling

Mocking Next.js APIs:
- Mock next/navigation (useRouter, usePathname)
- Mock next/image component
- Mock next/link component
- Mock Server Actions
- Mock API route handlers

Integration Testing:
- Test component interactions
- Test data flow
- Test form submissions
- Test navigation flows
- Use MSW for API mocking

E2E Testing with Playwright:
- Install @playwright/test
- Configure playwright.config.ts
- Write user journey tests
- Test critical paths
- Test across browsers

Playwright Best Practices:
- Use page object pattern
- Test user flows end-to-end
- Use data-testid attributes
- Test responsive layouts
- Implement visual regression testing

Testing API Routes:
- Test route handlers directly
- Mock database and external APIs
- Test authentication
- Test error responses
- Test rate limiting

Testing Middleware:
- Test middleware logic
- Mock NextRequest and NextResponse
- Test authentication checks
- Test redirects
- Test header modifications

Snapshot Testing:
- Use for component output
- Update snapshots carefully
- Review snapshot changes
- Use for API responses
- Avoid over-using snapshots

Performance Testing:
- Use Lighthouse CI
- Test Core Web Vitals
- Monitor bundle size
- Test loading performance
- Use @next/bundle-analyzer

Accessibility Testing:
- Use jest-axe for a11y tests
- Test keyboard navigation
- Test screen reader compatibility
- Test ARIA attributes
- Use Playwright for a11y E2E tests

Test Coverage:
- Aim for 80%+ coverage
- Use jest --coverage
- Focus on critical paths
- Test edge cases
- Don't chase 100% coverage

Best Practices:
- Write tests before fixing bugs
- Test user behavior, not implementation
- Use meaningful test descriptions
- Keep tests isolated and independent
- Mock external dependencies
- Use factories for test data
- Run tests in CI/CD
- Test error states
- Use TypeScript in tests
- Maintain test documentation