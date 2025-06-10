// package main

// import (
// 	"fmt"
// )

// type Income struct {
// 	Source string
// 	Amount int
// }

// func main() {

// 	//defining intial balance
// 	var bankBalance int

// 	//printing initial balance
// 	fmt.Println("Initial Balance-", bankBalance)

// 	//weekly income
// 	incomes := []Income{
// 		{"Job", 100},
// 		{"Stocks", 30},
// 		{"Crypto", 40},
// 		{"YouTube", 50},
// 	}

// 	//calculation yearly balance
// 	for _, i := range incomes {
// 		for j := 0; j < 52; j++ {
// 			bankBalance = bankBalance + i.Amount
// 		}
// 	}

// 	//printing yearly balance
// 	fmt.Println("Final Balance-", bankBalance)
// }

package main

import (
	"fmt"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	name := "kartik"

	wg.Add(1)
	go func(){fmt.Printf("%s 1\n",name)
	defer wg.Done()}()
	wg.Wait()

	wg.Add(1)
	go func(){fmt.Printf("%s 2\n",name)
	defer wg.Done()}()
	wg.Wait()

	wg.Add(1)
	go func(){fmt.Printf("%s 3\n",name)
	defer wg.Done()}()
	wg.Wait()
}