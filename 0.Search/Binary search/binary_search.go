package main

import "fmt"

func main() {
	nums := []int{1, 3, 67, 232, 240, 323, 391}
	result := binarySearch(nums, 391)
	fmt.Println(result)
}

func binarySearch(n []int, element int) bool {
	low := 0
	high := len(n)
	mid := 0

	for low <= high {
		mid = low + (high-low)/2
		if n[mid] == element {
			return true
		} else if n[mid] > element {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}
