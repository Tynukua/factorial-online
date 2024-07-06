package main

import (
	"math/big"
	"testing"

	"github.com/Tynukua/factorial-online/config"
	"github.com/Tynukua/factorial-online/services"
	"github.com/Tynukua/factorial-online/util"
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

type FactorialCase struct {
	a         int
	b         int
	expecteda *big.Int
	expectedb *big.Int
}

func TestDoubleFactorial(t *testing.T) {
	s := services.NewFactorialService(config.Config{})

	cases := []FactorialCase{
		{1, 2, big.NewInt(1), big.NewInt(2)},
		{1, 2, big.NewInt(1), big.NewInt(2)},
		{1, 5, big.NewInt(1), big.NewInt(120)},
		{5, 1, big.NewInt(120), big.NewInt(1)},
		{11, 11, big.NewInt(39916800), big.NewInt(39916800)},
	}

	for _, c := range cases {
		gota, gotb := s.DoubleFactorial(c.a, c.b)
		if gota.Cmp(c.expecteda) != 0 || gotb.Cmp(c.expectedb) != 0 {
			t.Fatalf("DoubleFactorial(%d, %d) = %d, %d, want %d, %d", c.a, c.b, gota, gotb, c.expecteda, c.expectedb)
		}
	}
}
