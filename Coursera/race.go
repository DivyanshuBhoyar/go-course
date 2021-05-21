package main

import (
    "fmt"
    "sync"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {

    for routine := 1; routine <= 2; routine++ {

        Wait.Add(1)
        go Routine(routine)
    }

    Wait.Wait()
    fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

    for count := 0; count < 2; count++ {

        value := Counter
        value++
        Counter = value
    }

    Wait.Done()
}

// A race condition is when two or more routines have access to the same resource, such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines.
// run the code through the Go race detector and see what it finds. 
// Open a Terminal session where the source code is located and build the code using the -race option.