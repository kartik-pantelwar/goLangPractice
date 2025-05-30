//Modifying a Value via Pointer in a Function

//Doubling a Number Using Pointers

package main

import "fmt"

// Function that modifies the original variable using a pointer
func Double(num *int) {
    *num = *num * 2 // Dereferencing to modify the actual value
}

func main() {
    x := 5
    fmt.Println("Before:", x)

    Double(&x) // Pass the address of x

    fmt.Println("After:", x) // x is modified inside the function
}
