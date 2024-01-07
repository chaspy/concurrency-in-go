package main

import "sync"

func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
