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

func (suite *CalculatorTestSuite) SetupSuite() {
	suite.calculator = New()
}

func (suite *CalculatorTestSuite) TestFactorial() {
	ctx := context.Background()

	suite.Require().Equal(suite.calculator.Factorial(ctx, 0), big.NewInt(1))
	suite.Require().Equal(suite.calculator.Factorial(ctx, 1), big.NewInt(1))
	suite.Require().Equal(suite.calculator.Factorial(ctx, 10), big.NewInt(3628800))

}

func TestCalculatorFactorial(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
