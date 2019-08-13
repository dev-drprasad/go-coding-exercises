package bignumber_test

import (
	"codingexercises/bignumber"
	"testing"
)

func TestBigNumber(t *testing.T) {
	max := bignumber.FindBigNumber([]uint{0, 1, 10, 4, 2, 13, 3, 16, 19})

	if max != 19 {
		t.Errorf("big number is %d, want %d", max, 19)
	}
}
