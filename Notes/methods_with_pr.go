package main

import "fmt"

type Number struct {
	Value float64
}

// Method with a pointer receiver (*type)
// func (n *Number) Scale(factor float64) {
//     n.Value *= factor // Modifies the original value
// }

// Method with value receiver (type)
func (n *Number) Scale(factor float64) {
	n.Value = n.Value * factor
}

func main() {
	num := Number{Value: 5}
	fmt.Println("Before Value-", num.Value)

	num.Scale(2) // Go automatically converts `num.Scale(2)` to `(&num).Scale(2)`

	fmt.Println("After Value-", num.Value) // Output: 10
}
