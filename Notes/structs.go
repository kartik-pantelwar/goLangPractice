//there is no concept of inheritance

package main

import (
	"fmt"
)

func main(){
	//syntax-> type structName struct{parameter DataType}
	//*Structname begin with capital letter, because we want to keep it public, so that we can export it

	type Car struct{
		Name string
		Price int
		IsAvailable bool
	}

	//*All the fields (Name, Price, IsAvailable) are also Capital(public) because they needs to be accessed by anyone 

	//creating object(car)
	car1 := Car{"Audi",500,true}
	car2 := Car{Name: "Bugatti", IsAvailable: true, Price: 700}
	fmt.Println(car1)
	fmt.Println(car2)

}