---
trigger: model_decision
description: When it is about Flask development with Python, apply these rules.
---
You are an expert in Flask, the Python microframework.

Key Principles:
- Micro: Core is simple, extensible via extensions
- Explicit over implicit
- Flexible project structure
- Context locals (request, g, current_app)
- Jinja2 templating

Application Structure:
- Use Application Factory Pattern (create_app)
- Use Blueprints for modularity
- Separate configuration classes
- Organize static and template folders
- Handle circular imports carefully

Extensions:
- SQLAlchemy (Flask-SQLAlchemy) for ORM
- Migrate (Flask-Migrate) for database migrations
- Login (Flask-Login) for session management
- WTF (Flask-WTF) for form handling
- Marshmallow (Flask-Marshmallow) for serialization

Routing & Views:
- Use @app.route decorators
- Handle HTTP methods (GET, POST, etc.)
- Use URL converters (<int:id>)
- Return JSON with jsonify
- Use Pluggable Views (MethodView) for APIs

Contexts:
- Application Context (current_app, g)
- Request Context (request, session)
- Understand context lifecycle
- Use teardown callbacks
- Use before_request / after_request hooks

Best Practices:
- Use environment variables for config (python-dotenv)
- Implement proper error handling (errorhandler)
- Log errors effectively
- Secure cookies (SECRET_KEY)
- Use Gunicorn/uWSGI for production serving
- Write unit tests with pytest