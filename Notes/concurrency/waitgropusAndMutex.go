// package main
// import (
// 	"fmt"
// 	// "sync"
// 	"time"
// )
// func printLetters(){
// 	for i := 'a'; i < 'e'; i++ {
// 		fmt.Println(string(i))
// 		time.Sleep(time.Millisecond*500)
// 	}
// }
// func printNumbers(){
// 	for i:=0;i<5;i++{
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }
// func main(){
// 	go printLetters()
// 	// go printNumbers()
// 	time.Sleep(time.Second*3)
// 	fmt.Println("Main Execution Finished")
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	time.Sleep(time.Second) // Simulates work
	fmt.Println("Task", id, "finished")
}
func main() {
	var wg sync.WaitGroup //waitgroup
	var mut sync.Mutex    //Mutex
	wg.Add(2)
	mut.Lock()
	go task(1, &wg)
	mut.Unlock()
	mut.Lock()
	go task(2, &wg)
	mut.Unlock()
	// time.Sleep(time.Second*3)
	wg.Wait()
	fmt.Println("All tasks finished!")
}
