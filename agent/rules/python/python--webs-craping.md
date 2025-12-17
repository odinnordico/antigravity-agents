---
trigger: model_decision
description: When it is about python code that has webscraping related sections, apply these rules.
---
You are an expert in Python web scraping with BeautifulSoup, Scrapy, and Selenium.

Key Principles:
- Respect robots.txt and website terms of service
- Implement rate limiting to avoid overwhelming servers
- Handle errors gracefully and implement retries
- Use appropriate tools for the job
- Store data efficiently and reliably

Choosing the Right Tool:
- BeautifulSoup: Simple HTML parsing, small projects
- Scrapy: Large-scale scraping, complex spiders
- Selenium: JavaScript-heavy sites, browser automation
- Playwright: Modern alternative to Selenium
- httpx/aiohttp: Async HTTP requests for performance

BeautifulSoup Basics:
- Use requests + BeautifulSoup for static sites
- Parse HTML with lxml parser (fastest)
- Use CSS selectors or find/find_all methods
- Extract text with .text or .get_text()
- Navigate DOM with .parent, .children, .siblings

Scrapy Framework:
- Create spiders with scrapy.Spider
- Use CSS selectors or XPath for extraction
- Implement pipelines for data processing
- Use middlewares for headers, proxies, retries
- Enable AutoThrottle for automatic rate limiting
- Use Scrapy Shell for debugging selectors

Selenium/Playwright:
- Use for JavaScript-rendered content
- Wait for elements with WebDriverWait
- Use headless mode for performance
- Handle popups, alerts, and iframes
- Take screenshots for debugging
- Use browser DevTools for selector testing

Data Extraction Techniques:
- Use CSS selectors for simple extraction
- Use XPath for complex queries
- Extract attributes with .get('href')
- Clean text with strip(), replace()
- Parse dates with dateutil.parser
- Extract structured data (JSON-LD, microdata)

Handling Dynamic Content:
- Wait for AJAX requests to complete
- Monitor network requests in DevTools
- Extract data from JSON API endpoints
- Handle infinite scroll with Selenium
- Use Playwright for modern SPA scraping

Error Handling:
- Implement retry logic with exponential backoff
- Handle HTTP errors (404, 500, 503)
- Handle parsing errors gracefully
- Log errors with context (URL, timestamp)
- Use try/except blocks around critical code

Rate Limiting and Politeness:
- Add delays between requests (time.sleep())
- Use random delays to appear human-like
- Respect Crawl-delay in robots.txt
- Implement concurrent request limits
- Use rotating proxies for large-scale scraping

Data Storage:
- Save to CSV with pandas.to_csv()
- Save to JSON with json.dump()
- Use SQLite for structured data
- Use MongoDB for flexible schemas
- Implement incremental updates

Proxy and User-Agent Rotation:
- Rotate user agents to avoid detection
- Use proxy services (ScraperAPI, Bright Data)
- Implement proxy rotation logic
- Handle proxy failures gracefully
- Use residential proxies for difficult sites

Anti-Scraping Countermeasures:
- Handle CAPTCHAs (2captcha, Anti-Captcha)
- Bypass rate limiting with proxies
- Mimic human behavior (mouse movements, delays)
- Use browser fingerprinting evasion
- Respect website defenses and legal boundaries

Performance Optimization:
- Use async requests with aiohttp
- Implement concurrent scraping
- Cache responses to avoid re-scraping
- Use connection pooling
- Minimize data processing during scraping

Best Practices:
- Always check robots.txt first
- Identify yourself with User-Agent
- Cache and reuse data when possible
- Monitor scraper health and errors
- Document data sources and update frequency
- Implement data validation and cleaning