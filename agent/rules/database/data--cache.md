---
trigger: model_decision
description: When it is about cache or Redis, apply these rules.
---
You are an expert in Redis and application caching strategies.

Key Principles:
- Cache for read-heavy workloads
- Handle cache invalidation correctly
- Use appropriate data structures
- Ensure high availability
- Monitor memory usage

Data Structures:
- Strings: Simple key-value, counters
- Lists: Queues, recent items
- Sets: Unique items, tags
- Sorted Sets: Leaderboards, priority queues
- Hashes: Objects, profiles
- Streams: Event logs, messaging
- HyperLogLog: Cardinality estimation

Caching Patterns:
- Cache-Aside (Lazy Loading): App checks cache, then DB
- Write-Through: App writes to cache and DB synchronously
- Write-Behind: App writes to cache, async to DB
- Cache Stampede prevention (locking, probabilistic expiration)

Persistence & Durability:
- RDB (Snapshots): Point-in-time backups
- AOF (Append Only File): Log every write
- Disable persistence for pure cache
- Hybrid approach (RDB + AOF)

Advanced Features:
- Pub/Sub for messaging
- Lua scripting for atomic operations
- Redis Transactions (MULTI/EXEC)
- Redis Cluster for sharding
- Redis Sentinel for HA
- Eviction policies (allkeys-lru, volatile-ttl)

Best Practices:
- Set TTL on all keys
- Use namespacing (user:123:profile)
- Monitor hit/miss ratio
- Avoid Keys command in production (use Scan)
- Pipeline commands for throughput
- Secure with password/ACL