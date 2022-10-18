https://deepu.tech/memory-management-in-golang/

![image-20220330140740742](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220330140740742.png)

\https://blog.learngoprogramming.com/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed



## Page Heap(`mheap`)



This is where Go stores dynamic data(any data for which size cannot be calculated at compile time). This is the biggest block of memory and this is where **Garbage Collection(GC)** takes place.



The resident set is divided into pages of 8KB each and is managed by one global `mheap` object.

> Large objects(Object of Size > 32kb) are allocated directly from `mheap`. These large requests come at an expense of central lock, so only one `P`’s request can be served at any given point in time.



`mheap` manages pages grouped into different constructs as below







**mspan**: `mspan` is the most basic structure that manages the pages of memory in `mheap`. It’s a double-linked list that holds the address of the start page, span size class, and the number of pages in the span. Like TCMalloc, Go also divides Memory Pages into a block of 67 different classes by size starting at 8 bytes up to 32 kilobytes as in the below image



![image-20220330145903251](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220330145903251.png)



Each span exists twice, one for objects with pointers (**scan** classes) and one for objects with no pointers (**`noscan`** classes). This helps during GC as `noscan` spans need not be traversed to look for live objects.



Each span exists twice, one for objects with pointers (**scan** classes) and one for objects with no pointers (**`noscan`** classes). This helps during GC as `noscan` spans need not be traversed to look for live objects.



- **mcentral**: `mcentral` groups spans of same size class together. Each `mcentral` contains two `mspanList`:

  - **empty**: Double linked list of spans with no free objects or spans that are cached in a `mcache`. When a span here is freed, it’s moved to the nonempty list.
  - **non-empty**: Double linked list of spans with a free object. When a new span is requested from `mcentral`, it takes that from the nonempty list and moves it into the empty list.

  When `mcentral` doesn’t have any free span, it requests a new run of pages from `mheap`.



- **arena**: The heap memory grows and shrinks as required within the virtual memory allocated. When more memory is needed, `mheap` pulls them from the virtual memory as a chunk of 64MB(for 64-bit architectures) called `arena`. The pages are mapped to spans here.
- **mcache**: This is a very interesting construct. `mcache` is a cache of memory provided to a `P`(Logical Processor) to store small objects(Object size <=32Kb). Though this resembles the thread stack, it is part of the heap and is used for dynamic data. `mcache` contains `scan` and `noscan` types of `mspan` for all class sizes. Goroutines can obtain memory from `mcache` without any locks as a `P` can have only one `G` at a time. Hence this is more efficient. `mcache` requests new spans from `mcentral` when required.

## Stack

This is the stack memory area and there is one stack per Goroutine(`G`). This is where static data including function frames, static structs, primitive values, and pointers to dynamic structs are stored. This is not the same as `mcache` which is assigned to a `P`





# Go Memory management



Go’s memory management involves automatic allocation when memory is needed and garbage collection when memory is not needed anymore. It’s done by the standard library. Unlike C/C++ the developer does not have to deal with it and the underlying management done by Go is well optimized and efficient.



## Memory Allocation



Many programming languages that employ Garbage collection use a generational memory structure to make collection efficient along with compaction to reduce fragmentation. Go takes a different approach here, as we saw earlier, Go structures memory quite differently. Go employs a thread-local cache to speed up small object allocations and maintains `scan`/`noscan` spans to speed up GC. This structure along with the process avoids fragmentation to a great extent making compact unnecessary during GC. Let’s see how this allocation takes place.



Go decides the allocation process of an object based on its size and is divided into three categories:



**Tiny(size < 16B)**: Objects of size less than 16 bytes are allocated using the `mcache`’s tiny allocator. This is efficient and multiple tiny allocations are done on a single 16-byte block.



**Small(size 16B ~ 32KB)**: Objects of size between 16 bytes and 32 Kilobytes are allocated on the corresponding size class(`mspan`) on `mcache` of the `P` where the `G` is running.



In both tiny and small allocation if the `mspan`’s list is empty the allocator will obtain a run of pages from the `mheap` to use for the `mspan`. If the `mheap` is empty or has no page runs large enough then it allocates a new group of pages (at least 1MB) from the OS.



**Large(size > 32KB)**: Objects of size greater than 32 kilobytes are allocated directly on the corresponding size class of `mheap`. If the `mheap` is empty or has no page runs large enough then it allocates a new group of pages (at least 1MB) from the OS.

