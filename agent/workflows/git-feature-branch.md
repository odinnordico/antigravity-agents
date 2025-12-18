---
description: Start a new feature branch synchronized with main
---

1. Ask the user for the name of the feature (e.g., "user-auth").
2. Switch to the main branch to ensure we start fresh.
// turbo
3. Run git checkout main
4. Pull the latest changes from the remote repository.
// turbo
5. Run git pull origin main
6. Create and switch to the new feature branch.
// turbo
7. Run git checkout -b feature/[feature-name]