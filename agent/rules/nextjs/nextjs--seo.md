---
trigger: model_decision
description: When it is about Next.js SEO and metadata optimization, apply these rules.
---
You are an expert in Next.js SEO and metadata optimization.

Key Principles:
- Use Metadata API for SEO
- Implement dynamic metadata
- Use proper Open Graph tags
- Create XML sitemaps
- Implement structured data

Metadata API:
- Export metadata object from page/layout
- Use generateMetadata for dynamic metadata
- Set title, description, keywords
- Configure viewport and theme color
- Use metadata templates

Dynamic Metadata:
- Use generateMetadata async function
- Fetch data for metadata
- Use params for dynamic routes
- Implement fallback metadata
- Cache metadata generation

Title Optimization:
- Use descriptive, unique titles
- Keep titles under 60 characters
- Include primary keywords
- Use title templates
- Set default title in root layout

Meta Description:
- Write compelling descriptions
- Keep under 160 characters
- Include call-to-action
- Use unique descriptions per page
- Include target keywords naturally

Open Graph:
- Set og:title, og:description, og:image
- Use high-quality images (1200x630)
- Set og:type appropriately
- Include og:url and og:site_name
- Test with Facebook Debugger

Twitter Cards:
- Set twitter:card type
- Use twitter:title and twitter:description
- Set twitter:image
- Include twitter:creator
- Test with Twitter Card Validator

Canonical URLs:
- Set canonical in metadata
- Use for duplicate content
- Implement for pagination
- Use absolute URLs
- Handle www vs non-www

Structured Data:
- Use JSON-LD format
- Implement Schema.org types
- Add to page components
- Use for articles, products, events
- Test with Google Rich Results Test

Sitemap Generation:
- Create app/sitemap.ts
- Export default async function
- Return array of URLs
- Include lastModified dates
- Set changeFrequency and priority
- Generate dynamically from database

Robots.txt:
- Create app/robots.ts
- Export default function
- Allow/disallow paths
- Include sitemap URL
- Handle different environments

Image Optimization:
- Use next/image component
- Set alt text for all images
- Use descriptive filenames
- Implement lazy loading
- Use WebP format
- Set proper dimensions

Performance for SEO:
- Optimize Core Web Vitals
- Implement proper caching
- Use ISR for dynamic content
- Minimize JavaScript
- Optimize images and fonts

Mobile Optimization:
- Use responsive design
- Set viewport meta tag
- Test mobile usability
- Implement mobile-first approach
- Use touch-friendly elements

International SEO:
- Use hreflang tags
- Implement locale-specific URLs
- Create locale sitemaps
- Use proper language codes
- Set lang attribute on html

Best Practices:
- Use semantic HTML
- Implement proper heading hierarchy
- Use descriptive URLs
- Avoid duplicate content
- Use HTTPS
- Implement breadcrumbs
- Create quality content
- Monitor with Google Search Console
- Test with Lighthouse
- Update metadata regularly