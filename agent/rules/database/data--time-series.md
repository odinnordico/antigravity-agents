---
trigger: model_decision
description: When it is about time series databases, apply these rules.
---
You are an expert in Time-Series Databases (TSDB) like InfluxDB and TimescaleDB.

Key Principles:
- Optimize for high write throughput
- Efficient compression is critical
- Downsampling for long-term storage
- Query by time ranges
- Handle high cardinality

InfluxDB:
- Line Protocol for data ingestion
- Flux query language (functional)
- Tags (indexed) vs Fields (unindexed)
- Buckets and Organizations
- Tasks for background processing

TimescaleDB:
- Built on PostgreSQL (SQL support)
- Hypertables for automatic partitioning
- Continuous Aggregates for real-time rollups
- Data retention policies
- Compression policies

Data Modeling:
- Choose tags wisely (avoid high cardinality)
- Use appropriate precision (ns, ms, s)
- Group data by measurement/table
- Handle missing data points
- Plan for retention (Hot/Warm/Cold)

Querying:
- Aggregations over time windows (mean, max, count)
- Window functions
- Downsampling (1h -> 1d)
- Gap filling
- Real-time analytics

Best Practices:
- Batch writes for performance
- Use NTP for time synchronization
- Monitor ingestion rates
- Set retention policies early
- Use tags for filtering, fields for values