package main

import (
	"fmt"
)

func main() {
	arr1 := []int{0, 1, 2, 3, 4, 6, 8, 10, 11, 16}
	arr2 := []int{1, 3, 5, 7, 8, 12, 14, 16, 17, 18}

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			fmt.Println(arr1[i])
			i++

			j++

		}

	}
}
