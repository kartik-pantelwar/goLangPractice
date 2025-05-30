//*Benefits of Interfaces*
//polymorphism
//Data Abstraction
//readble code 
//understandable
//testable codel (while using database commands)

//*Disadvantages*
//complexity in Design
//Limited to Method Signatures: Interfaces in Go are limited to method signatures, they can not include fields, which might be necessary in some scenarios.
//no genrics before 18

package main
import (
	"fmt"
)

type Car struct{
	price int;
	name string;
}

type Bike struct{
	price int;
	name string;
}

type Auto struct{
	price int;
	name string;
}

type Rate interface{
	tellPrice() int
}

func(c Car) tellPrice()int {
	return c.price;
}

func(b Bike) tellPrice()int {
	return b.price;
}

func(a Auto) tellPrice()int{
	return a.price;
}

func main(){
	audi := Car{price: 400, name: "audi"}
	ktm := Bike{price: 700, name: "ktm"}
	bajaj := Auto{price: 1000,name: "bajaj"}

	prices := []Rate{audi,ktm,bajaj}
	for _,value := range prices{
		fmt.Println("price of",value.tellPrice())
	}
	// fmt.Println(audi.tellPrice());
	// fmt.Println(ktm.tellPrice())
	// fmt.Println(bajaj.tellPrice())
}