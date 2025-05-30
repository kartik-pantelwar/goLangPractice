package main
import(
	"fmt"
)
func main(){
	//creating map
	//syntax-> mapName =  make(map[KeyType]ValueType)
	map1 := make(map[int]string)
	//adding values
	map1[36]= "Kartik";
	map1[20]= "Dhananjay";
	map1[46]= "Sharique";
	//printing all values
	fmt.Println(map1)
	//printing specific key value
	fmt.Println(map1[36])

	//Deleting values from Maps
	//*We can use this for slices as well
	//delete(mapName, Key)
	delete(map1,20)
	fmt.Println(map1)

	//traversing using loops
	//similar to foreach loop
	for key, value := range map1{ //we can write anything instead of 'key' and 'value'
		fmt.Printf("For Key %v, value is %v\n",key, value)
	}
}