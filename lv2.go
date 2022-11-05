package main

import (
	"fmt"
	"sync"
)

var ch = make(chan struct{})
var wg sync.WaitGroup

func ji() {
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
		ch <- struct{}{}
		<-ch
		wg.Done()
	}
}
func ou() {
	for i := 2; i <= 100; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- struct{}{}
		wg.Done()
	}
}
func main() {
	wg.Add(100)
	go ji()
	go ou()
	wg.Wait()
}
