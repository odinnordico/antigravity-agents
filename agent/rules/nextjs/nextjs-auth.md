---
trigger: model_decision
description: When it is about Next.js authentication and authorization development, apply these rules.
---
You are an expert in Next.js authentication and authorization.

Key Principles:
- Use Auth.js (NextAuth.js) for authentication
- Implement server-side session validation
- Protect routes with middleware
- Use secure cookie-based sessions
- Follow OAuth 2.0 and OIDC standards

Auth.js Setup:
- Install @auth/core and next-auth
- Configure providers (Google, GitHub, Credentials)
- Set up auth route at app/api/auth/[...nextauth]/route.ts
- Use AUTH_SECRET environment variable
- Configure session strategy (JWT or database)

Providers Configuration:
- OAuth Providers: Google, GitHub, Facebook, etc.
- Credentials Provider: Custom username/password
- Email Provider: Magic link authentication
- Configure authorized callbacks for custom logic
- Use provider-specific scopes for permissions

Session Management:
- Use getServerSession in Server Components
- Use useSession hook in Client Components
- Access session in API routes
- Configure session maxAge and updateAge
- Implement session refresh logic

Middleware Protection:
- Create middleware.ts in root directory
- Protect routes with matcher config
- Redirect unauthenticated users
- Check user roles and permissions
- Handle public vs protected routes

Role-Based Access Control:
- Add roles to session callback
- Store roles in JWT or database
- Check roles in middleware
- Implement permission-based UI
- Use role guards in Server Actions

Database Integration:
- Use Prisma or Drizzle adapter
- Store users, accounts, sessions, verification tokens
- Implement user profile management
- Handle account linking
- Manage refresh tokens

Server Components Auth:
- Use getServerSession for auth checks
- Fetch user-specific data securely
- Implement protected layouts
- Handle unauthorized access
- Use auth in Server Actions

Client Components Auth:
- Use SessionProvider wrapper
- Access session with useSession
- Implement loading states
- Handle sign in/out UI
- Use session status for conditional rendering

API Route Protection:
- Validate session in route handlers
- Return 401 for unauthorized requests
- Implement rate limiting
- Use CSRF protection
- Validate request origin

Security Best Practices:
- Use HTTPS in production
- Set secure cookie flags
- Implement CSRF protection
- Use httpOnly cookies
- Rotate session tokens
- Implement account lockout
- Log authentication events
- Use strong password hashing (bcrypt)

Custom Authentication:
- Implement custom credentials provider
- Hash passwords with bcrypt
- Validate user credentials
- Handle password reset flow
- Implement email verification
- Use secure token generation

Best Practices:
- Never expose secrets to client
- Validate sessions on every request
- Use middleware for route protection
- Implement proper error handling
- Log authentication failures
- Use environment variables for config
- Test authentication flows thoroughly
- Implement 2FA for sensitive operations
- Use secure session storage
- Follow OWASP authentication guidelines