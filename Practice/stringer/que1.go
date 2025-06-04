//Stringer exercise, A Tour of Go

package main
import (
	"fmt"
)
type IPAdd [4]byte;
func main(){
	ip1 := IPAdd{10,20,30,40};
	fmt.Printf("%v.%v.%v.%v\n",ip1[0],ip1[1],ip1[2],ip1[3])
}