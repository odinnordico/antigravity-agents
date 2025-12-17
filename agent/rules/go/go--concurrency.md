---
trigger: model_decision
description: When it is about go code that has concurrency or uses goroutines or channels, apply these rules.
---
You are an expert in Go concurrency, goroutines, and channels.

Key Principles:
- Share memory by communicating, don't communicate by sharing memory
- Use goroutines for concurrent execution
- Use channels for synchronization and data transfer
- Handle context for cancellation
- Prevent race conditions and deadlocks

Goroutines:
- Start goroutines with 'go' keyword
- Keep goroutines lightweight
- Manage goroutine lifecycle
- Avoid leaking goroutines
- Use WaitGroup to wait for completion

Channels:
- Use unbuffered channels for synchronization
- Use buffered channels for throughput
- Close channels from the sender side
- Use range to iterate over channels
- Use select for multiplexing channels

Synchronization:
- Use sync.Mutex for critical sections
- Use sync.RWMutex for read-heavy data
- Use sync.Once for one-time initialization
- Use sync.Cond for signaling
- Use atomic package for simple counters

Context:
- Pass context.Context as the first argument
- Use context for cancellation propagation
- Use context for timeouts and deadlines
- Use context values sparingly
- Always cancel contexts to release resources

Patterns:
- Worker Pool: Distribute work among workers
- Pipeline: Chain stages of processing
- Fan-out/Fan-in: Distribute and aggregate work
- Generator: Produce data in a goroutine
- Semaphore: Limit concurrency

Error Handling:
- Propagate errors through channels
- Use errgroup for group error handling
- Handle panics in goroutines
- Log errors in background tasks
- Cancel operations on error

Race Detection:
- Always run tests with -race
- Fix data races immediately
- Use atomic operations for shared counters
- Protect shared maps with mutexes
- Avoid concurrent read/write to same variable

Best Practices:
- Don't leave goroutines hanging
- Close channels gracefully
- Use select with default for non-blocking
- Limit the number of goroutines
- Use buffered channels carefully
- Design for cancellation
- Test concurrent code thoroughly