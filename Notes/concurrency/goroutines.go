// package main

// import (
// 	"fmt"
// 	"time"
// )
// func name(n string){
// 	for i:=0;i<5;i++ {
// 		time.Sleep(time.Second);
// 		fmt.Println(n)
// 	}
// }
// func main(){
// 	go name("kartik");
// 	name("arora")
// }



package main

import (
    "fmt"
    "time"
)

// Function to print numbers
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(500 * time.Millisecond) // Simulate work
        fmt.Println("Number:", i)
    }
}

// Function to print letters
func printLetters() {
    for _, letter := range []string{"A", "B", "C", "D", "E"} {
        time.Sleep(500 * time.Millisecond) // Simulate work
        fmt.Println("Letter:", letter)
    }
}

func main() {
    // Start both functions as goroutines
    go printNumbers()
    go printLetters()

    // Give goroutines time to complete
    time.Sleep(3 * time.Second)
    fmt.Println("Main function done!")
}
