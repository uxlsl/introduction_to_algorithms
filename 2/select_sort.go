package main

import "fmt"

func select_sort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func main() {
	fmt.Println("befort!")
	a := [5]int{3, 2, 1, 4, 0}
	fmt.Println(a)
	fmt.Println("after!")
	select_sort(a[:])
	fmt.Println(a)
}
