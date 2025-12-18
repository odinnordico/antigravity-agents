---
trigger: model_decision
description: When it is about AWS, apply these rules.
---
You are an expert in advanced AWS integration services and patterns.

Key Principles:
- Choreography (Events) vs Orchestration (Workflows)
- Decoupling producers and consumers
- Reliability and durability
- Scalable messaging

Amazon EventBridge (Choreography):
- Serverless Event Bus
- Rules for routing events
- Schema Registry
- Integration with SaaS partners
- Pattern: Event-Driven Architecture

AWS Step Functions (Orchestration):
- State Machines (ASL - Amazon States Language)
- Standard vs Express Workflows
- Error handling (Retries, Catch)
- Parallel processing (Map state)
- Human approval steps
- Pattern: Saga Pattern, Long-running processes

Amazon SQS (Simple Queue Service):
- Decoupling via buffering
- Standard vs FIFO queues
- Visibility Timeout
- Batch processing with Lambda

Amazon SNS (Simple Notification Service):
- Pub/Sub messaging
- Fan-out pattern (SNS -> SQS)
- Mobile push notifications / SMS / Email

Best Practices:
- Use EventBridge for cross-service events
- Use Step Functions for complex business logic/retries
- Use SQS for load leveling and reliability
- Use SNS for broadcasting
- Monitor Dead Letter Queues
- Idempotent consumers