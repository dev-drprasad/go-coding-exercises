package bignumberrecursive_test

import (
	bignumber "codingexercises/bignumber-recursive"
	"testing"
)

func TestBigNumber(t *testing.T) {
	max := bignumber.FindBigNumber([]uint{0, 1, 10, 4, 2, 13, 3, 16, 19}, 9-1, 0)

	if max != 19 {
		t.Errorf("big number is %d, want %d", max, 19)
	}
}
