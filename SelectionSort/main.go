package main

import (
	"fmt"
)

func main() {
	a := []int{2, 5, 11, 30, 4}
	SelectionSort(a)

}

func SelectionSort(a []int) {
	var i, j, min int
	for i = 0; i < len(a); i++ {
		min = i
		for j = i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
				Swap(&a[i], &a[min])

			}
		}

	}
	for i := 0; i < len(a); i++ {
		fmt.Printf(" %d-> ", a[i])
	}
}

func Swap(b *int, c *int) {
	temp := *b
	*b = *c
	*c = temp
}
