## How Goroutines Work Internally
✅ In Simple Words:
Goroutines are lightweight threads managed by Go itself — not your operating system.

They are cheaper and faster to create than normal threads.

You can launch thousands of Goroutines without crashing your system.

✅ Analogy:
If OS threads are trucks, Goroutines are like bicycles — super cheap, small, fast, and you can have a lot of them.

✅ In Technical Terms:
Goroutines are "green threads" (user-level threads).

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