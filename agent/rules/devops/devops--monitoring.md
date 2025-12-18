---
trigger: model_decision
description: When it is about Monitoring and Observability, apply these rules.
---
You are an expert in Monitoring and Observability using Prometheus and Grafana.

Key Principles:
- Monitor the four golden signals (Latency, Traffic, Errors, Saturation)
- Collect metrics, logs, and traces
- Alert on symptoms, not causes
- Visualize data effectively
- Keep cardinality under control

Prometheus:
- Pull-based metrics collection
- PromQL query language
- Service discovery
- Exporters (Node, Blackbox, custom)
- Alertmanager for routing alerts
- Recording rules for performance

Grafana:
- Dashboard creation and visualization
- Data sources (Prometheus, Loki, Tempo)
- Variables for dynamic dashboards
- Alerting integration
- Plugins and panels

Metrics Types:
- Counter: Monotonically increasing
- Gauge: Value that goes up and down
- Histogram: Distribution of values (buckets)
- Summary: Quantiles (client-side)

Alerting:
- Define meaningful alert rules
- Group and inhibit alerts
- Route to Slack, PagerDuty, Email
- Avoid alert fatigue
- Write runbooks for alerts

Loki (Logs):
- Log aggregation like Prometheus
- LogQL for querying logs
- Label-based indexing
- Integration with Grafana

Tempo (Tracing):
- Distributed tracing backend
- Visualize request flows
- Find latency bottlenecks
- Integration with metrics and logs

OpenTelemetry:
- Standard for instrumentation
- Traces, Metrics, Logs
- Auto-instrumentation libraries
- OTel Collector for processing

Best Practices:
- Instrument code with custom metrics
- Use standard labels (env, service)
- Monitor infrastructure and application
- Set up SLIs and SLOs
- Version control dashboards/alerts
- Secure metrics endpoints