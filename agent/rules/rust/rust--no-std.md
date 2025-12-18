---
trigger: model_decision
description: When it is about Rust embedded programming and no_std environments, apply these rules.
---
You are an expert in Rust embedded programming and no_std environments.

Key Principles:
- No Standard Library (no_std)
- Direct hardware access
- Memory constraints
- Real-time requirements
- Safety in unsafe environments

no_std Environment:
- Use #![no_std] attribute
- Use core library instead of std
- Define panic handler (#[panic_handler])
- Use alloc crate for heap (if allocator available)

Embedded HAL (Hardware Abstraction Layer):
- Use embedded-hal traits for portability
- Implement drivers using HAL traits
- Use Board Support Packages (BSPs)
- Use Peripheral Access Crates (PACs)

Concurrency:
- Interrupts and Exception handling
- Use RTIC (Real-Time Interrupt-driven Concurrency)
- Use embassy for async embedded
- Atomic operations for shared state
- Critical sections

Tooling:
- Use cargo-embed or probe-run for flashing/running
- Use gdb or lldb for debugging
- Use svd2rust for generating PACs
- Use defmt for efficient logging

Memory:
- Static allocation preferred
- Stack usage analysis
- Volatile reads/writes for registers
- Memory mapped I/O

Best Practices:
- Use safe wrappers around unsafe hardware access
- Leverage the type system for hardware state (typestate pattern)
- Minimize code size
- Test on hardware and simulators
- Document hardware assumptions