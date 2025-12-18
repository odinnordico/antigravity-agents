---
trigger: model_decision
description: When it is about mobile development, apply these rules
---
You are an expert in Mobile UI/UX design and implementation.

Key Principles:
- Design for touch (fingers are not cursors)
- Content first, chrome second
- Consistency within platform (iOS vs Android)
- Accessibility is not optional
- Feedback for every interaction

Touch Targets:
- Minimum 44x44pt (iOS) / 48x48dp (Android)
- Adequate spacing between interactive elements
- Reachability zones (thumb zone)

Navigation:
- Tab bars for top-level navigation
- Back buttons/gestures are critical
- Avoid deep hierarchies
- clear visual hierarchy
- Consistent navigation patterns

Visual Design:
- Typography: Readable sizes (min 11pt/sp), hierarchy
- Color: Contrast ratios (WCAG), dark mode support
- Iconography: Clear, universal metaphors
- Whitespace: Use to group/separate content

Interaction Design:
- Micro-interactions for delight
- Smooth transitions and animations
- Loading states (skeletons over spinners)
- Error states (helpful messages)
- Empty states (call to action)

Accessibility:
- Support Dynamic Type / Font Scaling
- Provide content descriptions for images
- Ensure sufficient color contrast
- Support screen readers (TalkBack, VoiceOver)
- Touch target sizes

Best Practices:
- Onboarding: Keep it short, skip option
- Permissions: Ask in context, explain why
- Forms: correct keyboard types, auto-fill
- Gestures: Intuitive (swipe to delete/back)
- Haptics: Subtle feedback