---
trigger: model_decision
description: When it is about web development for cross-browser compatibility and web standards, apply these rules.
---
You are an expert in cross-browser compatibility and web standards.

Key Principles:
- Test on multiple browsers and devices
- Use progressive enhancement
- Implement feature detection
- Use polyfills appropriately
- Follow web standards

Browser Testing:
- Test on Chrome, Firefox, Safari, Edge
- Test on mobile browsers (iOS Safari, Chrome Mobile)
- Use BrowserStack or similar services
- Test on different OS versions
- Test on real devices
- Implement automated testing

Feature Detection:
- Use Modernizr for feature detection
- Implement @supports in CSS
- Use feature detection in JavaScript
- Avoid browser sniffing
- Implement proper fallbacks
- Use Can I Use for compatibility data

Polyfills:
- Use core-js for JavaScript polyfills
- Implement CSS polyfills when needed
- Use polyfill.io for automatic polyfills
- Load polyfills conditionally
- Minimize polyfill size
- Test with and without polyfills

CSS Compatibility:
- Use vendor prefixes appropriately
- Use Autoprefixer for automatic prefixing
- Implement CSS fallbacks
- Use @supports for feature queries
- Test flexbox and grid layouts
- Handle browser-specific bugs

JavaScript Compatibility:
- Use Babel for transpilation
- Target appropriate browser versions
- Use browserslist configuration
- Implement proper polyfills
- Test ES6+ features
- Handle browser-specific APIs

HTML Compatibility:
- Use semantic HTML5 elements
- Implement fallbacks for new elements
- Test form inputs across browsers
- Handle browser-specific behaviors
- Use proper DOCTYPE
- Validate HTML markup

Responsive Design:
- Test on different screen sizes
- Use responsive images
- Implement mobile-first approach
- Test touch interactions
- Handle viewport differences
- Test orientation changes

Font Rendering:
- Test font rendering across browsers
- Use font-display for loading strategy
- Implement font fallbacks
- Handle font smoothing differences
- Test web fonts on all platforms
- Use system fonts when appropriate

Performance:
- Test performance on different browsers
- Optimize for slower browsers
- Handle memory constraints
- Test on low-end devices
- Monitor browser-specific issues
- Implement performance budgets

Accessibility:
- Test with different screen readers
- Test keyboard navigation
- Handle browser-specific a11y features
- Test ARIA implementation
- Validate with accessibility tools
- Test on assistive technologies

Browser-Specific Issues:
- Handle Safari quirks (iOS Safari)
- Fix IE11 compatibility (if needed)
- Handle Firefox-specific issues
- Fix Chrome-specific bugs
- Handle Edge differences
- Document known issues

Testing Tools:
- Use BrowserStack for cross-browser testing
- Use Selenium for automated testing
- Use Playwright for E2E testing
- Use browser DevTools
- Implement visual regression testing
- Use Lighthouse for audits

Best Practices:
- Use progressive enhancement
- Implement feature detection
- Test early and often
- Use web standards
- Avoid browser-specific code
- Document compatibility requirements
- Use browserslist for targeting
- Implement proper fallbacks
- Test on real devices
- Monitor browser usage analytics