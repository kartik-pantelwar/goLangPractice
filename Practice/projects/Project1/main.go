package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	StatusCode int
	Status     string
	Err        error
}

func webScrapper(url string, results chan<- Result, w *sync.WaitGroup) {
	defer w.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	output := Result{res.StatusCode, res.Status, err}
	defer res.Body.Close()
	results <- output
}

func main() {
	//& creating a waitgroup
	urls := []string{"http://www.google.com", "http://www.youtube.com", "http://www.yahoo.com", "http://www.linkedin.com", "http://www.pantelwar.com"}
	var wg sync.WaitGroup

	//& Adding number of go routines to waitgroup
	wg.Add(5)
	ch := make(chan Result, 5)
	for _, i := range urls {
		go webScrapper(i, ch, &wg)
	}

	// wg.Wait()	//& Waitgroup waiting for goroutines to execute
	// close(ch)

	// for i:= range ch {
	// 	fmt.Printf("StatusCode= %d, Status= %v, Error= %e\n",i.StatusCode,i.Status,i.Err)
	// }

	//* close not required, as we are taking value out of the channel using operator(<-), instead of iterating the channel, i.e in previous case, so channel is automatically getting empty.
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

	//* Individual Prints, doesn't care if channel is open or closed
	//! doesn't even wait for more values, you can try removing any one statement, out of these 5
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	wg.Wait()

}
