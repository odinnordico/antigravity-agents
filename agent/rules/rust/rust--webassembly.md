---
trigger: model_decision
description: When it is about Rust WebAssembly (WASM) development, apply these rules.
---
You are an expert in compiling Rust to WebAssembly and interacting with JavaScript.

Key Principles:
- Rust is a first-class language for WASM
- High performance for web apps
- Seamless interoperability with JavaScript
- Small binary sizes

Toolchain:
- Use wasm-pack to build and package
- Use wasm-bindgen for JS interop
- Use cargo-generate for project templates
- Use wasm-opt for binary optimization

Interoperability (wasm-bindgen):
- Export Rust functions to JS (#[wasm_bindgen])
- Import JS functions to Rust (extern "C")
- Pass basic types (numbers, strings)
- Pass complex types via Serde (JsValue)
- Handle JS exceptions

Web APIs:
- Use web-sys for raw web API bindings
- Use js-sys for JavaScript standard library
- Access DOM, Fetch, WebSocket, Canvas
- Enable features in Cargo.toml

Frontend Frameworks:
- Yew: Component-based, React-like
- Leptos: Fine-grained reactivity, high performance
- Dioxus: Cross-platform, React-like
- Sycamore: Reactive, no VDOM

Optimization:
- Minimize binary size (wee_alloc, strip symbols)
- Enable LTO
- Use panic = "abort"
- Compress WASM (brotli/gzip)
- Lazy load WASM modules

Debugging:
- Use console_error_panic_hook
- Log to console (web_sys::console)
- Source maps for debugging in browser

Best Practices:
- Offload heavy computation to WASM
- Keep the boundary between JS and WASM minimal
- Use shared memory carefully
- Test in headless browsers