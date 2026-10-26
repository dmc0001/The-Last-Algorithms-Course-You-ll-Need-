package main

import "fmt"

func bubbleSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] < n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}

func main() {
	unsorted := []int{1, 7, 6, 3, 100, 2, 10}
	result := bubbleSort(unsorted)
	fmt.Println(result)
}
