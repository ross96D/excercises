package main

import (
	"fmt"
	"sync"
	"time"
)

func Syncgroup() {
	fmt.Println("Holas")

	w := async()
	w.Wait()
	fmt.Println("COMIDA")
	time.Sleep(2 * time.Second)
}

func async() *sync.WaitGroup {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		fmt.Println("MANDAR LA GOROUTINE")
		w.Done()
		fmt.Println("TERMINO LA GOROUTINE")
	}()

	return &w
}
