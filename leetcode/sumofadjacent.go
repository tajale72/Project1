package main

import (
	"fmt"
)

func Max(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}

	sumArray := make([]int, n)

	sumArray[0] = arr[0]
	sumArray[1] = Greater(arr[0], arr[1])

	for i := 2; i < n; i++ {
		sumArray[i] = Greater(sumArray[i-1], sumArray[i-2]+arr[i])
	}

	return sumArray[n-1]
}

func Greater(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{3, 7, 4, 6, 5}
	fmt.Println(Max(arr))
	fmt.Println(QuickSort(arr))

}

func QuickSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	pi := arr[0]

	var left, right []int

	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v < pi {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pi), right...)
}
