// And for synchronizing multiple goroutines

package main

import "fmt"
import "sync"

var wg sync.WaitGroup // 1

func routine(i int) {
    defer wg.Done() // 3
    fmt.Printf("routine %v finished\n", i)
}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1) // 2
        go routine(i) // *
    }
    wg.Wait() // 4
    fmt.Println("main finished")
}

// WaitGroup usage in order of execution.

//1  Declaration of global variable. Making it global is the easiest way to make it visible to all functions and methods.
//2  Increasing the counter. This must be done in main goroutine because there is no guarantee that newly started goroutine will execute before 4 due to memory model guarantees.
//3  Decreasing the counter. This must be done at the exit of goroutine. Using deferred call, we make sure that it will be called whenever function ends no matter but no matter how it ends.
//4  Waiting for the counter to reach 0. This must be done in main goroutine to prevent program exit.

//* The actual parameters are evaluated before starting new gouroutine. Thus it is needed to evaluate them explicitly before wg.Add(1) so the possibly panicking code would not leave increased counter.

// Use

//    param := f(x)
//    wg.Add(1)
//    go g(param)
//    instead of
// 
//    wg.Add(1)
//    go g(f(x))
