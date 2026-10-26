package main

import "fmt"

func partation(n []int, low int, high int) int {
	pivot := n[low]
	i := low
	j := high
	for i < j {
		for i <= high && n[i] <= pivot {
			i++
		}

		for j >= low && n[j] > pivot {
			j--
		}
		if i < j {
			n[i], n[j] = n[j], n[i]
		}
	}
	n[j], n[low] = n[low], n[j]
	return j
}
func quickSort(n []int, low int, high int) {
	if low < high {
		j := partation(n, low, high)
		quickSort(n, low, j-1)
		quickSort(n, j+1, high)
	}
}
func main() {
	numbers := []int{7, 2, 9, 4, 1, 6, 3, 8, 5}

	fmt.Println("Before sorting:")
	fmt.Println(numbers)

	quickSort(numbers, 0, len(numbers)-1)

	fmt.Println("\nAfter sorting:")
	fmt.Println(numbers)

}
