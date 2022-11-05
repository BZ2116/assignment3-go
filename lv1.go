package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

func add() {
	ch := make(chan int, 1)
	for i := 0; i <= 100000; i++ {
		ch <- i
		x = <-ch
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
