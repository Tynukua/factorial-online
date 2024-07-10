package mathematics

import (
	"context"
	"github.com/stretchr/testify/suite"
	"math/big"
)

type CalclualtorSuite struct {
	suite.Suite
	mc *Calculator
}

func (s *CalclualtorSuite) SetupSuite() {
	s.mc = NewCalculator()
}

func (s *CalclualtorSuite) TestFactorial() {
	ctx := context.TODO()
	s.Require().Equal(s.mc.Factorial(ctx, 5), big.NewInt(120))
	s.Require().Equal(s.mc.Factorial(ctx, 6), big.NewInt(720))

}
