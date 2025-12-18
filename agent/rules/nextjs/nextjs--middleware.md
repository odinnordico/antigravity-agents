---
trigger: model_decision
description: When it is about Next.js middleware, apply these rules
---
You are an expert in Next.js Middleware and Edge Functions.

Key Principles:
- Use middleware for request/response manipulation
- Run logic before request completes
- Use Edge Runtime for global performance
- Implement proper matcher configuration
- Keep middleware lightweight

Middleware Basics:
- Create middleware.ts in root directory
- Export middleware function
- Use NextRequest and NextResponse
- Configure matcher for specific routes
- Run before cached content

Common Use Cases:
- Authentication and authorization
- Redirects and rewrites
- Geolocation-based routing
- A/B testing
- Bot detection
- Rate limiting
- Header manipulation
- Cookie management

Authentication:
- Check auth token in middleware
- Redirect to login if unauthorized
- Add user info to request headers
- Implement role-based access
- Use next-auth middleware

Redirects and Rewrites:
- Use NextResponse.redirect()
- Use NextResponse.rewrite()
- Implement locale-based redirects
- Handle legacy URLs
- Implement custom routing logic

Geolocation:
- Access geo data from request
- Route based on country/city
- Implement region-specific content
- Use for compliance (GDPR)
- Personalize user experience

A/B Testing:
- Assign users to variants
- Use cookies for persistence
- Rewrite to variant pages
- Track experiment data
- Implement gradual rollouts

Headers and Cookies:
- Set custom headers
- Read and set cookies
- Implement security headers
- Add CORS headers
- Set cache headers

Matcher Configuration:
- Use matcher to filter routes
- Support glob patterns
- Exclude static files
- Use negative patterns
- Match specific paths only

Edge Runtime:
- Runs globally on edge network
- Use for low-latency responses
- Limited Node.js APIs
- Use Web APIs instead
- Optimize for cold starts

Edge Functions:
- Create in app/api with edge runtime
- Export runtime = 'edge'
- Use for dynamic content
- Implement streaming responses
- Use for real-time features

Performance:
- Keep middleware fast (<50ms)
- Avoid heavy computations
- Use edge-compatible libraries
- Minimize external API calls
- Cache when possible

Security:
- Validate all inputs
- Implement rate limiting
- Add security headers
- Prevent CSRF attacks
- Use secure cookies
- Implement bot protection

Testing Middleware:
- Test locally with next dev
- Mock NextRequest/NextResponse
- Test matcher patterns
- Test edge cases
- Use integration tests

Debugging:
- Use console.log (appears in terminal)
- Test in production-like environment
- Use Vercel deployment previews
- Check middleware execution order
- Monitor performance metrics

Best Practices:
- Keep middleware lightweight
- Use matcher to limit execution
- Avoid blocking operations
- Use edge-compatible code
- Implement proper error handling
- Test thoroughly before deploying
- Monitor middleware performance
- Use TypeScript for type safety
- Document middleware logic
- Version control middleware changes