package main

import "fmt"

func main() {
	nums := []int{1, 3, 67, 232, 0, 23, 1}
	result := linearSearch(nums, 3)
	fmt.Println(result)
}

func linearSearch(n []int, element int) bool {
	for i := 0; i < len(n); i++ {
		if n[i] == element {
			return true
		}
	}
	return false
}
