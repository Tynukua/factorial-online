package webservices

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/config"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type FactorialWebSuiteTest struct {
	suite.Suite
	s FactorialService
}
type doubleFactorialCase struct {
	a, b                 int
	expectedA, expectedB *big.Int
}

func (suite *FactorialWebSuiteTest) SetupSuite() {
	suite.s = NewFactorialService(config.Config{DBType: config.Memory})
}

func TestFactorialService(t *testing.T) {
	suite.Run(t, new(FactorialWebSuiteTest))
}

func (suite *FactorialWebSuiteTest) TestDoubleFactorial() {
	ctx := context.Background()
	cases := []doubleFactorialCase{
		{1, 1, big.NewInt(1), big.NewInt(1)},
		{5, 5, big.NewInt(120), big.NewInt(120)},
		{10, 11, big.NewInt(3628800), big.NewInt(39916800)},
	}
	for _, c := range cases {
		gotA, gotB := suite.s.DoubleFactorial(ctx, c.a, c.b)
		suite.Require().Equal(c.expectedA, gotA)
		suite.Require().Equal(c.expectedB, gotB)
	}
}
