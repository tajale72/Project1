package main

import (
	"fmt"
)

func main() {
	var sum int
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range arr {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Printf("%s", string(sum))
}
