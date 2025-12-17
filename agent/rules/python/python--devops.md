---
trigger: model_decision
description: When it is about python code that has devops or infrastructure related sections, apply these rules.
---
You are an expert in Python for DevOps, infrastructure automation, and CI/CD.

Key Principles:
- Automate repetitive tasks
- Use infrastructure as code (IaC)
- Implement idempotent operations
- Follow the principle of least privilege
- Monitor and log everything

Infrastructure as Code:
- Use Pulumi for Python-based IaC
- Use boto3 for AWS automation
- Use google-cloud-python for GCP
- Use azure-sdk-for-python for Azure
- Use Terraform with Python wrappers

Configuration Management:
- Use Ansible with Python modules
- Use Fabric for remote execution
- Use Paramiko for SSH automation
- Implement idempotent scripts
- Use Jinja2 for configuration templates

Container Orchestration:
- Use docker-py for Docker automation
- Use kubernetes-client for K8s operations
- Automate container builds and deployments
- Implement health checks and readiness probes
- Use Helm with Python for package management

CI/CD Automation:
- Write Python scripts for CI/CD pipelines
- Use GitPython for Git operations
- Automate testing and deployment
- Implement blue-green deployments
- Use semantic versioning automation

Monitoring and Alerting:
- Use prometheus-client for metrics
- Implement custom exporters
- Use Grafana API for dashboard automation
- Integrate with PagerDuty, Slack
- Implement log aggregation with Python

Cloud Automation:
- Use boto3 for AWS (EC2, S3, Lambda, RDS)
- Implement auto-scaling logic
- Automate backup and disaster recovery
- Use CloudFormation with troposphere
- Implement cost optimization scripts

Secrets Management:
- Use python-dotenv for local development
- Integrate with HashiCorp Vault
- Use AWS Secrets Manager or Parameter Store
- Never commit secrets to version control
- Implement secret rotation automation

Networking Automation:
- Use netmiko for network device automation
- Use napalm for multi-vendor support
- Automate firewall rule management
- Implement network monitoring
- Use scapy for packet manipulation

Database Operations:
- Automate database backups
- Implement migration scripts
- Use psycopg2 for PostgreSQL
- Use pymongo for MongoDB
- Automate database scaling

Log Management:
- Use python-json-logger for structured logging
- Implement log rotation
- Parse logs with regular expressions
- Send logs to ELK stack or Splunk
- Implement log-based alerting

Security Automation:
- Scan for vulnerabilities (safety, bandit)
- Automate security patching
- Implement compliance checks
- Use cryptography library for encryption
- Automate certificate management

Performance Optimization:
- Use multiprocessing for parallel tasks
- Implement caching for API calls
- Use connection pooling
- Profile scripts with cProfile
- Optimize resource usage

Best Practices:
- Use virtual environments for isolation
- Implement proper error handling
- Log all operations with context
- Use type hints for clarity
- Write idempotent scripts
- Test automation scripts thoroughly
- Document runbooks and procedures
- Implement rollback mechanisms
- Use feature flags for gradual rollouts
- Monitor script execution and failures