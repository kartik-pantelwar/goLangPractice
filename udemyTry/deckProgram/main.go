package main 

import(
	"fmt"
)
func main(){
	merePatte := NewDeck()
	merePatte.print()
	
	//after shuffling
	fmt.Println("After Shuffling")
	shuffleDeck(merePatte)
	merePatte.print()
}