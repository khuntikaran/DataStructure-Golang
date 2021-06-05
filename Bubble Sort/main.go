package main

import (
	"fmt"
)

func main() {
	a := []int{55, 32, 12, 46, 99, 2, 30, 21, 11, 19, 15}
	BubbleSort(a)
}

func BubbleSort(a []int) {
	var i, j int
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				Swap(&a[j], &a[j+1])
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
