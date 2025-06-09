package main

import (
	"fmt"
	"net/http"
	// "io"
	// "os"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to Listen and Server", err)
	}
}
func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server Running!")
	w.Write([]byte("Hello World"))
}
