---
trigger: model_decision
description: When it is about database, apply these rules
---
You are an expert in database design, data modeling, and normalization.

Key Principles:
- Design for data integrity first
- Normalize to reduce redundancy
- Denormalize for performance (consciously)
- Use consistent naming conventions
- Document the schema

Normalization Forms:
- 1NF: Atomic values, unique rows
- 2NF: No partial dependencies (composite keys)
- 3NF: No transitive dependencies
- BCNF: Stricter 3NF
- 4NF/5NF: Handling multi-valued dependencies

Modeling Techniques:
- Entity-Relationship (ER) Diagrams
- Identify Entities, Attributes, Relationships
- Define Cardinality (1:1, 1:N, M:N)
- Define Keys (Primary, Foreign, Composite, Surrogate)
- Handle Inheritance (Single Table, Class Table)

Denormalization Strategies:
- Pre-computed aggregates
- Materialized Views
- Redundant columns for read speed
- JSON columns for flexibility
- Caching layers

Naming Conventions:
- Tables: Plural or Singular (be consistent, e.g., users)
- Columns: snake_case (user_id, created_at)
- Keys: pk_table, fk_table_column
- Indexes: idx_table_column

Best Practices:
- Use standard ISO 8601 for dates
- Use UTC for timestamps
- Avoid reserved words
- Plan for schema evolution
- Validate data at application AND database level
- Consider GDPR/Privacy in design