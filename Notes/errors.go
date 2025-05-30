package main
import (
	"fmt"
	"errors"
)

func main(){
	err := errors.New("Display Error")
	fmt.Println(err)	
}