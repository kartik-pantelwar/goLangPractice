package main
import (
	"fmt"
)
type Animal struct {
	Name string
	Colour string
	Price int
}
func (a Animal) EMI()int{
	return a.Price/10
}

func main(){
	dog := Animal{"Lebra", "Cream", 10000}
	fmt.Println("EMI of Dog is-",dog.EMI())
}