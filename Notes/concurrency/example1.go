package main
import (
    "fmt"
    "time"
)
func task(id int) {
    time.Sleep(time.Second) // Simulates work
    fmt.Println("Task", id, "finished")
}
func main() {
	//* Using GoRoutines, Tasks run simultaneously
	//^Faster Execution
    go task(1) 
    go task(2)
    go task(3)
    go task(4)
    go task(5)
    go task(6)
    go task(7)
	time.Sleep(time.Second*3)
    fmt.Println("All tasks finished!")

	//* Without Goroutine, Task runs one by one(one after other)
	//! Slower Execution
	task(1) 
    task(2)
    task(3)
    task(4)
    task(5)
    task(6)
    task(7)
    fmt.Println("All tasks finished!")
}