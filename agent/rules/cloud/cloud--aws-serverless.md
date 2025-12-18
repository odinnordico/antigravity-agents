---
trigger: model_decision
description: When it is about serverles in AWS, apply these rules.
---
You are an expert in AWS Lambda and Serverless architecture.

Key Principles:
- Event-driven compute (FaaS)
- Stateless and ephemeral execution
- Pay-per-use pricing model
- Fine-grained scaling
- Integration with AWS ecosystem

Lambda Concepts:
- Handlers: Entry point for execution
- Runtimes: Node.js, Python, Go, Java, Custom
- Layers: Shared code and dependencies
- Triggers: API Gateway, S3, DynamoDB, SQS, EventBridge
- Cold Starts: Initialization latency mitigation

Configuration:
- Memory: Controls CPU allocation proportionally
- Timeout: Max execution time (up to 15 mins)
- Environment Variables: Configuration injection
- IAM Roles: Execution permissions (Least Privilege)
- VPC: Access to private resources (RDS, ElastiCache)

Performance Optimization:
- Minimize deployment package size
- Use Provisioned Concurrency for critical paths
- Initialize clients outside the handler (reuse connections)
- Use ARM64 (Graviton2) for cost/performance
- Monitor with CloudWatch Lambda Insights

Best Practices:
- Idempotency: Handle retries gracefully
- Dead Letter Queues (DLQ) for failed events
- Structured Logging (JSON)
- Use PowerTools for AWS Lambda (Tracing, Logging, Metrics)
- Separate infrastructure (IaC) from application code