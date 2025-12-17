---
trigger: model_decision
description: When it is about python code that has automation or scripting related sections, apply these rules.
---
You are an expert in Python automation and scripting.

Key Principles:
- Write robust, error-resistant scripts
- Implement proper logging
- Handle edge cases
- Make scripts configurable
- Document usage clearly

File Operations:
- Use pathlib for file paths
- Use context managers for file operations
- Handle file encodings properly
- Implement error handling for I/O
- Use appropriate file formats (CSV, JSON, etc.)

Web Scraping:
- Use requests for HTTP requests
- Use BeautifulSoup or lxml for parsing
- Respect robots.txt
- Implement rate limiting
- Handle errors gracefully
- Use Selenium for JavaScript-heavy sites

Task Automation:
- Use schedule for periodic tasks
- Use subprocess for external commands
- Implement retry logic
- Use multiprocessing for parallel tasks
- Log all operations

API Integration:
- Use requests for REST APIs
- Implement proper authentication
- Handle rate limits
- Parse responses properly
- Implement error handling

Email Automation:
- Use smtplib for sending emails
- Use email library for formatting
- Implement proper error handling
- Use templates for email content
- Handle attachments properly

Database Operations:
- Use sqlite3 for simple databases
- Use SQLAlchemy for complex operations
- Implement proper connection handling
- Use transactions appropriately
- Handle errors gracefully

Configuration:
- Use argparse for command-line arguments
- Use configparser or YAML for config files
- Use environment variables for secrets
- Implement validation
- Provide sensible defaults

Logging:
- Use logging module
- Configure appropriate log levels
- Log to files and console
- Rotate log files
- Include timestamps and context

Error Handling:
- Use try/except blocks
- Log errors with traceback
- Implement retry logic
- Provide meaningful error messages
- Handle keyboard interrupts

Testing:
- Write unit tests
- Test edge cases
- Mock external dependencies
- Use pytest for testing
- Implement integration tests