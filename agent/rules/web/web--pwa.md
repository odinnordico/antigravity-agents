---
trigger: model_decision
description: When it is about web progressive app development, apply these rules.
---
You are an expert in Progressive Web App development.

Key Principles:
- Implement offline-first strategy
- Use service workers for caching
- Make app installable
- Ensure fast loading
- Provide app-like experience

Service Workers:
- Implement proper caching strategies
- Use Cache API effectively
- Handle offline scenarios
- Implement background sync
- Use workbox for easier implementation
- Handle service worker updates

Manifest:
- Create comprehensive web app manifest
- Define app icons for all sizes
- Set appropriate display mode
- Define theme and background colors
- Set start URL and scope
- Add screenshots for app stores

Caching Strategies:
- Use cache-first for static assets
- Use network-first for dynamic content
- Implement stale-while-revalidate
- Use cache-only for offline pages
- Implement proper cache versioning

Offline Experience:
- Provide offline fallback page
- Cache critical resources
- Implement background sync
- Show offline indicator
- Queue failed requests

Performance:
- Implement lazy loading
- Use code splitting
- Optimize images
- Minimize JavaScript
- Use HTTP/2 push
- Implement resource hints

Installability:
- Meet PWA criteria
- Implement beforeinstallprompt
- Provide install UI
- Test installation flow
- Handle app updates

Push Notifications:
- Implement push notification API
- Request permission appropriately
- Handle notification clicks
- Implement notification best practices
- Test on multiple platforms

Security:
- Serve over HTTPS
- Implement CSP headers
- Validate all inputs
- Use secure authentication
- Implement proper CORS

Testing:
- Use Lighthouse for audits
- Test offline functionality
- Test on multiple devices
- Test installation flow
- Test push notifications

Best Practices:
- Follow PWA checklist
- Implement progressive enhancement
- Provide app shell architecture
- Use PRPL pattern
- Monitor performance metrics