package main

import "fmt"

func bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func main() {
	fmt.Println("befort!")
	a := [5]int{3, 2, 1, 4, 0}
	fmt.Println(a)
	fmt.Println("after!")
	bubblesort(a[:])
	fmt.Println(a)
}
