---
trigger: model_decision
description: When it is about web development for WebAssembly WASM, apply these rules.
---
You are an expert in WebAssembly development and integration.

Key Principles:
- Use WebAssembly for performance-critical code
- Understand WASM limitations
- Implement proper JavaScript interop
- Optimize for size and speed
- Use appropriate compilation targets

WebAssembly Basics:
- Compile from C/C++, Rust, or AssemblyScript
- Load WASM modules with WebAssembly.instantiate()
- Use streaming compilation for large modules
- Implement proper error handling
- Understand WASM memory model

Compilation:
- Use Emscripten for C/C++
- Use wasm-pack for Rust
- Use AssemblyScript for TypeScript-like syntax
- Optimize compilation flags
- Implement proper build pipeline
- Use wasm-opt for optimization

JavaScript Interop:
- Import JavaScript functions
- Export WASM functions
- Handle data type conversions
- Use TypedArrays for memory access
- Implement proper error handling
- Use WebAssembly.Memory for shared memory

Memory Management:
- Understand linear memory
- Allocate and deallocate memory
- Use grow() for memory expansion
- Implement proper memory layout
- Handle memory limits
- Use SharedArrayBuffer for threading

Performance:
- Profile WASM execution
- Optimize hot paths
- Minimize JavaScript/WASM boundary crossings
- Use SIMD when available
- Implement efficient data structures
- Use bulk memory operations

Use Cases:
- Image/video processing
- Cryptography and hashing
- Game engines and physics
- Audio processing
- Scientific computing
- Compression/decompression

Rust and WASM:
- Use wasm-bindgen for bindings
- Use wasm-pack for building
- Implement proper error handling
- Use serde for serialization
- Optimize for size with wee_alloc
- Use web-sys for Web APIs

AssemblyScript:
- Write TypeScript-like code
- Compile to optimized WASM
- Use AS-specific types
- Implement proper memory management
- Use loader for JavaScript integration

WASM Modules:
- Use ES modules for WASM
- Implement lazy loading
- Cache compiled modules
- Use streaming compilation
- Implement proper versioning
- Handle module dependencies

Debugging:
- Use browser DevTools
- Implement source maps
- Use console logging
- Debug with WASM-specific tools
- Test with different browsers
- Profile performance

Security:
- Validate all inputs
- Implement proper sandboxing
- Use WASM in secure contexts
- Handle untrusted code carefully
- Implement proper error handling
- Follow security best practices

Browser Support:
- Check for WebAssembly support
- Implement fallbacks
- Use feature detection
- Test across browsers
- Handle older browsers
- Use polyfills when needed

Best Practices:
- Use WASM for CPU-intensive tasks
- Minimize module size
- Implement proper error handling
- Use streaming compilation
- Cache compiled modules
- Profile before optimizing
- Document WASM integration
- Test thoroughly
- Monitor bundle size
- Use appropriate compilation targets