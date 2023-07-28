package main

import (
	"math"

	s "sort"
)

const bookPrice = 800

var discountTiers = [...]int{0, 5, 10, 20, 25}

// Cost computes the combination of books with the lowest total price

func CostComunity(books []int) int {

	organize(books)

	return cost(books, 0)

}

func organize(books []int) {

	type kv struct {
		Key int

		Value int
	}

	freq := make(map[int]int)

	for i := range books {

		freq[books[i]]++

	}

	ss := make([]kv, len(freq))

	for k, v := range freq {

		ss = append(ss, kv{k, v})

	}

	s.Slice(ss, func(i, j int) bool {

		return ss[i].Value > ss[j].Value

	})

	p := 0

	for _, kv := range ss {

		for i := 0; i < kv.Value; i++ {

			books[p] = kv.Key

			p++

		}

	}

}

func cost(books []int, priceSoFar int) int {

	if len(books) == 0 {

		return priceSoFar

	}

	distinctBooks, remainingBooks := getDistinctBooks(books)

	minPrice := math.MaxInt32

	for i := 1; i <= len(distinctBooks); i++ {

		newRemainingBooks := make([]int, len(remainingBooks))

		copy(newRemainingBooks, remainingBooks)

		newRemainingBooks = append(newRemainingBooks, distinctBooks[i:]...)

		price := cost(newRemainingBooks, priceSoFar+groupCost(i))

		if price < minPrice {

			minPrice = price

		}

	}

	return minPrice

}

func getDistinctBooks(books []int) (distinct []int, remaining []int) {

	exists := make(map[int]bool)

	for _, book := range books {

		if exists[book] {

			remaining = append(remaining, book)

		} else {

			distinct = append(distinct, book)

			exists[book] = true

		}

	}

	return

}

func groupCost(groupSize int) int {

	normalPrice := bookPrice * groupSize

	discount := (normalPrice * discountTiers[groupSize-1]) / 100

	return normalPrice - discount

}
