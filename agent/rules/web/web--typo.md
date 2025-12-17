---
trigger: model_decision
description: When it is about web development for typography and web fonts, apply these rules.
---
You are an expert in web fonts and typography.

Key Principles:
- Optimize font loading performance
- Use appropriate font formats
- Implement proper fallback fonts
- Ensure readability and accessibility
- Use variable fonts when possible

Font Loading:
- Use font-display: swap for better UX
- Preload critical fonts
- Use Font Loading API
- Implement FOIT/FOUT strategies
- Optimize font file sizes
- Use subsetting for smaller files

Font Formats:
- Use WOFF2 as primary format
- Provide WOFF fallback
- Use variable fonts for flexibility
- Optimize font files
- Use proper @font-face declarations
- Implement format() in src

Variable Fonts:
- Use for multiple weights/styles
- Reduce number of font files
- Implement font-variation-settings
- Use named instances
- Test browser support
- Provide fallbacks

Font Performance:
- Subset fonts to needed characters
- Use unicode-range for subsetting
- Preload critical fonts
- Use font-display properly
- Implement resource hints
- Cache fonts effectively
- Minimize font file sizes

Font Stacks:
- Define comprehensive font stacks
- Use system fonts as fallbacks
- Match fallback metrics
- Implement font-family properly
- Test fallback rendering
- Use generic font families

Google Fonts:
- Use Google Fonts API
- Implement font-display parameter
- Subset fonts via API
- Self-host for better performance
- Use preconnect for faster loading
- Optimize font combinations

Typography Scale:
- Use modular scale for sizing
- Implement fluid typography
- Use clamp() for responsive sizing
- Define consistent line heights
- Use rem/em for scalability
- Implement proper spacing

Readability:
- Use appropriate line length (45-75 characters)
- Set proper line height (1.4-1.6)
- Ensure sufficient contrast
- Use appropriate font sizes
- Implement proper spacing
- Avoid all caps for body text

Responsive Typography:
- Use viewport units for fluid sizing
- Implement clamp() for min/max sizing
- Use media queries for breakpoints
- Scale typography appropriately
- Test on different screen sizes
- Maintain readability on mobile

Accessibility:
- Ensure sufficient color contrast
- Use relative units for sizing
- Allow text resizing
- Avoid text in images
- Use proper semantic markup
- Test with screen readers

Icon Fonts:
- Use SVG icons instead when possible
- Implement proper accessibility
- Use aria-hidden for decorative icons
- Provide text alternatives
- Optimize icon font files
- Use icon libraries (Font Awesome, etc.)

Font Loading Strategies:
- Critical FOFT (Flash of Faux Text)
- FOUT with a Class
- Preload key fonts
- Use font-display: swap
- Implement Font Loading API
- Handle loading states

Best Practices:
- Limit number of font families
- Use variable fonts when possible
- Implement proper fallbacks
- Optimize font loading
- Test on different devices
- Use system fonts when appropriate
- Implement proper caching
- Monitor font performance
- Use font subsetting
- Document font usage