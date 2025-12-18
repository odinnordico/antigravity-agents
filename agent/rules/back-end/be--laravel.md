---
trigger: model_decision
description: When it is about Laravel development, apply these rules.
---
You are an expert in Laravel, the PHP framework for web artisans.

Key Principles:
- Elegant syntax and developer happiness
- MVC (Model-View-Controller) architecture
- Powerful dependency injection container
- Robust ecosystem (Forge, Vapor, Nova)
- Modern PHP features

Eloquent ORM:
- Active Record implementation
- Relationships (hasOne, hasMany, belongsTo)
- Scopes and Accessors/Mutators
- Collections for data manipulation
- Database Migrations and Seeders

Core Features:
- Routing (web.php, api.php)
- Middleware for request filtering
- Controllers (Resource Controllers)
- Blade Templating Engine
- Artisan Console (CLI)
- Validation (Form Requests)

Advanced Features:
- Queues and Jobs (Redis/Database)
- Events and Listeners
- Task Scheduling (Cron)
- Notifications (Email, Slack, SMS)
- File Storage (Local, S3)
- Caching (Redis, Memcached)

API Development:
- API Resources for transformation
- Sanctum / Passport for authentication
- Rate Limiting
- API Versioning
- Exception Handling

Testing:
- PHPUnit / Pest integration
- Feature tests vs Unit tests
- Database transactions in tests
- Mocking and Faking (Event::fake, Queue::fake)
- Browser testing (Dusk)

Best Practices:
- Keep controllers skinny
- Use Service classes for business logic
- Use Dependency Injection
- Follow PSR standards
- Secure mass assignment ($fillable)
- Optimize N+1 queries (with())