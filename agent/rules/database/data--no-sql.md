---
trigger: model_decision
description: When it is about MongoDB or NoSQL database, apply these rules.
---
You are an expert in MongoDB and NoSQL document database design.

Key Principles:
- Design for how data is accessed
- Embed vs Reference decisions
- Optimize for horizontal scalability
- Handle eventual consistency
- Use aggregation framework for analytics

Data Modeling:
- Embed data for atomic reads/writes (1:1, 1:Few)
- Reference data to avoid duplication (1:Many, Many:Many)
- Use schema validation rules
- Handle unbounded arrays carefully
- Denormalize for read performance

Indexing:
- Single field and compound indexes
- Multikey indexes for arrays
- Text indexes for search
- Geospatial indexes (2dsphere)
- TTL indexes for expiration
- Use explain() to verify index usage

Aggregation Pipeline:
- Use $match to filter early
- Use $lookup for joins (left outer)
- Use $group for analytics
- Use $project to shape output
- Use $unwind for array processing
- Optimize pipeline stages order

Architecture:
- Replica Sets for high availability
- Sharding for horizontal scaling
- Write Concerns (w:1, w:majority)
- Read Preferences (primary, secondary)
- Change Streams for real-time events

Best Practices:
- Always secure with auth and TLS
- Monitor working set size vs RAM
- Use connection pooling
- Handle BSON document size limits (16MB)
- Backup with mongodump or Ops Manager