package main

import (
	"testing"
)

type FactorialCase struct {
	a         int
	b         int
	expecteda int
	expectedb int
}

func TestFactorial(t *testing.T) {
	cases := []FactorialCase{
		{1, 2, 1, 2},
		{2, 1, 2, 1},

		{2, 3, 2, 6},
		{3, 2, 6, 2},

		{5, 5, 120, 120},

		{5, 11, 120, 39916800},
		{11, 5, 39916800, 120},

		{10, 11, 39916800 / 11, 39916800},
	}

	for _, c := range cases {
		gota, gotb := doubleFactorial(c.a, c.b)
		if gota != c.expecteda || gotb != c.expectedb {
			t.Fatalf("doubleFactorial(%d, %d) = (%d, %d), want (%d, %d)", c.a, c.b, gota, gotb, c.expecteda, c.expectedb)
		}
	}

}
