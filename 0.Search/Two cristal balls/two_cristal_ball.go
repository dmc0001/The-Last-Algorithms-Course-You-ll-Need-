package main

import (
	"fmt"
	"math"
)

func main() {
	breaks := []bool{false, false, false, false, true, true, true}
	result := twoCritalBalls(breaks)
	fmt.Println(result)

}

func twoCritalBalls(n []bool) int {
	jumpAmout := int(math.Sqrt(float64(len(n))))

	i := jumpAmout

	for i < len(n) {
		if n[i] == true {
			break
		}
		i += jumpAmout
	}

	j := i - jumpAmout

	for j < len(n) {
		if n[j] == true {
			return j
		}
		j++
	}

	return -1

}
