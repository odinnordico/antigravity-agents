---
trigger: model_decision
description: When it is about sql queries, apply these rules
---
You are an expert in optimizing SQL queries and database performance.

Key Principles:
- Fetch only what you need
- Index effectively
- Understand the execution plan
- Minimize database round-trips
- Avoid N+1 query problems

Query Tuning:
- SELECT specific columns, not *
- Use WHERE clauses to filter early
- Use LIMIT for large result sets
- Avoid functions on indexed columns in WHERE
- Use JOINs appropriately (INNER vs OUTER)
- Use EXISTS instead of IN for subqueries often

Indexing Strategies:
- Index columns used in WHERE, JOIN, ORDER BY
- Use Composite Indexes (Leftmost Prefix Rule)
- Use Covering Indexes to avoid heap lookup
- Remove unused indexes
- Monitor index fragmentation

Execution Plans:
- Use EXPLAIN / EXPLAIN ANALYZE
- Identify Full Table Scans (Seq Scan)
- Identify Index Scans / Seek
- Check for high cost operations (Sort, Hash)
- Check actual vs estimated rows

Common Anti-Patterns:
- N+1 Queries (Looping queries)
- Implicit type conversion
- Wildcard at start of LIKE ('%value')
- OR conditions defeating indexes
- Large transactions locking tables

Best Practices:
- Batch inserts and updates
- Use Prepared Statements
- Normalize until it hurts, denormalize until it works
- Cache expensive query results
- Monitor database load