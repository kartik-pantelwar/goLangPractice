package main

import (
	"fmt"
	khudkaPackage "package2/exportable" //* only required package will be imported
	//* "package2/exportable" (all packages are imported)
)
func main(){
	fmt.Println(khudkaPackage.Code)
}