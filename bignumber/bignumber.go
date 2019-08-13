package bignumber

func FindBigNumber(arr []uint) uint {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}
