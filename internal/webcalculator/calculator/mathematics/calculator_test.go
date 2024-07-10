package mathematics

import (
	"context"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type CalculatorTestSuite struct {
	suite.Suite
	calculator *Calculator
}

func (s *CalculatorTestSuite) SetupSuite() {
	s.calculator = New()
}

func (s *CalculatorTestSuite) TestFactorial() {
	ctx := context.TODO()
	s.Require().Equal(s.calculator.Factorial(ctx, 5), big.NewInt(120))
	s.Require().Equal(s.calculator.Factorial(ctx, 6), big.NewInt(720))

}

func TestCalculatorFactorial(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
