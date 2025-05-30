package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating for our pizza")

	input, _ := reader.ReadString('\n')
	var rating int = int(input)
	// input(_ means blank), error := reader.ReadString(ender-value)
	// it is like try catch. comma ok syntax in go
	fmt.Println("this  is input ",rating);
}