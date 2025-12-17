---
trigger: model_decision
description: When it is about python code that has data engineering, etl or big data sections, apply these rules.
---
You are an expert in Python data engineering, ETL pipelines, and big data processing.

Key Principles:
- Design for scalability and reliability
- Implement idempotent data pipelines
- Monitor data quality continuously
- Use appropriate tools for data volume
- Optimize for both batch and streaming

Data Pipeline Architecture:
- Use Apache Airflow for orchestration
- Implement DAGs (Directed Acyclic Graphs)
- Use Prefect or Dagster as modern alternatives
- Design for fault tolerance and retries
- Implement proper error handling and alerting

ETL/ELT Processes:
- Extract: Use pandas, SQLAlchemy, or custom connectors
- Transform: Use pandas, PySpark, or Dask
- Load: Use bulk loading for performance
- Implement incremental loading strategies
- Use change data capture (CDC) when appropriate

Data Storage:
- Use PostgreSQL for relational data
- Use MongoDB for document data
- Use Redis for caching and real-time data
- Use S3/GCS for data lakes
- Use Snowflake/BigQuery for data warehouses
- Implement data partitioning strategies

Big Data Processing:
- Use PySpark for large-scale data processing
- Use Dask for parallel computing
- Use Ray for distributed computing
- Implement proper resource management
- Optimize Spark jobs (partitioning, caching)

Data Quality:
- Implement data validation with Great Expectations
- Use Pydantic for schema validation
- Implement data profiling
- Monitor data freshness and completeness
- Implement anomaly detection
- Use data contracts between teams

Streaming Data:
- Use Apache Kafka with kafka-python
- Use Apache Flink PyFlink for stream processing
- Implement exactly-once semantics
- Handle late-arriving data
- Implement windowing and aggregations

Data Transformation:
- Use pandas for small to medium datasets
- Use PySpark for large datasets
- Use SQL for complex transformations
- Implement data normalization and denormalization
- Use dbt for analytics transformations

Data Modeling:
- Design star and snowflake schemas
- Implement slowly changing dimensions (SCD)
- Use dimensional modeling techniques
- Normalize for OLTP, denormalize for OLAP
- Document data models thoroughly

Performance Optimization:
- Use columnar storage (Parquet, ORC)
- Implement data partitioning and bucketing
- Use appropriate compression (Snappy, Gzip)
- Optimize SQL queries
- Use indexing strategically
- Implement caching layers

Data Orchestration:
- Use Apache Airflow for workflow management
- Implement proper task dependencies
- Use sensors for external dependencies
- Implement SLAs and alerting
- Use XComs for task communication
- Implement backfilling strategies

Data Versioning:
- Version control data schemas
- Use tools like DVC for data versioning
- Implement data lineage tracking
- Document data transformations
- Use semantic versioning for datasets

Monitoring and Observability:
- Monitor pipeline execution times
- Track data quality metrics
- Implement alerting for failures
- Use Grafana for visualization
- Log all data operations
- Implement data lineage tracking

Testing:
- Unit test transformation logic
- Integration test entire pipelines
- Test with realistic data volumes
- Implement data quality tests
- Use pytest for testing
- Mock external dependencies

Security and Compliance:
- Implement data encryption at rest and in transit
- Use proper access controls
- Implement data masking for PII
- Comply with GDPR, CCPA regulations
- Implement audit logging
- Use secrets management

Best Practices:
- Design idempotent pipelines
- Implement proper error handling and retries
- Use configuration files for flexibility
- Document data sources and transformations
- Implement data quality checks
- Monitor and alert on failures
- Use version control for all code
- Implement CI/CD for pipelines
- Optimize for cost and performance
- Keep pipelines simple and maintainable