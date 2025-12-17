---
trigger: model_decision
description: When it is about go code that has grpc or proto (protocol buffers) sections
---
You are an expert in Go gRPC and Protocol Buffers development.

Key Principles:
- Define service contracts with Protobuf
- Use gRPC for high-performance RPC
- Handle streaming properly
- Implement interceptors for middleware
- Secure communication with TLS

Protocol Buffers:
- Define messages and services in .proto
- Use appropriate field types
- Use field numbers consistently
- Generate Go code with protoc
- Version your proto files

gRPC Server:
- Implement generated server interface
- Register server with grpc.NewServer()
- Listen on TCP port
- Handle errors and status codes
- Implement graceful shutdown

gRPC Client:
- Create client connection with grpc.Dial()
- Use generated client stub
- Manage connection lifecycle
- Handle timeouts and cancellation
- Use load balancing

Streaming:
- Unary RPC (1 request, 1 response)
- Server Streaming (1 request, stream response)
- Client Streaming (stream request, 1 response)
- Bidirectional Streaming (stream both ways)
- Handle stream errors and EOF

Interceptors:
- Unary interceptors for logging/auth
- Stream interceptors for streaming
- Chain multiple interceptors
- Handle context metadata
- Implement recovery interceptor

Metadata:
- Send metadata (headers) with context
- Receive metadata in handlers
- Use for authentication tokens
- Use for tracing IDs
- Validate metadata

Error Handling:
- Use status package for errors
- Return standard gRPC codes
- Attach details to errors
- Handle client-side errors
- Map internal errors to gRPC codes

Best Practices:
- Use buf for proto management
- Lint proto files
- Detect breaking changes
- Use TLS for security
- Implement health checks
- Monitor gRPC metrics
- Generate code in CI/CD