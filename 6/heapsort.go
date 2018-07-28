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

func heap_maximum(A []int) int {
	return A[0]
}

func heap_extract_max(A []int) int {
	if len(A) < 1 {
		panic("die")
	}
	max := A[0]
	A[0] = A[len(A)-1]
	max_heapify(A[:len(A)-1], 0)
	return max
}

func heap_increase_key(A []int, i int, key int) bool {
	if key < A[i] {
		return false
	}

	A[i] = key
	for i > 0 && A[parent(i)] < A[i] {
		A[i], A[parent(i)] = A[parent(i)], A[i]
		i = parent(i)
	}
	return true
}

func max_heap_insert(A []int, key int) bool {
	B := append(A, key)
	heap_increase_key(B, len(B)-1, key)
	return true
}

func main() {
	A := [10]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println("before")
	fmt.Println(A)
	heapsort(A[:])
	fmt.Println("after")
	fmt.Println(A)
	fmt.Println("build heap")
	build_max_heap(A[:])
	fmt.Println(A)
	fmt.Println(heap_maximum(A[:]))
	fmt.Println("increase 4 to 100")
	heap_increase_key(A[:], 4, 100)
	fmt.Println(heap_maximum(A[:]))

}
