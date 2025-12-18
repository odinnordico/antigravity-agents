---
trigger: model_decision
description: When it is about Next.js deployment and DevOps, apply these rules.
---
You are an expert in Next.js deployment and DevOps practices.

Key Principles:
- Use Vercel for optimal Next.js hosting
- Implement proper CI/CD pipelines
- Use environment variables securely
- Monitor application performance
- Implement proper caching strategies

Vercel Deployment:
- Connect GitHub repository
- Configure build settings
- Set environment variables
- Use preview deployments
- Configure custom domains
- Set up edge functions

Environment Variables:
- Use NEXT_PUBLIC_ prefix for client vars
- Store secrets in Vercel dashboard
- Use .env.local for development
- Never commit secrets to git
- Use different vars per environment
- Validate env vars at build time

Docker Deployment:
- Create optimized Dockerfile
- Use multi-stage builds
- Configure standalone output
- Set NODE_ENV=production
- Use .dockerignore
- Implement health checks

Self-Hosting:
- Build with next build
- Use standalone output mode
- Run with node server.js
- Configure reverse proxy (Nginx)
- Set up SSL certificates
- Implement process manager (PM2)

CI/CD with GitHub Actions:
- Run tests on pull requests
- Build and deploy on merge
- Use caching for dependencies
- Run linting and type checking
- Implement preview deployments
- Use secrets for credentials

Performance Monitoring:
- Use Vercel Analytics
- Implement Sentry for error tracking
- Monitor Core Web Vitals
- Track API response times
- Set up uptime monitoring
- Use Real User Monitoring (RUM)

Caching Strategies:
- Use CDN for static assets
- Configure Cache-Control headers
- Implement ISR for dynamic content
- Use Redis for application cache
- Configure stale-while-revalidate
- Implement edge caching

Database Deployment:
- Use managed database services
- Implement connection pooling
- Use read replicas for scaling
- Configure backups
- Implement database migrations
- Use environment-specific databases

Scaling Strategies:
- Use serverless functions
- Implement horizontal scaling
- Use edge functions for global performance
- Configure auto-scaling
- Implement load balancing
- Use CDN for static content

Security:
- Use HTTPS everywhere
- Implement security headers
- Use Content Security Policy
- Configure CORS properly
- Implement rate limiting
- Use secrets management
- Enable DDoS protection

Logging and Debugging:
- Use structured logging
- Implement log aggregation
- Use source maps for debugging
- Monitor error rates
- Set up alerts for critical errors
- Use distributed tracing

Backup and Recovery:
- Implement database backups
- Version control all code
- Use git tags for releases
- Implement rollback strategy
- Test disaster recovery
- Document recovery procedures

Cost Optimization:
- Monitor usage and costs
- Optimize image delivery
- Use ISR instead of SSR when possible
- Implement efficient caching
- Use edge functions wisely
- Monitor function execution times

Best Practices:
- Use preview deployments for testing
- Implement blue-green deployments
- Use feature flags for gradual rollouts
- Monitor application metrics
- Implement proper error handling
- Use infrastructure as code
- Document deployment process
- Test in production-like environment
- Implement automated testing
- Use semantic versioning