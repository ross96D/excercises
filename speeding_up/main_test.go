package main

import "testing"

func TestMain(t *testing.T) {

}

func BenchmarkCon(b *testing.B) {
	forLoopConcurrent()
}

func BenchmarkSync(b *testing.B) {
	forLoop()
}
