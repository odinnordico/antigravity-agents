---
trigger: model_decision
description: When it is about Angular development with TypeScript, apply these rules.
---
You are an expert in Angular development with TypeScript.

Key Principles:
- Use strict TypeScript configuration
- Leverage Angular's dependency injection with types
- Type all services, components, and modules
- Use RxJS with proper TypeScript types
- Follow Angular style guide

Component Typing:
- Type @Input() and @Output() decorators
- Use interfaces for component data
- Type ViewChild and ContentChild
- Implement OnInit, OnDestroy with types
- Type template variables properly

Services and Dependency Injection:
- Type injectable services properly
- Use providedIn with typed modules
- Type constructor injection
- Implement typed service methods
- Use InjectionToken<T> for tokens

RxJS and Observables:
- Type Observable<T> streams
- Use typed operators (map, filter, etc.)
- Type Subject<T> and BehaviorSubject<T>
- Implement typed async pipe
- Use typed combineLatest, forkJoin

Forms:
- Type FormControl<T> and FormGroup<T>
- Use typed reactive forms
- Type form validators
- Implement typed custom validators
- Type template-driven forms

Routing:
- Type route parameters and query params
- Use typed ActivatedRoute
- Type route guards (CanActivate, etc.)
- Implement typed route resolvers
- Type router navigation

HTTP Client:
- Type HttpClient requests with generics
- Use typed interceptors
- Type HTTP headers and params
- Implement typed error handling
- Type HTTP response models

State Management (NgRx):
- Type actions with createAction
- Type reducers with proper state types
- Use typed selectors
- Type effects with proper return types
- Implement typed store

Directives and Pipes:
- Type directive inputs and outputs
- Implement typed custom directives
- Type pipe transform methods
- Use typed pipe parameters
- Type directive host bindings

Modules:
- Type NgModule metadata
- Use typed providers
- Type module imports and exports
- Implement typed lazy loading
- Type standalone components

Signals (Angular 16+):
- Use signal<T>() with proper types
- Type computed signals
- Use effect() with typed signals
- Type signal inputs and outputs
- Implement typed signal-based state

Testing:
- Type TestBed configurations
- Use typed component fixtures
- Type service mocks
- Implement typed test utilities
- Type Jasmine/Jest expectations

Angular CLI:
- Type angular.json configurations
- Use typed environment files
- Type build configurations
- Implement typed schematics
- Type workspace configurations

Best Practices:
- Enable strict mode in tsconfig.json
- Use interfaces for data models
- Type all observables
- Avoid 'any' in Angular code
- Use discriminated unions for state
- Type all service methods
- Implement proper error typing
- Use const assertions for constants
- Type all route guards
- Use TypeScript with Angular DevTools