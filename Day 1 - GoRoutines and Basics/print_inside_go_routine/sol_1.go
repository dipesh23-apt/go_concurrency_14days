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
