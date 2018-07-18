package main

import "fmt"

func insertion_sort(a []int) {
	for j := 1; j < len(a); j++ {
		v := a[j]
		i := j - 1
		for ; 0 <= i && a[i] > v; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = v
	}
}

func main() {
	fmt.Println("befort!")
	a := [5]int{3, 2, 1, 4, 0}
	fmt.Println(a)
	fmt.Println("after!")
	insertion_sort(a[:])
	fmt.Println(a)
}
