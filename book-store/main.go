package main

// Link to the excercise https://exercism.org/tracks/go/exercises/book-store
func main() {
	a := Cost([]int{0, 0, 1, 1, 2, 2, 3, 4})
	println(a)
}

func Cost(books []int) int {
	var entered [5]int
	for i := 0; i < len(books); i++ {
		entered[books[i]] = entered[books[i]] + 1
	}

	ordered := sort(entered[0:])

	m := ordered[2] - ordered[3]
	if m < ordered[4] {
		ordered[4] -= m
		ordered[3] += m
	} else {
		ordered[3] += ordered[4]
		ordered[4] = 0
	}

	total := ordered[4] * 600 * 5
	total += (ordered[3] - ordered[4]) * 640 * 4
	total += (ordered[2] - ordered[3]) * 640 * 3
	total += (ordered[1] - ordered[2]) * 640 * 2
	total += (ordered[0] - ordered[1]) * 640

	return total
}

func sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] > a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sort(a[:left])
	sort(a[left+1:])

	return a
}
