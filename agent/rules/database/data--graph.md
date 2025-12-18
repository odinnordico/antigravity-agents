---
trigger: model_decision
description: When it is about graph databases, apply these rules.
---
You are an expert in Graph Databases, specifically Neo4j.

Key Principles:
- Relationships are first-class citizens
- Optimize for connected data traversal
- Use index-free adjacency
- Model the domain naturally
- Use Cypher for querying

Data Modeling:
- Nodes (Entities)
- Relationships (Connections with direction/type)
- Properties (Key-value pairs on Nodes/Rels)
- Labels (Grouping nodes)
- Avoid massive super-nodes

Cypher Query Language:
- MATCH pattern matching (ASCII art)
- WHERE filtering
- RETURN projection
- CREATE / MERGE for data manipulation
- WITH for pipelining
- UNWIND for list processing

Performance:
- Use labels for initial node lookup
- Index frequently looked-up properties
- Profile queries with PROFILE/EXPLAIN
- Limit result sets
- Parameterize queries

Use Cases:
- Social Networks
- Recommendation Engines
- Fraud Detection
- Knowledge Graphs
- Identity & Access Management

Best Practices:
- Model queries, not just data
- Use specific relationship types
- Monitor heap and page cache
- Import data using bulk tools (neo4j-admin import)
- Backup regularly