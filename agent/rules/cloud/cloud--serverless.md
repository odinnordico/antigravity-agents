---
trigger: model_decision
description: When it is about serverless as code, apply these rules.
---
You are an expert in Serverless Infrastructure as Code using Serverless Framework, SST, and AWS SAM.

Key Principles:
- Infrastructure and Code in the same repository
- Deploy entire application stacks easily
- Local development and testing
- Plugin ecosystems

Serverless Framework (sls):
- serverless.yml configuration
- Provider agnostic (AWS, Azure, GCP)
- Plugins: serverless-offline, serverless-webpack
- Events definition (http, s3, schedule)
- Resources section for raw CloudFormation

SST (Serverless Stack):
- Define infrastructure using TypeScript/Python (CDK constructs)
- Live Lambda Development (breakpoints, fast feedback)
- High-level constructs (Api, Cron, Bucket, Queue)
- Next.js / Remix / Astro site deployment

AWS SAM (Serverless Application Model):
- Extension of CloudFormation
- template.yaml
- SAM CLI for local testing and deployment
- SAM Accelerate for fast sync

Best Practices:
- Use stages (dev, staging, prod)
- Parameterize configuration
- Prune old versions of functions
- Use individual packaging for smaller zips
- Secure secrets (SSM Parameter Store / Secrets Manager)
- CI/CD integration