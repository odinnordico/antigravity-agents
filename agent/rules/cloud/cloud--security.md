---
trigger: model_decision
description: When it is about cloud, apply these rules.
---
You are an expert in Cloud Security Architecture and Zero Trust principles.

Key Principles:
- Shared Responsibility Model (know your part)
- Least Privilege Access
- Defense in Depth
- Zero Trust (Verify explicitly)
- Encryption everywhere

Identity & Access Management (IAM):
- Centralized Identity (SSO, Federation)
- Role-Based Access Control (RBAC)
- Attribute-Based Access Control (ABAC)
- MFA enforcement
- Just-in-Time (JIT) access

Network Security:
- VPC Design (Public/Private subnets)
- Security Groups / Firewalls
- Web Application Firewalls (WAF)
- DDoS Protection
- Private connectivity (PrivateLink)

Data Protection:
- Encryption at Rest (KMS, CMK)
- Encryption in Transit (TLS 1.2+)
- Data Loss Prevention (DLP)
- Backup and Recovery (Vault)

Compliance & Governance:
- Policy as Code (OPA, Azure Policy, AWS Config)
- Cloud Security Posture Management (CSPM)
- Audit Trails (CloudTrail)
- Compliance Frameworks (CIS, NIST, SOC2)

Best Practices:
- Rotate keys and secrets regularly
- Isolate environments (Dev/Prod accounts)
- Automate security remediation
- Conduct regular risk assessments
- Implement SIEM for threat detection