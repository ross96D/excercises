package main

import "testing"

func TestCost(t *testing.T) {
	a := Cost([]int{0, 0, 1, 1, 2, 2, 3, 4})
	println(a)
}

func BenchmarkCost(b *testing.B) {
	Cost([]int{0, 1, 2, 0, 1, 4, 3, 1, 2, 0, 1, 3, 4})
}
