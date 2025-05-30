package main

import (
	"fmt"
	// "strconv"
)

// working
func binary_check(n int) string {
	i := n
	k := 0
	for i > 0 {
		if i%10 != 0 && i%10 != 1 {
			k = 1
			return "Non Binary"
		}
		i = i / 10
	}
	if k == 0 {
		return "Binary"
	} else {
		return "error"
	}
}

// func vowel()

// perfect
func n_sum() {
	var a int
	fmt.Print("Enter value of a-")
	fmt.Scan(&a)
	for i := 1; i < a+1; i++ {
		fmt.Print(i, " ")
	}
}
func odd_even(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
func swap_nums(a int, b int) (int, int) {
	return b, a
}

type Car struct{
	Name string;
}

func carname(c Car){
	fmt.Println(Car.Name)
}
func main() {
	car1 := Car{"Audi"}
	// fmt.Println(car1.name)
	carname(car1)
}
