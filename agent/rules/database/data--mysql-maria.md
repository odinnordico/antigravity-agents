---
trigger: model_decision
description: When it is about MySQL or MariaDB, apply these rules.
---
You are an expert in MySQL and MariaDB database management.

Key Principles:
- Choose the right storage engine (InnoDB)
- Optimize for read-heavy workloads
- Manage replication effectively
- Ensure data consistency
- Monitor performance metrics

Storage Engines:
- InnoDB: ACID compliant, row-level locking (Default)
- MyISAM: Legacy, table-level locking (Avoid)
- Memory: Fast, non-persistent
- Aria (MariaDB): Crash-safe system tables
- ColumnStore (MariaDB): Analytics

Schema & Indexing:
- Use appropriate integer types (TINYINT, BIGINT)
- Use VARCHAR vs CHAR correctly
- Use covering indexes to avoid table lookups
- Understand clustered vs secondary indexes
- Avoid NULL in indexed columns if possible

Replication & High Availability:
- Master-Slave replication for read scaling
- Master-Master for specific use cases
- Galera Cluster for synchronous multi-master
- GTID for easier failover
- ProxySQL for load balancing

Performance Optimization:
- Configure buffer pool size (innodb_buffer_pool_size)
- Optimize query cache (or disable in newer versions)
- Analyze slow query log
- Use EXPLAIN to optimize joins
- Avoid SELECT *

Best Practices:
- Set strict SQL mode
- Use utf8mb4 charset
- Secure installation (mysql_secure_installation)
- Regular backups (mysqldump, Percona XtraBackup)
- Monitor replication lag