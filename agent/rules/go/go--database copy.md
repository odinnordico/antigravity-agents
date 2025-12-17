---
trigger: model_decision
description: When it is about go code that has database interaction sections or uses gorm, sqlx or database/sql, apply these rules.
---
You are an expert in Go database integration using database/sql, GORM, and sqlx.

Key Principles:
- Use connection pooling
- Use prepared statements
- Handle transactions properly
- Prevent SQL injection
- Manage migrations

database/sql:
- Open connection with sql.Open
- Ping to verify connection
- Use Query/QueryRow for reads
- Use Exec for writes
- Scan results into variables

GORM (ORM):
- Define models with struct tags
- Use AutoMigrate for schema
- Use Create/First/Find/Update/Delete
- Use Preload for associations
- Use Hooks for lifecycle events

sqlx (Extensions):
- Use StructScan for mapping
- Use NamedExec for named parameters
- Use Select/Get for convenience
- Better support for bulk operations
- Compatible with database/sql

Connection Management:
- SetMaxOpenConns
- SetMaxIdleConns
- SetConnMaxLifetime
- Monitor connection stats
- Handle connection errors

Transactions:
- Begin transaction with tx.Begin()
- Commit with tx.Commit()
- Rollback with tx.Rollback()
- Use defer for rollback
- Handle transaction isolation levels

Migrations:
- Use golang-migrate or goose
- Version control migrations
- Apply migrations on startup or CLI
- Handle up/down migrations
- Test migrations

Performance:
- Index columns properly
- Optimize queries
- Batch inserts/updates
- Use prepared statements
- Profile database performance

NoSQL:
- Use mongo-driver for MongoDB
- Use go-redis for Redis
- Handle BSON/JSON mapping
- Manage connections
- Handle specific NoSQL patterns

Best Practices:
- Use context for timeouts
- Handle NULL values (sql.NullString)
- Sanitize inputs
- Log slow queries
- Use interfaces for repositories
- Mock database for testing
- Secure credentials