package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 5, 8, 3, 55, 88}
	var i int = size/2 - 1

	heapify(a, len(a), i)
	//Insert(a, 1000)
	printheap(a, len(a))
}
func Swap(b *int, c *int) {
	temp := *b
	*b = *c
	*c = temp
}

var size int = 0

func heapify(a []int, size int, i int) {
	for i := 0; i < size; i++ {

		if size == 1 {
			fmt.Printf("only one element in heap...")

		} else {
			left := 2*i + 1
			right := 2*i + 2
			largest := i

			if left < size && a[left] > a[largest] {
				largest = left

			}
			if right < size && a[right] > a[largest] {
				largest = right

			}
			if largest != i {
				Swap(&a[i], &a[largest])
				heapify(a, size, largest)

			}
		}
	}
}

func printheap(a []int, size int) {

	for i := 0; i < size; i++ {
		fmt.Printf("%d->", a[i])
	}

}

func Insert(a []int, n int) {
	i := a[0]
	if size == 0 {
		a[0] = n
		size++
	} else {
		a[size] = n
		size++
		heapify(a, size, i)
	}
}
