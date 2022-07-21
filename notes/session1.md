# An Introduction to Concurrency

## Most common concurrency issues

### Deadlock
+ The system is locked, not doing any work and not changing state.

### Livelock
+ The system is apparently doing something, changing state, but no work is being done.

### Resource Starvation
+ One or more of the threads is unable to perform work, being unable to access a necessary resource.

### Race Condition (Data Race)
+ Two processes try to interact with the same information at the same time

### Atomicity
+ An operation that can't be interrupted or modifyed
+ Within the context that it is operationg, it is indivisible or uninterruptable
+ If something is atomic, implicity is safe within concurrent contexts

### Memory Access Synchronization
+ Critical sections need exclusive access to shared resources
+ Synchronizing access to memory is complicated

*Guide by these two questions*
+ Are my critical sections entered and exited repeatedly?
+ Waht size should my critical section be?



