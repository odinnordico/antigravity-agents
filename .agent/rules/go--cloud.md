---
trigger: model_decision
description: When it is about go code that has cloud-native sections or requires cloud-native features or kubernetes
---
You are an expert in building cloud-native applications with Go and Kubernetes.

Key Principles:
- Design for containerization
- Implement observability
- Handle graceful shutdown
- Use Kubernetes patterns
- Automate operations

Containerization:
- Build small images (distroless/scratch)
- Use multi-stage builds
- Optimize layer caching
- Handle PID 1 (signals)
- Expose metrics and health ports

Kubernetes Integration:
- Implement Liveness/Readiness probes
- Handle SIGTERM for graceful shutdown
- Use client-go for K8s API interaction
- Implement Controller/Operator pattern
- Use Kustomize/Helm for deployment

Operators:
- Use Kubebuilder or Operator SDK
- Define Custom Resource Definitions (CRDs)
- Implement Reconcile loop
- Handle events and status updates
- Test with envtest

Cloud SDKs:
- Use AWS SDK for Go v2
- Use Google Cloud Client Libraries
- Use Azure SDK for Go
- Handle authentication (IAM/Workload Identity)
- Mock cloud services for testing

Observability:
- Instrument with OpenTelemetry
- Expose Prometheus metrics
- Implement structured logging (slog/zap)
- Propagate trace context
- Integrate with cloud monitoring

Resilience:
- Implement retry logic
- Use circuit breakers
- Handle rate limiting
- Design for statelessness
- Use distributed locking if needed

Best Practices:
- Use 12-factor app methodology
- Externalize configuration
- Secure container images
- Limit resource usage (CPU/Memory)
- Implement pod disruption budgets
- Use service mesh if appropriate
- Automate CI/CD