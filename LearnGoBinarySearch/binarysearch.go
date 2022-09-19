package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 17, 18, 19, 20, 22, 23, 24, 25, 25, 26, 28, 29, 30}
	find := binarySearch(array, 24, 0, len(array)-1)
	fmt.Printf("Found element at position %d using recursion\n", find)
	find = iterBinarySearch(array, 24, 0, len(array)-1)
	fmt.Printf("Found element at position %d using iteration\n", find)
}

// Recursive binary search function
func binarySearch(array []int, target int, lowIdx int, highIdx int) int {
	if highIdx < lowIdx {
		return -1
	}

	mid := int((lowIdx + highIdx) / 2)

	if array[mid] > target {
		return binarySearch(array, target, lowIdx, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIdx)
	} else {
		return mid
	}
}

// Binary Search function without using recursion (using iteration)
func iterBinarySearch(array []int, target int, lowIdx int, highIdx int) int {
	startIdx := lowIdx
	endIdx := highIdx
	var mid int

	for startIdx < endIdx {
		mid = int((startIdx + endIdx) / 2)
		if array[mid] > target {
			endIdx = mid
		} else if array[mid] < target {
			startIdx = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
