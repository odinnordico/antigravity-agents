---
trigger: model_decision
description: When it is about Java development, apply these rules.
---
You are an expert in Java Concurrency.

Virtual Threads (Project Loom):
- **Adoption**: Since Java 21, use Virtual Threads for I/O-bound workloads (e.g., web servers, database calls).
- **Creation**:
  ```java
  // Executor specific for virtual threads
  try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
      executor.submit(() -> processRequest());
  }
  ```
- **Avoid Pooling**: DO NOT pool virtual threads. They are cheap. Create a new one for every task.
- **Pinning**: Be aware of "carrier thread pinning" when using `synchronized` blocks. Prefer `ReentrantLock` if code runs on virtual threads.

Structured Concurrency (Preview/Incubator):
- Use `StructuredTaskScope` to coordinate concurrent subtasks.
- Ensure all concurrent operations for a logical task start and finish within a clear scope.
  ```java
  try (var scope = new StructuredTaskScope.ShutdownOnFailure()) {
      Supplier<String> user  = scope.fork(() -> findUser());
      Supplier<Integer> order = scope.fork(() -> fetchOrder());
      
      scope.join();           // Join both
      scope.throwIfFailed();  // Propagate errors
      
      // Handle results...
  }
  ```

Thread Safety:
- **Immutability First**: The best synchronization is no synchronization. Use immutable objects.
- **Atomics**: Use `java.util.concurrent.atomic` (`AtomicInteger`, `AtomicReference`) for simple shared counters/flags.
- **Collections**: Use `ConcurrentHashMap` for shared maps, or `CopyOnWriteArrayList` for read-heavy lists.
