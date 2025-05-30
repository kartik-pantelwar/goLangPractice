package main
import(
	"fmt"
)

func main(){
	//we can allocated memory using new() and make()
	//make() is most commonly used, because in it, memory is allocated and initialized also.
	//new() is not commonly used, becasuse in it, mmemory  is allocated, but not initialized.
	//garbage collection is automatic in GoLang

	a:=10
	var ptr *int 
	ptr = &a
	fmt.Println("Value store in pointer-",*ptr)	//print the value kept at the memory, at which the pointer is pointing

	fmt.Println("Memory block at which pointer is pointing-",ptr)	//printing the memory block at which pointer is pointing

	fmt.Println("Address where the pointer is stored in memory-",&ptr)	//printing the address where the pointer is stored
		
}