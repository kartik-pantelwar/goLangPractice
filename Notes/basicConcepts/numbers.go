package main 
import (
	"fmt"
)
// %b = binary, o = octal, d = decimal, X = hexadecimal 
// # = address 
// x(small x) = for small hexadecimal letters(a-f)
// %q = UTF-8 character value(ASCII)
func main(){
	fmt.Printf("%b - %d - %o - %#X", 42, 42, 42, 42)
}