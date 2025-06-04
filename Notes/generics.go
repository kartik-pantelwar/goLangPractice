package main

import (
	"fmt"
)

// Generic Functions->
func Sum[T int | float64](a, b T) T {
	var summ T
	summ = a + b
	return summ
} //'any' can not be used here, because we are performing addition operation, so only those data types which support addition, can be assigned
func printt[P any, T any](First T, Second P) {
	fmt.Println(First, Second)
}

// Generic Types->
type Age[T any] struct {
	age T
}

func main() {
	fmt.Println("sum of Integers-", Sum(10, 20))
	fmt.Println("sum of Floats-", Sum(10.5, 20.1))

	printt(10, 20)
	printt("Kartik", "Arora")
	printt("Kartik", 10)

	ApproxAge := Age[int]{21}
	DetailedAge := Age[float64]{22};
	StringAge := Age[string]{"23"}
	fmt.Println(ApproxAge)
	fmt.Println(DetailedAge)
	fmt.Println(StringAge)
}
