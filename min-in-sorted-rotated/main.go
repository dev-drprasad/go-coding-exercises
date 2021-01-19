package main

import "fmt"

// Find the minimum element in a sorted and rotated array
// https://www.geeksforgeeks.org/find-minimum-element-in-a-sorted-and-rotated-array/

/**
A simple solution is to traverse the complete array and find a minimum. This solution requires O(n) time.
We can do it in O(Logn) using Binary Search. If we take a closer look at the above examples, we can easily figure out the following pattern:

	- The minimum element is the only element whose previous is greater than it. If there is no previous element, then there is no rotation (the first element is minimum). We check this condition for the middle element by comparing it with (mid-1)’th and (mid+1)’th elements.
	- If the minimum element is not at the middle (neither mid nor mid + 1), then the minimum element lies in either the left half or right half.
	- If the middle element is smaller than the last element, then the minimum element lies in the left half,
		Else minimum element lies in the right half.
*/

func findMin(arr []int, leftmostIndex int, rightmostIndex int) int {

	// [1,2] , rightmostIndex will be -1 in 2nd depth
	if rightmostIndex < leftmostIndex {
		return arr[0]
	}

	// there is only one element left
	if leftmostIndex == rightmostIndex {
		return arr[leftmostIndex]
	}

	midIndex := (leftmostIndex + rightmostIndex) / 2

	// [1]
	if midIndex < rightmostIndex {

		//  cases like [3, 4, 5, 1, 2]
		if arr[midIndex+1] < arr[midIndex] {
			return arr[midIndex+1]
		}
	}

	// [1,2]
	if leftmostIndex < midIndex {
		// [4, 5, 1, 2, 3]
		if arr[midIndex] < arr[midIndex-1] {
			return arr[midIndex]
		}
	}

	if arr[rightmostIndex] > arr[midIndex] {
		return findMin(arr, leftmostIndex, midIndex-1)
	}

	return findMin(arr, midIndex+1, rightmostIndex)
}

func main() {
	input := []int{5, 6, 1, 2, 3, 4}

	min := findMin(input, 0, len(input)-1)
	fmt.Println(min)

	input2 := []int{5, 6, 7, 1, 2, 3, 4}
	fmt.Println(findMin(input2, 0, len(input2)-1))

	input3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(findMin(input3, 0, len(input3)-1))

	input4 := []int{1, 2}
	fmt.Println(findMin(input4, 0, len(input4)-1))
}
