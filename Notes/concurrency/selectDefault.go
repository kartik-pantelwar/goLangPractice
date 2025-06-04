package main

import (
    "fmt"
    "time"
)

func main() {
    order1 := make(chan string)
    order2 := make(chan string)

    go func() {
        time.Sleep(3 * time.Second) // Simulate delay for order1
        order1 <- "Burger Ready!"
    }()

    go func() {
        time.Sleep(1 * time.Second) // Simulate delay for order2
        order2 <- "Pizza Ready!"
    }()

    // `select` listens to both channels and picks whichever is ready first
    select {
    case msg := <-order1:
        fmt.Println(msg) // Prints "Burger Ready!" if it's received first
    case msg := <-order2:
        fmt.Println(msg) // Prints "Pizza Ready!" if it's received first
	default:
        fmt.Println("Waiting for orders...") // Executes if both channels are blocked
    }
}
