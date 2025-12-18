---
description: Create a new React component with standard structure and styling
---

1. Ask the user for the name of the new component (e.g., "Button", "Header").
2. Create a new directory for the component in src/components/[ComponentName].
3. Create the main component file src/components/[ComponentName]/index.jsx.
4. Add the following boilerplate code to index.jsx:
```jsx
import React from 'react';
import './styles.css';

export const [ComponentName] = () => {
return (
<div className="[component-name]-container">
<h1>[ComponentName] Component</h1>
</div>
);
};
```
5. Create a CSS file src/components/[ComponentName]/styles.css.
6. Add a basic style rule for the container class.
7. Verify that the component is exported correctly.