package bignumberrecursive

func FindBigNumber(arr []uint, n uint, max uint) uint {
	if n == 0 {
		return max
	}

	if arr[n] > max {
		max = arr[n]
	}

	return FindBigNumber(arr, n-1, max)
}
