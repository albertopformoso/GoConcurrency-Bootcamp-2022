# Modeling Your Code: Communicating Sequential Processes

## Principal Problems when Communicating Sequential Processes

### Deadlocks

+ All concurrent process are waiting on one another
+ Will never recover without intervention

**The Coffman Conditions** *(these conditions will generate a deadlock)*
+ Mutual exclusion
+ Wait for a condition
+ No preemption: the state will not change unless there is an external intervention
+ Circular Wait

### Livelocks

+ Similar to a deadlock
+ Current processes are experforming, but not moving the state of the program forward

### Resource Starvation

+ When a concurrent process cannot get what it needs to perform work
+ Greedy workers vs Polite workers
+ Based on the critical sections of our program

## Determining Concurrency Safety
+ The most difficult aspect of developing concurrent code: people
+ The code most answer these three questions:
    + Who is responsible for the concurrency?
    + How is the problem mapped onto concurrency primitives?
    + Who is responsible for the synchronization?
+ Clarity vs performance

## The differnce between Concurrency and Parallelism
+ Concurrency is a property of the code
+ Parallelism is a property of the running program

## What is CSP?
+ Communicating Sequential Processes, invented by Charles Antony Richard Hoare
+ Inspiration for Go's concurrency models
+ Message-passing style with inputs and outputs

## Go's Phylosophy on Concurrency
+ Go supports two approaches for concurrency:
    + CSP
    + Traditional means
+ Between them, use whichever is most expenssive and/or simple
