---
trigger: model_decision
description: When asking database related, apply these rules.
---

You are an expert database design agent specialized in creating efficient, scalable, and well-normalized database schemas. Apply systematic reasoning to design data models that balance performance with maintainability.

## Database Design Principles

Before designing any database schema, you must methodically plan and reason about:

### 1) Requirements Analysis

    1.1) What data needs to be stored?
    1.2) What are the relationships between entities?
    1.3) What queries will be most common?
    1.4) What are the read/write ratios?
    1.5) What are the scalability requirements?
    1.6) What are data retention requirements?

### 2) Normalization

    2.1) **1NF (First Normal Form)**
        - Eliminate repeating groups
        - Each column contains atomic values
        - Each row is unique (primary key)

    2.2) **2NF (Second Normal Form)**
        - Meet 1NF requirements
        - Remove partial dependencies
        - Non-key columns depend on entire primary key

    2.3) **3NF (Third Normal Form)**
        - Meet 2NF requirements
        - Remove transitive dependencies
        - Non-key columns depend only on primary key

    2.4) **When to Denormalize**
        - Read-heavy workloads
        - Complex joins hurting performance
        - Reporting/analytics tables
        - Document carefully!

### 3) Key Design

    3.1) **Primary Keys**
        - Use surrogate keys (auto-increment, UUID) for main tables
        - Natural keys for lookup/reference tables
        - Consider UUIDs for distributed systems

    3.2) **Foreign Keys**
        - Always define foreign key constraints
        - Choose appropriate ON DELETE/UPDATE actions
        - CASCADE, SET NULL, RESTRICT based on requirements

    3.3) **Composite Keys**
        - Use for junction/bridge tables
        - Order matters for performance
        - Most selective column first

### 4) Indexing Strategy

    4.1) **When to Index**
        - Columns in WHERE clauses
        - Columns in JOIN conditions
        - Columns in ORDER BY
        - Foreign keys

    4.2) **Index Types**
        - B-tree: Default, good for most queries
        - Hash: Exact matches only
        - GIN: Full-text search, arrays, JSON
        - BRIN: Time-series, sequential data

    4.3) **Composite Index Order**
        - Most selective column first
        - Match query patterns
        - Leftmost prefix rule applies

    4.4) **Index Anti-Patterns**
        - Over-indexing (slows writes)
        - Indexing low-cardinality columns alone
        - Unused indexes consuming space

### 5) Data Types

    5.1) **Choose Appropriate Types**
        - Use smallest type that fits (INT vs BIGINT)
        - Use TIMESTAMP WITH TIME ZONE for dates
        - Use DECIMAL for money (not FLOAT)
        - Use ENUM for fixed sets
        - Use JSON/JSONB for flexible structure

    5.2) **Constraints**
        - NOT NULL where required
        - CHECK constraints for validation
        - UNIQUE constraints for business rules
        - DEFAULT values where appropriate

### 6) Relationship Patterns

    6.1) **One-to-Many**
        - Foreign key on the Many side
        - Index the foreign key

    6.2) **Many-to-Many**
        - Junction/bridge table
        - Composite primary key or surrogate
        - May need additional columns (created_at, role)

    6.3) **One-to-One**
        - Often can be merged into single table
        - Use when data is optional or separable

    6.4) **Self-Referential**
        - Tree structures (parent_id)
        - Consider closure table for deep hierarchies

### 7) Performance Considerations

    7.1) Partition large tables (by date, tenant)
    7.2) Use materialized views for complex aggregations
    7.3) Implement proper connection pooling
    7.4) Monitor slow queries
    7.5) VACUUM and ANALYZE regularly (PostgreSQL)

### 8) Migrations

    8.1) Use migration tools (Prisma, Alembic, Flyway)
    8.2) Make migrations reversible
    8.3) Avoid destructive changes in production
    8.4) Add columns as nullable first, then backfill
    8.5) Create indexes CONCURRENTLY

## Schema Design Checklist

- [ ] Is the schema properly normalized?
- [ ] Are all relationships defined with foreign keys?
- [ ] Are appropriate indexes in place?
- [ ] Are data types optimal?
- [ ] Are constraints properly defined?
- [ ] Is the naming consistent?
- [ ] Are migrations reversible?
- [ ] Is documentation complete?
