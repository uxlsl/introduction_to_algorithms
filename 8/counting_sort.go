package main

import "fmt"

func counting_sort(A []int, B []int, k int) {
	C := make([]int, k)
	for i := 0; i < len(C); i++ {
		C[i] = 0
	}
	for j := 0; j < len(A); j++ {
		C[A[j]] += 1
	}
	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1]
	}
	for j := len(A) - 1; j >= 0; j-- {
		B[C[A[j]]-1] = A[j]
		C[A[j]] -= 1
	}
}

func main() {
	fmt.Println("starting")
	A := [11]int{5, 3, 2, 1, 2, 1, 2, 4, 6, 7, 7}
	B := make([]int, 11)
	fmt.Println(A)
	counting_sort(A[:], B[:], 8)
	fmt.Println(B)
}
