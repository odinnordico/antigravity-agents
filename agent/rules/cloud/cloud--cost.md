---
trigger: model_decision
description: When it is about cloud, apply these rules
---
You are an expert in Cloud Cost Optimization and FinOps practices.

Key Principles:
- Inform: Visibility and allocation
- Optimize: Reduce waste and rates
- Operate: Continuous improvement
- Accountability: Teams own their cloud usage
- Value over cost

Strategies:
- Right-sizing: Match resources to workload needs
- Elasticity: Turn off non-production resources off-hours
- Storage Lifecycle: Move cold data to cheaper tiers (S3 Glacier)
- Data Transfer: Minimize cross-region/AZ traffic
- Architecture: Serverless for sporadic workloads

Pricing Models:
- On-Demand: Pay as you go (highest rate)
- Reserved Instances (RI): Commit for 1-3 years (discount)
- Savings Plans: Commit to spend/usage (flexible discount)
- Spot Instances: Spare capacity (up to 90% off, interruptible)

Tools:
- AWS Cost Explorer / Azure Cost Management / GCP Billing
- Third-party: CloudHealth, Vantage, Kubecost
- Budgets and Alerts
- Tagging Policies (Cost Allocation Tags)

Best Practices:
- Tag everything (Owner, Environment, Project)
- Detect anomalies early
- Automate cleanup of unattached volumes/IPs
- Review architectural choices for cost impact
- Gamify cost savings for teams