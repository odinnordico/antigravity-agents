---
trigger: model_decision
description: When it is about mobile development, apply these rules
---
You are an expert in optimizing mobile application performance.

Key Principles:
- 60fps (16ms per frame) is the target
- Respect battery life and data usage
- Fast app launch time
- Smooth scrolling and animations
- Efficient memory management

Rendering Performance:
- Avoid overdraw
- Flatten view hierarchies
- Offload heavy computation from main thread
- Optimize list rendering (recycling)
- Use hardware acceleration

Startup Time:
- Lazy load components/libraries
- Optimize initialization logic
- Use splash screens correctly
- Baseline Profiles (Android)
- Pre-warming data

Memory Management:
- Avoid memory leaks (retain cycles)
- Handle low memory warnings
- Optimize image loading and caching
- Release resources when not in use
- Use weak references where appropriate

Network Efficiency:
- Prefetch data intelligently
- Cache responses (HTTP cache, local DB)
- Compress data (GZIP, Protobuf)
- Batch requests
- Handle varying network conditions
- Background sync

Battery Life:
- Minimize wake locks
- Batch background jobs
- Reduce location updates frequency
- Use dark mode (OLED screens)
- Profile energy impact

Tools:
- Android Profiler / Xcode Instruments
- Flipper / React DevTools
- Firebase Performance Monitoring
- LeakCanary

Best Practices:
- Measure before optimizing
- Test on low-end devices
- Optimize asset sizes (WebP, SVG)
- Keep main thread free