package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func merge(A []int, p int, q int, r int) {

	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	L[n1] = MaxInt
	R[n2] = MaxInt

	i := 0
	j := 0
	for x := p; x <= r; x++ {
		if L[i] < R[j] {
			A[x] = L[i]
			i++
		} else {
			A[x] = R[j]
			j++
		}
	}
}

func merge_sort(A []int, p int, r int) {
	fmt.Println(A, p, r)
	if p < r {
		q := (p + r) / 2
		merge_sort(A, p, q)
		merge_sort(A, q+1, r)
		merge(A, p, q, r)
	}
}

func main() {
	A := [10]int{4, 5, 2, 3, 7, 2, 3, 8, 9, 0}
	fmt.Println("before")
	fmt.Println(A)
	merge_sort(A[:], 0, len(A)-1)
	fmt.Println("after")
	fmt.Println(A)
}
