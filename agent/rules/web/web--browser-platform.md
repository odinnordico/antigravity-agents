---
trigger: model_decision
description: When it is about web development for browser APIs, apply these rules.
---
You are an expert in browser APIs and the modern web platform.

Key Principles:
- Use modern browser APIs
- Implement progressive enhancement
- Handle browser compatibility
- Use feature detection
- Follow web standards

Storage APIs:
- localStorage for persistent data
- sessionStorage for session data
- IndexedDB for large datasets
- Cache API for offline support
- Use proper error handling
- Implement storage quotas

Fetch API:
- Use fetch for HTTP requests
- Handle responses properly
- Implement error handling
- Use AbortController for cancellation
- Implement request/response interceptors
- Handle CORS properly

Intersection Observer:
- Use for lazy loading
- Implement infinite scroll
- Track element visibility
- Optimize performance
- Use proper thresholds
- Handle cleanup properly

Mutation Observer:
- Watch DOM changes
- Implement proper callbacks
- Use for dynamic content
- Handle performance implications
- Disconnect when not needed
- Use specific observation targets

Resize Observer:
- Watch element size changes
- Implement responsive components
- Handle resize events efficiently
- Use for container queries
- Implement proper cleanup

Geolocation API:
- Request user permission
- Handle position data
- Implement error handling
- Use watchPosition for tracking
- Handle privacy concerns
- Implement fallbacks

Notifications API:
- Request permission properly
- Create rich notifications
- Handle notification clicks
- Implement notification actions
- Use service worker for background notifications
- Respect user preferences

Web Workers:
- Offload heavy computation
- Use for background tasks
- Implement message passing
- Handle worker lifecycle
- Use SharedWorker for shared state
- Implement proper error handling

WebRTC:
- Implement peer-to-peer communication
- Handle media streams
- Use for video/audio calls
- Implement screen sharing
- Handle connection states
- Implement proper error handling

WebSockets:
- Implement real-time communication
- Handle connection lifecycle
- Implement reconnection logic
- Use for chat and live updates
- Handle message queuing
- Implement proper error handling

Clipboard API:
- Use for copy/paste functionality
- Request permissions properly
- Handle different data types
- Implement fallbacks
- Use async clipboard API
- Handle security restrictions

Payment Request API:
- Implement web payments
- Handle payment methods
- Validate payment data
- Implement proper error handling
- Use for checkout flows
- Handle payment completion

Web Share API:
- Implement native sharing
- Check for API support
- Handle share data
- Implement fallbacks
- Use for social sharing
- Handle share targets

Best Practices:
- Use feature detection
- Implement progressive enhancement
- Handle errors gracefully
- Use polyfills when needed
- Test across browsers
- Follow web standards
- Implement proper permissions
- Use HTTPS for secure APIs
- Handle privacy concerns
- Document API usage