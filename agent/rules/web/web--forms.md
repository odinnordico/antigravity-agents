---
trigger: model_decision
description: When it is about web development with forms, apply these rules.  
---
You are an expert in web forms and form validation.

Key Principles:
- Implement proper form UX
- Use HTML5 validation attributes
- Provide clear error messages
- Ensure accessibility
- Validate on client and server

Form Structure:
- Use semantic HTML form elements
- Group related fields with fieldset
- Use legend for fieldset labels
- Associate labels with inputs
- Use appropriate input types
- Implement proper form hierarchy

Input Types:
- Use email for email inputs
- Use tel for phone numbers
- Use url for URLs
- Use number for numeric inputs
- Use date/time for temporal data
- Use search for search inputs
- Use color for color pickers

HTML5 Validation:
- Use required attribute
- Use pattern for regex validation
- Use min/max for numeric ranges
- Use minlength/maxlength for text
- Use step for numeric increments
- Implement custom validation messages

Client-Side Validation:
- Validate on blur and submit
- Provide real-time feedback
- Use Constraint Validation API
- Implement custom validators
- Show validation states visually
- Use JavaScript validation libraries

Server-Side Validation:
- Always validate on server
- Never trust client data
- Return clear error messages
- Validate data types and formats
- Implement rate limiting
- Log validation failures

Error Handling:
- Show errors near relevant fields
- Use clear, actionable messages
- Highlight invalid fields
- Provide error summaries
- Use aria-invalid and aria-describedby
- Don't clear valid fields on error

Form UX:
- Use clear, descriptive labels
- Provide helpful placeholder text
- Show password strength indicators
- Implement autocomplete attributes
- Use appropriate input modes
- Provide inline help text
- Show character counters

Accessibility:
- Associate labels with inputs
- Use fieldset and legend
- Implement proper ARIA attributes
- Ensure keyboard navigation
- Provide clear error messages
- Test with screen readers
- Use proper focus management

Form Submission:
- Disable submit during processing
- Show loading states
- Prevent double submission
- Handle network errors
- Provide success feedback
- Implement proper redirects

File Uploads:
- Validate file types and sizes
- Show upload progress
- Implement drag-and-drop
- Preview uploaded files
- Handle upload errors
- Use proper security measures

Multi-Step Forms:
- Show progress indicators
- Save progress automatically
- Allow navigation between steps
- Validate each step
- Provide step summaries
- Handle abandoned forms

Form Libraries:
- Use Formik for React forms
- Use React Hook Form for performance
- Use Zod for schema validation
- Use Yup for validation schemas
- Implement proper integration

Security:
- Implement CSRF protection
- Sanitize all inputs
- Use HTTPS for form submission
- Implement rate limiting
- Validate on server always
- Use secure session management

Best Practices:
- Keep forms simple and short
- Use appropriate input types
- Provide clear labels and instructions
- Validate on client and server
- Show validation feedback immediately
- Make error messages actionable
- Test with real users
- Ensure mobile-friendly forms
- Implement proper accessibility
- Monitor form completion rates