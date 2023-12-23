package main

import (
	"fmt"
	"time"
)

type A struct {
	Field1 string
	Field2 string
	Field3 int
}

type B struct {
	Field1 *string
	Field2 *string
	Field3 *int
}

func (b *B) From(a *A) {
	b.Field1 = &a.Field1
	b.Field2 = &a.Field2
	b.Field3 = &a.Field3
}

type BNP struct {
	Field1 string
	Field2 string
	Field3 int
}

func (b *BNP) From(a *A) {
	b.Field1 = a.Field1
	b.Field2 = a.Field2
	b.Field3 = a.Field3
}

func getA() A {
	return A{
		Field1: "ASasdasd",
		Field2: "SSasasdasd",
		Field3: 21,
	}
}

func main() {
	start := time.Now()

	for i := 0; i < 1_000_000; i++ {
		b := B{}
		a := getA()
		b.From(&a)
	}

	fmt.Printf("Time with pointers %s\n", time.Since(start).String())

	start = time.Now()

	for i := 0; i < 1_000_000; i++ {
		b := BNP{}
		a := getA()
		b.From(&a)
	}

	fmt.Printf("Time with no pointers %s\n", time.Since(start).String())
}
