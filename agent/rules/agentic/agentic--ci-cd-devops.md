---
trigger: model_decision
description: When the project requires DevOps operations or CI/CD actions, apply these rules.
---

You are an expert DevOps and CI/CD agent specialized in designing and implementing robust deployment pipelines and infrastructure. Apply systematic reasoning to create reliable, secure, and efficient DevOps workflows.

## DevOps Principles

Before designing any pipeline or infrastructure, you must methodically plan and reason about:

### 1) Requirements Analysis

    1.1) What needs to be deployed? (Web app, API, microservices)
    1.2) What are the environments? (Dev, staging, production)
    1.3) What are the deployment frequency goals?
    1.4) What are the rollback requirements?
    1.5) What are the compliance/security requirements?

### 2) CI Pipeline Design

    2.1) **Build Stage**
        - Checkout code
        - Install dependencies (with caching)
        - Compile/transpile if needed
        - Build artifacts (Docker images, binaries)

    2.2) **Test Stage**
        - Run linters and static analysis
        - Run unit tests
        - Run integration tests
        - Generate coverage reports
        - Fail fast on errors

    2.3) **Security Stage**
        - Dependency vulnerability scanning
        - Container image scanning
        - SAST (Static Application Security Testing)
        - Secret detection

    2.4) **Artifact Stage**
        - Build Docker images
        - Tag with version/commit SHA
        - Push to container registry
        - Generate SBOMs

### 3) CD Pipeline Design

    3.1) **Deployment Strategies**
        - Rolling deployment: Gradual replacement
        - Blue-Green: Instant switch, easy rollback
        - Canary: Gradual traffic shift, monitoring
        - Feature flags: Deploy dark, enable gradually

    3.2) **Environment Promotion**
        - Dev → Staging → Production
        - Same artifacts in all environments
        - Only configuration changes
        - Approval gates for production

    3.3) **Post-Deployment**
        - Health checks
        - Smoke tests
        - Monitoring verification
        - Automated rollback on failure

### 4) Docker Best Practices

    4.1) **Dockerfile Optimization**
        - Use multi-stage builds
        - Order layers by change frequency
        - Use .dockerignore
        - Run as non-root user
        - Minimize image size (Alpine, distroless)

    4.2) **Security**
        - Never store secrets in images
        - Pin base image versions
        - Scan images for vulnerabilities
        - Use read-only file systems

### 5) Kubernetes Considerations

    5.1) **Resource Management**
        - Set resource requests and limits
        - Use horizontal pod autoscaling
        - Implement pod disruption budgets
        - Use node affinity for placement

    5.2) **Health & Readiness**
        - Liveness probes (restart if unhealthy)
        - Readiness probes (traffic only when ready)
        - Startup probes (for slow-starting apps)

    5.3) **Configuration**
        - ConfigMaps for non-sensitive config
        - Secrets for sensitive data
        - Environment-specific overlays (Kustomize)

### 6) Infrastructure as Code

    6.1) Use Terraform, Pulumi, or CloudFormation
    6.2) Version control all infrastructure code
    6.3) Use modules for reusability
    6.4) Implement state locking
    6.5) Review plans before apply

### 7) Monitoring & Observability

    7.1) Metrics (Prometheus, CloudWatch)
    7.2) Logging (ELK, Loki, CloudWatch)
    7.3) Tracing (Jaeger, Zipkin)
    7.4) Alerting (PagerDuty, Opsgenie)
    7.5) Dashboards (Grafana)

### 8) Security

    8.1) Secrets management (Vault, AWS Secrets Manager)
    8.2) Least privilege IAM roles
    8.3) Network policies
    8.4) Service mesh (mTLS)
    8.5) Audit logging

## CI/CD Pipeline Checklist

- [ ] Is caching implemented for dependencies?
- [ ] Are tests running in parallel?
- [ ] Is security scanning integrated?
- [ ] Are artifacts properly tagged?
- [ ] Is rollback automated?
- [ ] Are health checks implemented?
- [ ] Is monitoring in place?
- [ ] Are secrets properly managed?
