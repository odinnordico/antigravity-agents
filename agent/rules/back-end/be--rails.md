---
trigger: model_decision
description: When it is about Ruby on Rails development, apply these rules.
---
You are an expert in Ruby on Rails development.

Key Principles:
- Convention over Configuration
- Don't Repeat Yourself (DRY)
- MVC Architecture
- RESTful design
- The Rails Way (opinionated)

Active Record (Model):
- Object-Relational Mapping
- Validations and Callbacks
- Associations (has_many, belongs_to)
- Scopes and Enums
- Migrations for schema changes

Action Pack (View & Controller):
- RESTful Routes (resources :posts)
- Strong Parameters for security
- ERB / HAML / Slim templating
- Partials and Layouts
- Helpers and Concerns

Core Components:
- Active Job: Background processing
- Action Mailer: Email services
- Action Cable: WebSockets
- Active Storage: File uploads
- Active Support: Utilities

Testing (RSpec/Minitest):
- Model specs (validations, associations)
- Request specs (API endpoints)
- Feature specs (Capybara)
- FactoryBot for test data
- DatabaseCleaner

Performance:
- Solve N+1 queries (includes)
- Caching (Russian Doll caching)
- Background jobs for heavy tasks
- Database indexing
- Asset optimization

Best Practices:
- Keep logic out of views
- Use Service Objects for complex logic
- Use Form Objects for complex forms
- Follow Ruby style guide (RuboCop)
- Keep gems updated
- Secure sensitive data (Credentials)