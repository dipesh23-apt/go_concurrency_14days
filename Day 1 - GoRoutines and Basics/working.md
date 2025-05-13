## How Goroutines Work Internally
✅ In Simple Words:
Goroutines are lightweight threads managed by Go runtime itself — not your operating system.

They are cheaper and faster to create than normal threads.

They are Multiplexed over a small number of OS threads

You can launch thousands of Goroutines without crashing your system.

✅ Analogy:
If OS threads are trucks, Goroutines are like bicycles — super cheap, small, fast, and you can have a lot of them.

✅ In Technical Terms:
Goroutines are "green threads" (user-level threads).

## Key Components:
G (Goroutine): Represents the actual goroutine (stack, function to run, etc.)

M (Machine): A wrapper around an OS thread.

P (Processor): Represents a logical processor that executes Go code. It holds the run queue of goroutines.

> 🧠 Conceptually: P picks a G (goroutine) and runs it on an M (OS thread). There’s usually a 1:1 relation between P and M, but multiple Gs can run on that M over time.

They run on M:N scheduler:

Many Goroutines (M) are mapped to fewer OS threads (N).

```
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("i:", i)
		}()
	}
	time.Sleep(1 * time.Second)
}
```
> Expected : 0 1 2 3 4

> Actual: 5 5 5 5 5 

All Goroutines share the same i, and by the time they run, the loop is done (i == 5).


### Programs
1. Write a simple program that prints numbers 1 to 10 inside a Goroutine. Use time.Sleep in main to ensure it completes. [solution](print_inside_go_routine/readme.md)
2. Launch two Goroutines:[solution](two_go_routines/readme.md)
          <br/> One prints even numbers (2 to 10). <br/> Other prints odd numbers (1 to 9).
3. Add synchronization to make sure the main program waits for Goroutines (using sync.WaitGroup).[solution](sol_3.go)
