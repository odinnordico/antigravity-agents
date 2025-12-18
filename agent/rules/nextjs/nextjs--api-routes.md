---
trigger: model_decision
description: When it is about Next.js API development, apply these rules.
---
You are an expert in Next.js API Routes and Route Handlers.

Key Principles:
- Use Route Handlers in App Router (route.ts)
- Implement proper HTTP methods
- Use TypeScript for type safety
- Handle errors gracefully
- Follow REST API conventions

Route Handler Basics:
- Create route.ts in app directory
- Export named functions: GET, POST, PUT, PATCH, DELETE
- Use NextRequest and NextResponse
- Access params and searchParams
- Return JSON responses

HTTP Methods:
- GET: Fetch data (use cache options)
- POST: Create resources
- PUT: Replace resources
- PATCH: Update resources partially
- DELETE: Remove resources
- OPTIONS: CORS preflight

Request Handling:
- Access request body with request.json()
- Get query params from request.nextUrl.searchParams
- Read headers with request.headers
- Access cookies with request.cookies
- Handle FormData with request.formData()

Response Types:
- JSON: NextResponse.json(data, { status })
- Text: new Response(text)
- Redirect: NextResponse.redirect(url)
- Stream: Use ReadableStream
- Error: NextResponse.json(error, { status: 4xx/5xx })

Validation:
- Use Zod for request validation
- Validate body, params, query
- Return 400 for validation errors
- Provide detailed error messages
- Use type-safe schemas

Authentication:
- Check session in route handlers
- Use getServerSession from next-auth
- Return 401 for unauthorized
- Implement API key authentication
- Use middleware for auth checks

Error Handling:
- Use try-catch blocks
- Return appropriate status codes
- Log errors server-side
- Don't expose sensitive errors
- Implement global error handler
- Use custom error classes

CORS Configuration:
- Set CORS headers in responses
- Handle OPTIONS requests
- Configure allowed origins
- Set allowed methods and headers
- Use middleware for CORS

Rate Limiting:
- Implement rate limiting per IP/user
- Use Upstash Redis for distributed rate limiting
- Return 429 Too Many Requests
- Add rate limit headers
- Implement different limits per endpoint

Caching:
- Use cache: 'force-cache' for static data
- Use cache: 'no-store' for dynamic data
- Implement revalidate for ISR
- Use NextResponse with cache headers
- Implement ETags for conditional requests

Streaming Responses:
- Use ReadableStream for large data
- Implement Server-Sent Events (SSE)
- Stream AI responses
- Handle backpressure
- Close streams properly

File Uploads:
- Handle FormData in POST requests
- Validate file types and sizes
- Use uploadthing or similar services
- Store files in cloud storage
- Return upload URLs

Webhooks:
- Verify webhook signatures
- Use raw body for signature verification
- Implement idempotency
- Return 200 quickly, process async
- Log webhook events

API Versioning:
- Use /api/v1/ prefix
- Maintain backward compatibility
- Document breaking changes
- Implement deprecation warnings
- Use headers for version negotiation

Best Practices:
- Use TypeScript for all routes
- Validate all inputs
- Return consistent response format
- Use proper HTTP status codes
- Implement request logging
- Use environment variables for config
- Test API routes thoroughly
- Document API with OpenAPI/Swagger
- Implement health check endpoint
- Monitor API performance and errors