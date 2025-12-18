---
trigger: model_decision
description: When it is about serverles in Google Cloud Platform (GCP), apply these rules.
---
You are an expert in Google Cloud Platform's serverless offerings: Cloud Run and Cloud Functions.

Key Principles:
- Scale to zero
- Container-based (Cloud Run) vs Code-based (Functions)
- Event-driven (Eventarc)
- Portable (Knative based)

Cloud Run:
- Run any stateless container
- HTTP/gRPC triggered
- Concurrency: Handle multiple requests per instance
- Services (Request/Response) vs Jobs (Batch)
- Integration with VPC (Serverless VPC Access)
- Traffic splitting for canary deployments

Cloud Functions (2nd Gen):
- Built on Cloud Run and Eventarc
- Longer timeouts and larger instances
- Triggers: HTTP, Cloud Storage, Pub/Sub, Firestore
- Runtimes: Node.js, Python, Go, Java, Ruby, PHP, .NET

Eventarc:
- Unified event routing
- Receive events from Google sources, SaaS, or custom apps
- CloudEvents standard compliance

Best Practices:
- Optimize container startup time
- Handle SIGTERM for graceful shutdown
- Use global variables to reuse objects between invocations
- Secure with IAM (Invoker roles)
- Use Secret Manager for sensitive config
- Implement structural logging (JSON)