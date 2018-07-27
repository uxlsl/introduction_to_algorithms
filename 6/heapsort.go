package main

import "fmt"

func parent(i int) int {
	return i / 2 // 刚好取整
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func max_heapify(A []int, i int) {
	l := left(i)
	r := right(i)
	largest := i
	if l < len(A) && A[l] > A[i] {
		largest = l
	}

	if r < len(A) && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		max_heapify(A, largest)
	}

}

func build_max_heap(A []int) {
	for i := len(A) / 2; i >= 0; i-- {
		max_heapify(A, i)
	}
}

func heapsort(A []int) {
	build_max_heap(A)
	for i := len(A) - 1; i > 0; i-- {
		A[i], A[0] = A[0], A[i]
		max_heapify(A[:i], 0)
	}
}

func main() {
	A := [10]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println("before")
	fmt.Println(A)
	heapsort(A[:])
	fmt.Println("after")
	fmt.Println(A)
}
