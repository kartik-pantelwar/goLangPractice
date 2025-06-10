package main

import (
	// "fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)
func Test_printSomething(t *testing.T){
	stdOut:= os.Stdout	//this type store the terminal output, it can directly be string, because terminal output can contain anything(string, error, or integer)
	r, w, _ := os.Pipe()
	stdOut = w
	
	var wg sync.WaitGroup
	wg.Add(1)
	printSomething(&wg, 1, "hello")
	wg.Wait()

	_ = w.Close()
	result,_:= io.ReadAll(r)
	output:= string(result)
	os.Stdout = stdOut

	if !strings.Contains(output,"hello"){
		t.Errorf("Expected to find hello, but got something else")
	}
}