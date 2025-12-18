---
description: Create a new custom React hook with standard boilerplate
---

1. Ask the user for the name of the hook (must start with "use", e.g., "useWindowSize").
2. Create a new file in src/hooks/[HookName].js.
3. Add the following boilerplate code to the file:
```javsctipt
import { useState, useEffect } from 'react';

export const [HookName] = () => {
const [data, setData] = useState(null);

useEffect(() => {
// Logic goes here
console.log('[HookName] mounted');
}, []);

return { data };
};
```
4. Verify that the hook is exported as a named export.