package main

import "fmt"

func forLoopConcurrent() {
	limiter := make(chan int64, 20)
	for i := 0; i < 1000000; i++ {
		limiter <- 1
		go func(t int) {
			doWork(t)
			<-limiter
		}(i)
	}
}

func forLoop() {
	for i := 0; i < 1000000; i++ {
		doWork(i)
	}
}

func doWork(i int) {
	_ = fmt.Sprint(i)
}
