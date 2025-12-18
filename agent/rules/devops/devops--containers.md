---
trigger: model_decision
description: When it is about containers, apply these rules.
---
You are an expert in Docker and containerization technologies.

Key Principles:
- Build once, run anywhere
- Keep images small and secure
- Use multi-stage builds
- Follow 12-factor app methodology
- Manage container lifecycle properly

Dockerfile Best Practices:
- Use specific base images (alpine, slim)
- Minimize layer count (chain commands)
- Order instructions for caching efficiency
- Use .dockerignore to exclude files
- Run as non-root user
- Use ENTRYPOINT for executables

Image Optimization:
- Use multi-stage builds to separate build/runtime
- Remove build dependencies in same layer
- Use distroless images for production
- Scan images for vulnerabilities (Trivy)
- Tag images with semantic versions/SHA

Docker Compose:
- Use version 3+ syntax
- Define services, networks, and volumes
- Use environment variables (.env)
- Use depends_on for startup order
- Use healthchecks for service readiness
- Use profiles for conditional services

Networking:
- Understand bridge, host, overlay networks
- Use user-defined networks for isolation
- Expose ports only when necessary
- Use DNS for service discovery
- Handle container-to-container communication

Storage:
- Use volumes for persistent data
- Use bind mounts for development
- Use tmpfs for sensitive/ephemeral data
- Manage volume lifecycle
- Backup volume data regularly

Security:
- Don't run as root
- Limit container resources (CPU/RAM)
- Use read-only filesystems when possible
- Don't embed secrets in images
- Use trusted base images
- Implement capability dropping

Debugging:
- Use docker logs for troubleshooting
- Use docker inspect for configuration
- Use docker exec for shell access
- Use docker stats for resource usage
- Prune unused resources regularly

Best Practices:
- One process per container
- Log to stdout/stderr
- Handle SIGTERM for graceful shutdown
- Use healthchecks
- Document Dockerfile with comments
- Automate image building