package main

import "fmt"

func partition(A []int, p int, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j <= r; j++ {
		if A[j] < x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	i++
	A[i], A[r] = A[r], A[i]
	return i
}

func quicksort(A []int, p int, r int) {
	if p < r {
		q := partition(A, p, r)
		//哨兵不用排了
		quicksort(A, p, q-1)
		quicksort(A, q+1, r)
	}
}

func main() {
	A := [10]int{1, 3, 4, 5, 21, 23, 4, 1, 2, 3}
	fmt.Println("before")
	fmt.Println(A)
	quicksort(A[:], 0, len(A)-1)
	fmt.Println("after")
	fmt.Println(A)
}
