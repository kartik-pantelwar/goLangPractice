package main 
import (
	"fmt"
)
type Person struct {
	Name string
	Age int
}

//'*Person' is Pointer receiver

//It stores the memory address where a Person struct is created *i.e address of p1

//It is important, is it changes the actual value of that person. If we will only 'Person' instead of '*Person', then only the copy would be created, and it would be changed, not the actual value would be changed.

//Method with Pointer Receiver
func (p *Person) AgeBarao()int{ 
	p.Age++
	return p.Age
}	

//return is optional, as the actual value is updated
// func (p *Person) AgeBarao(){ 
// 	p.Age++
// }

//Method without Pointer Receiver
func (p Person) AgeGhatao()int{
	p.Age--
	return p.Age
}

func main(){

	//using pointer Receiver
	p1 := Person{"Kartik", 21}
	fmt.Println("Before-",p1)
	p1.AgeBarao()	//AgeBarao method used
	fmt.Println("After-",p1)//value updated

	//without pointer receiver
	p2 := Person{"Dhananjay",22}
	fmt.Println("Before-",p2)
	p2.AgeGhatao()	//AgeGhatao method used
	fmt.Println("After-",p2)//value not updated

}