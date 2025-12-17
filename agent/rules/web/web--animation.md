---
trigger: model_decision
description: When it is about web development for animations and motion design, apply these rules.
---
You are an expert in web animation and motion design.

Key Principles:
- Use animations purposefully
- Respect prefers-reduced-motion
- Optimize for performance
- Use appropriate animation techniques
- Follow motion design principles

CSS Animations:
- Use @keyframes for complex animations
- Use animation properties properly
- Implement timing functions
- Use animation-fill-mode
- Optimize with will-change
- Use animation-play-state for control

CSS Transitions:
- Use for simple state changes
- Transition specific properties
- Use appropriate durations
- Implement easing functions
- Avoid transitioning all properties
- Use transition-delay for sequencing

Transform:
- Use transform for better performance
- Use translate instead of position
- Use scale for size changes
- Use rotate for rotations
- Combine transforms efficiently
- Use transform-origin

Performance:
- Animate transform and opacity only
- Use will-change sparingly
- Avoid animating layout properties
- Use requestAnimationFrame
- Implement hardware acceleration
- Monitor frame rate

JavaScript Animation:
- Use Web Animations API
- Use requestAnimationFrame
- Implement proper timing
- Handle animation lifecycle
- Use libraries (GSAP, Anime.js)
- Implement cancellation

GSAP (GreenSock):
- Use for complex animations
- Implement timelines
- Use easing functions
- Implement scroll-triggered animations
- Use plugins for advanced features
- Optimize performance

Framer Motion:
- Use for React animations
- Implement variants
- Use layout animations
- Implement gesture animations
- Use AnimatePresence for exit animations
- Optimize re-renders

Scroll Animations:
- Use Intersection Observer
- Implement parallax effects
- Use scroll-triggered animations
- Implement smooth scrolling
- Use libraries (ScrollMagic, GSAP ScrollTrigger)
- Optimize performance

SVG Animation:
- Animate SVG properties
- Use CSS for simple animations
- Use JavaScript for complex animations
- Implement path animations
- Use SMIL (with fallbacks)
- Optimize SVG files

Loading Animations:
- Implement skeleton screens
- Use spinners appropriately
- Implement progress indicators
- Use shimmer effects
- Provide feedback during loading
- Optimize perceived performance

Micro-interactions:
- Implement button hover effects
- Use focus states
- Implement form feedback
- Use loading states
- Implement success/error states
- Keep animations subtle

Accessibility:
- Respect prefers-reduced-motion
- Provide alternatives to animations
- Don't rely on animation alone
- Avoid flashing content
- Implement proper focus management
- Test with assistive technologies

Motion Design Principles:
- Use easing for natural motion
- Implement proper timing
- Use choreography for sequences
- Follow 12 principles of animation
- Use motion to guide attention
- Maintain consistency

Best Practices:
- Use animations purposefully
- Keep animations short (< 500ms)
- Use appropriate easing
- Test on low-end devices
- Implement fallbacks
- Monitor performance
- Use CSS over JavaScript when possible
- Respect user preferences
- Document animation behavior
- Test across browsers