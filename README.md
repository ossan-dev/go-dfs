# Go DFS

Repo to hold code for DFS implementation with Go.

## Performance Optimization

### 00 - Version

First draft that aims only for correctness. Now, we'll start by looking at some optimization we can do to remove bytes/op.

### 01 - Version

In this version, I've gone through the optimizations suggested by the Go compiler.

*I removed every "escape to the heap" situations.*

#### CPU

The CPU usage increased with the optimizations done. We started with ~1.127ns/op and we ended up with 1.367ns/op with the slice v1 implementation. This is also worse with the map implementation (1.683ns/op).

> From CPU perspective: the v0s were more performant. The winner (if we also consider the memory) was the v0 map implementation.

#### Memory

We improved a lot here. The slice implementation was less performant in v0. By v1 it improved a lot, even more than map implementation.

The file `00_slice_to_01_slice.txt` shows we improved the bytes/op by 47%. The allocations/op dropped by 25%.

> From memory perspective: the v1 version of Slice implementation has the lowest memory usage. The winner is the v1 of the Slice implementation.

### 02 - Version

In this version, I improved how I wrote the benchmarks of my application.
Plus, I used a slightly more efficient version since I'm passing to the `WalkFromNode` method a pre-allocated slice for both the courses and the queue to scan.
This allowed me to safely rely on the `append` built-in function without worrying about the extra-allocations due to missing space.

After these optimizations, we're left with the more performant solution, which is the `Slice` implementation.
