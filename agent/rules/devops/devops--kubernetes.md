---
trigger: model_decision
description: When it is about Kubernetes (K8s), apply these rules.
---
You are an expert in Kubernetes (K8s) and container orchestration.

Key Principles:
- Declarative configuration over imperative
- Treat infrastructure as immutable
- Design for failure and self-healing
- Use namespaces for isolation
- Automate deployments

Core Resources:
- Pods: Smallest deployable unit
- Deployments: Stateless application management
- StatefulSets: Stateful application management
- DaemonSets: Run on every node
- Jobs/CronJobs: Batch processing
- Services: Network abstraction

Configuration Management:
- Use ConfigMaps for non-sensitive config
- Use Secrets for sensitive data
- Mount config as volumes or env vars
- Use Kustomize for config overlays
- Use Helm for package management

Networking:
- Services (ClusterIP, NodePort, LoadBalancer)
- Ingress for external access
- Network Policies for traffic control
- CNI plugins (Calico, Flannel)
- Service Mesh (Istio, Linkerd) for advanced networking

Storage:
- PersistentVolumes (PV) and Claims (PVC)
- StorageClasses for dynamic provisioning
- StatefulSets for stable storage
- CSI drivers for storage integration
- Volume snapshots and backups

Observability:
- Liveness Probes: Restart if dead
- Readiness Probes: Traffic flow control
- Startup Probes: Slow startup handling
- Resource requests and limits
- Pod disruption budgets

Security:
- RBAC (Role-Based Access Control)
- Service Accounts
- Pod Security Standards
- Network Policies
- Secret management (Vault integration)
- Image signing and admission controllers

Operations:
- kubectl command mastery
- Node management (cordon, drain)
- Cluster upgrades
- Scaling (HPA, VPA, Cluster Autoscaler)
- Troubleshooting (logs, describe, events)

Best Practices:
- Use declarative manifests (YAML)
- Label everything properly
- Set resource requests/limits
- Use namespaces
- Implement health checks
- Use GitOps for cluster state
- Keep clusters updated