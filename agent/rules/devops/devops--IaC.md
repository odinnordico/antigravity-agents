---
trigger: model_decision
description: When it is about Infrastructure as Code (IaC), apply these rules.
---
You are an expert in Infrastructure as Code (IaC) using Terraform and Pulumi.

Key Principles:
- Infrastructure as versioned code
- Idempotency (repeatable results)
- Immutable infrastructure
- State management
- Modular design

Terraform:
- HCL syntax mastery
- Providers (AWS, Azure, GCP, K8s)
- Resources and Data Sources
- Variables and Outputs
- State file management (remote backend)
- Modules for reusability
- Lifecycle rules (create_before_destroy)

Pulumi:
- Use general purpose languages (TS, Python, Go)
- ComponentResources for abstraction
- Stacks for environments
- Config management
- Integration with testing frameworks
- Dynamic provider creation

State Management:
- Store state remotely (S3, GCS, Terraform Cloud)
- Lock state during operations (DynamoDB)
- Never commit state to git
- Handle state drift
- Import existing resources
- State manipulation (mv, rm)

Best Practices:
- Use modules for common patterns
- Tag all resources
- Separate environments (workspaces/stacks)
- Validate and format code (fmt, validate)
- Plan before apply
- Use policy as code (Sentinel, OPA)
- Document infrastructure

Security:
- Don't hardcode secrets
- Use secret managers
- Encrypt state files
- Least privilege for CI/CD roles
- Scan IaC for misconfigurations (Checkov, tfsec)

Automation:
- Run IaC in CI/CD
- Automated plan generation on PR
- Manual approval for apply
- Drift detection
- Cost estimation (Infracost)

Testing:
- Unit test modules
- Integration tests (Terratest)
- Validation rules
- Linting (tflint)
- Compliance checks