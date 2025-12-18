---
trigger: model_decision
description: When it is about antigravity workflows, apply these rules.
---
You are an expert in creating Antigravity Workflows.

Key Principles:
- Workflows are step-by-step recipes for complex tasks
- Must be in .agent/workflows/ directory
- Must be .md (Markdown) files
- Must start with YAML frontmatter containing description
- Use turbo mode for safe, trusted commands

Workflow Structure:
```markdown
---
description: Short description of what this workflow does (max 250 chars)
---

1. First step with clear instruction
2. Second step
// turbo
3. Run `safe-command` (auto-runs with turbo)
4. Final step
```

Location Requirements:
- Workspace workflows: .agent/workflows/ in project root
- Global workflows: Available across all projects
- File naming: use lowercase with underscores (e.g., create_component.md)

Triggering Workflows:
- Smart Detection: AI detects relevant workflow from your question
- Slash Commands: Type /workflow-name to trigger directly
- Example: /deploy triggers .agent/workflows/deploy.md

Turbo Mode:
- // turbo: Auto-run single step (place above the step)
- // turbo-all: Auto-run ALL commands in workflow
- Only use for safe, non-destructive commands
- Great for setup scripts and trusted operations

Best Practices:
- Keep steps numbered and clear
- Ask user for input when needed
- Use placeholders like [ComponentName]
- Include code blocks for boilerplate
- Document prerequisites
- Test workflows before sharing