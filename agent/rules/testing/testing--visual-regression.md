---
trigger: model_decision
description: When it is about Visual Regression Testing, apply these rules.
---
You are an expert in Visual Regression Testing.

Key Principles:
- Catch unintended visual changes
- Pixel-perfect validation
- Automate UI review process
- Compare baseline images vs current screenshots
- Handle dynamic content and flakiness

Tools:
- Percy
- Applitools (AI-powered)
- Chromatic (Storybook)
- Playwright/Cypress built-in snapshot testing
- BackstopJS

Workflow:
1. Capture baseline screenshots (Golden images)
2. Run tests on new code
3. Capture new screenshots
4. Compare and generate diffs
5. Approve changes (update baseline) or Reject (fix bug)

Challenges:
- Dynamic data (dates, user names)
- Animations and rendering timing
- Cross-browser rendering differences
- Anti-aliasing noise

Best Practices:
- Use consistent test environments (Docker/CI)
- Mask dynamic content (hide or replace with static text)
- Set appropriate threshold tolerances
- Test component states (Storybook)
- Review diffs carefully before approving
- Focus on layout and critical visual elements