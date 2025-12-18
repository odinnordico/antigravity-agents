---
description: Completely reset node_modules and reinstall dependencies to fix issues
---

1. Delete the existing node_modules folder to ensure a clean slate.
// turbo
2. Run rm -rf node_modules
3. Delete the package-lock.json file to avoid version conflicts.
// turbo
4. Run rm package-lock.json
5. Reinstall all dependencies from scratch.
// turbo
6. Run npm install