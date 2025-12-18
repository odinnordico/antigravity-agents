---
trigger: model_decision
description: When it is about Rust development, apply these rules.
---
You are an expert in Rust asynchronous programming using the Tokio runtime.

Key Principles:
- Async/await is syntactic sugar for Futures
- Futures do nothing unless polled
- Tokio provides the runtime to execute futures
- Avoid blocking the executor thread
- Use message passing over shared state

Tokio Runtime:
- Use #[tokio::main] macro for the entry point
- Use tokio::spawn to create new tasks (Green threads)
- Understand the difference between concurrency and parallelism
- Configure the runtime (worker threads, blocking threads)

Async/Await:
- Use async fn to define asynchronous functions
- Use .await to yield control until the future completes
- Futures are lazy; they must be awaited or spawned
- Handle cancellation by dropping the future

Synchronization:
- Use tokio::sync::Mutex for async-aware mutual exclusion
- Use tokio::sync::RwLock for read-heavy workloads
- Use tokio::sync::Semaphore for limiting concurrency
- Use tokio::sync::Notify for signaling
- Use tokio::sync::Barrier for synchronization points

Channels:
- Use tokio::sync::mpsc for multi-producer, single-consumer
- Use tokio::sync::oneshot for one-time values
- Use tokio::sync::broadcast for fan-out
- Use tokio::sync::watch for state changes

I/O Operations:
- Use tokio::fs for async file system operations
- Use tokio::net for async networking (TCP, UDP)
- Use tokio::io::AsyncRead and AsyncWrite traits
- Use tokio::time for sleeping and intervals

Control Flow:
- Use tokio::select! to wait on multiple branches
- Use tokio::join! to wait for all branches
- Use tokio::try_join! for fallible branches
- Handle timeouts with tokio::time::timeout

Streams:
- Streams are async iterators
- Use tokio_stream for stream utilities
- Use StreamExt traits (map, filter, fold)
- Iterate over streams with while let some = stream.next().await

Best Practices:
- Never call blocking code (std::thread::sleep, synchronous I/O) in async context
- Use tokio::task::spawn_blocking for CPU-bound or blocking operations
- Pin futures when necessary (Box::pin)
- Use tracing for async logging and diagnostics
- Test async code with #[tokio::test]