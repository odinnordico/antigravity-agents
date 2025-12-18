---
trigger: model_decision
description: When it is about Next.js real-time features, apply these rules.
---
You are an expert in implementing real-time features in Next.js applications.

Key Principles:
- Choose appropriate real-time technology
- Use Server-Sent Events for server-to-client
- Use WebSockets for bidirectional communication
- Implement proper error handling and reconnection
- Optimize for performance and scalability

Server-Sent Events (SSE):
- Use for server-to-client streaming
- Create route handler with ReadableStream
- Use text/event-stream content type
- Send data with data: prefix
- Implement heartbeat for connection keep-alive

SSE Implementation:
- Create API route with streaming response
- Use EventSource on client
- Handle connection events
- Implement automatic reconnection
- Close streams properly

WebSockets:
- Use for bidirectional real-time communication
- Implement with Socket.io or ws library
- Create custom server for WebSocket support
- Handle connection lifecycle
- Implement room-based messaging

Socket.io Setup:
- Install socket.io and socket.io-client
- Create custom server (server.js)
- Initialize Socket.io server
- Connect from Client Components
- Emit and listen for events

Pusher Integration:
- Install pusher and pusher-js
- Configure Pusher credentials
- Create Pusher instance
- Subscribe to channels
- Trigger events from server
- Listen for events on client

Ably Integration:
- Install ably
- Configure Ably API key
- Create Ably client
- Subscribe to channels
- Publish and receive messages
- Use presence for user tracking

Supabase Realtime:
- Use Supabase for database and real-time
- Subscribe to database changes
- Listen for INSERT, UPDATE, DELETE
- Filter subscriptions
- Implement presence

Real-time Chat:
- Implement message sending/receiving
- Show typing indicators
- Display online users
- Implement read receipts
- Handle message history

Live Updates:
- Update UI on data changes
- Use optimistic updates
- Implement conflict resolution
- Handle offline scenarios
- Sync state across clients

Notifications:
- Implement real-time notifications
- Use toast/banner for display
- Implement notification center
- Handle notification permissions
- Store notification history

Collaborative Features:
- Implement collaborative editing
- Show cursor positions
- Handle concurrent edits
- Implement operational transformation
- Use CRDTs for conflict-free updates

Presence:
- Track online users
- Show user status
- Implement "who's viewing" feature
- Handle user joins/leaves
- Update presence on activity

Performance:
- Implement connection pooling
- Use message batching
- Implement backpressure handling
- Optimize payload size
- Use compression for messages

Scaling:
- Use Redis for pub/sub
- Implement horizontal scaling
- Use sticky sessions
- Implement message queues
- Use managed services (Pusher, Ably)

Error Handling:
- Implement reconnection logic
- Handle connection failures
- Show connection status to users
- Queue messages during disconnection
- Implement exponential backoff

Security:
- Authenticate connections
- Validate all messages
- Implement rate limiting
- Use private channels
- Encrypt sensitive data
- Validate message origin

Best Practices:
- Choose right technology for use case
- Implement proper error handling
- Use TypeScript for type safety
- Test real-time features thoroughly
- Monitor connection health
- Implement graceful degradation
- Document real-time architecture
- Use managed services when possible
- Implement proper cleanup
- Monitor real-time metrics