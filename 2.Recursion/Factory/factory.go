package main

import "fmt"

func main() {
	fmt.Println(factory(5))
}

func factory(n int) int {
	//Base case
	if n == 1 {
		return 1
	}
	//We shall recurse
	return factory(n-1) * n
}
