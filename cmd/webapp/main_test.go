package main

import (
	"github.com/Tynukua/factorial-online/internal/util"
	"math/big"
	"testing"
)

type MulRangeCase struct {
	a        int
	b        int
	expected *big.Int
}

func TestMulRange(t *testing.T) {

	cases := []MulRangeCase{
		{1, 2, big.NewInt(2)},
		{1, 5, big.NewInt(120)},
		{3, 5, big.NewInt(60)},
	}

	for _, c := range cases {
		got := util.MulRange(c.a, c.b)
		if got.Cmp(c.expected) != 0 {
			t.Fatalf("MulRange(%d, %d) = %d, want %d", c.a, c.b, got, c.expected)
		}
	}

}
func TestMulRangeParrallel(t *testing.T) {
	cases := []MulRangeCase{
		{1, 2, big.NewInt(2)},
		{1, 5, big.NewInt(120)},
		{3, 5, big.NewInt(60)},
	}

	for _, c := range cases {
		got := util.MulRangeParallel(c.a, c.b, 2)
		if got.Cmp(c.expected) != 0 {
			t.Fatalf("MulRange(%d, %d) = %d, want %d", c.a, c.b, got, c.expected)
		}
	}
}
