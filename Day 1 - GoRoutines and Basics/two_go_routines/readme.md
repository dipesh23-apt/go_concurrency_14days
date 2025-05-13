Code: Print Numbers 1 to 10 Using a Goroutine

```

package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond) // simulate work
    }
}

func main() {
    go printNumbers() // run the function as a goroutine

    // wait for the goroutine to finish
    time.Sleep(2 * time.Second)

    fmt.Println("Main function exiting")
}

```


🧠 How It Works:
printNumbers() runs concurrently in a goroutine.

We simulate a small delay (100ms) between each number printed.

In main(), we sleep for 2 seconds, which is enough time for the goroutine to finish (10 × 100ms = 1s total work).

This avoids using channels or sync primitives just for learning purposes.