### Repository Description: Lock-Free Data Structures in Go - Benchmarks and Tests

This repository contains implementations of **lock-free data structures** in Go, along with comprehensive **benchmarks** and **tests** to evaluate their performance and correctness. Lock-free data structures are designed to handle concurrent access without using traditional locking mechanisms (e.g., mutexes), making them ideal for high-performance, multi-threaded applications.

The repository focuses on the following lock-free structures:
1. **Lock-Free Stack**
2. **Lock-Free Ring Buffer**
3. **Lock-Free Queue** (optional, if implemented)
4. **Lock-Free LRU cache**
5. **Cache of sync.Map** non lock-free 

Each implementation is rigorously tested for thread safety and correctness, and benchmarked to compare performance against traditional locking-based approaches.

---

### Key Features:
- **Lock-Free Implementations**: All data structures are implemented using atomic operations (`sync/atomic`) to ensure thread safety without locks.
- **Benchmarks**: Detailed benchmarks are provided to measure the performance of each lock-free structure under various levels of concurrency.
- **Tests**: Unit tests and stress tests are included to verify the correctness and robustness of the implementations.
- **Documentation**: Clear explanations of each data structure, its use cases, and performance characteristics are provided.

---

### Implementations:
1. **Lock-Free Stack**:
   - A simple stack that supports `Push` and `Pop` operations.
   - Uses `atomic.CompareAndSwapPointer` for lock-free synchronization.

2. **Lock-Free Ring Buffer**:
   - A fixed-size circular buffer for producer-consumer scenarios.
   - Uses atomic operations to manage producer and consumer indices.

3. **Lock-Free Queue**:
   - A queue implementation that supports concurrent `Enqueue` and `Dequeue` operations.
   - Uses atomic operations to manage head and tail pointers.

4. **Lock-Free LRU cache**:
   - A LRU cache that supports concurrent `Set`, `Pop`, and `Get` operations.
   - Uses atomic operations to manage node pointers.

5. **Cache over sync.Map**:
   - Non lock-free cache implementation
   - A cache implementation that supports concurrent `Set`, `Get`, `Delete` and `Range`.

---

### Benchmarks:
The benchmarks compare the performance of lock-free implementations against traditional locking-based approaches (e.g., using `sync.Mutex` or `sync.RWMutex`). Metrics include:
- **Throughput**: Number of operations per second.
- **Latency**: Time taken for individual operations.
- **Scalability**: Performance under increasing levels of concurrency.

Example benchmark results:
```
goos: windows
goarch: amd64
pkg: AlekseyPromet/algo/lockfree
cpu: AMD Ryzen 7 7840HS w/ Radeon 780M Graphics
=== RUN   BenchmarkNewRingBufferFull
BenchmarkNewRingBufferFull
BenchmarkNewRingBufferFull-16
18812934                64.84 ns/op           23 B/op          1 allocs/op
```

---

### Tests:
The tests ensure that each lock-free implementation is correct and thread-safe. They include:
- **Concurrency Tests**: Stress tests with multiple goroutines to verify correctness under heavy concurrent access.
- **Edge Cases**: Tests for edge cases, such as empty/full buffers or stacks.

Example test output:
```
PASS: TestLockFreeRingBuffer_ConcurrentPushPop
PASS: TestLockFreeStack_ConcurrentPushPop
```

---

### How to Use:
1. Clone the repository:
   ```bash
   git clone https://github.com/AlekseyPromet/lockfree.git
   cd lockfree
   ```
2. Run the tests:
   ```bash
   go test -v ./...
   ```
3. Run the benchmarks:
   ```bash
   go test -bench=. ./...
   ```

---

### Dependencies:
- Recomended Go 1.18+
- `sync/atomic` package for atomic operations.

---

### Contributions:
Contributions are welcome! If you'd like to add new lock-free data structures, improve existing implementations, or add more benchmarks and tests, feel free to open a pull request.

---

### License:
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

### Why Lock-Free?
Lock-free data structures are ideal for high-performance applications where:
- Low latency is critical.
- High concurrency is required.
- Traditional locking mechanisms (e.g., mutexes) introduce contention and bottlenecks.

By eliminating locks, these structures can achieve better scalability and performance in multi-threaded environments.

---

Explore the repository, run the benchmarks, and see how lock-free data structures can improve the performance of your Go applications! 🚀