package mathematics

import (
	"context"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type CalculatorSuite struct {
	suite.Suite
	calculator *Calculator
}

func (s *CalculatorSuite) SetupSuite() {
	s.calculator = New()
}

func (s *CalculatorSuite) TestFactorial() {
	ctx := context.TODO()
	s.Require().Equal(s.calculator.Factorial(ctx, 5), big.NewInt(120))
	s.Require().Equal(s.calculator.Factorial(ctx, 6), big.NewInt(720))

}

func TestCalculatorFactorial(t *testing.T) {
	suite.Run(t, new(CalculatorSuite))
}
