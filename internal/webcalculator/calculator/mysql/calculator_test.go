package mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online/internal/webcalculator/calculator/mathematics"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type CalculatorTestSuite struct {
	suite.Suite
	calculator *Calculator
}

func (suite *CalculatorTestSuite) SetupSuite() {

	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/testdb")
	suite.Require().NoError(err)
	suite.calculator = New(db, mathematics.New())
}

func (suite *CalculatorTestSuite) TestFactorial() {
	ctx := context.Background()
	suite.Require().Equal(suite.calculator.Factorial(ctx, 0), big.NewInt(1))
	suite.Require().Equal(suite.calculator.Factorial(ctx, 1), big.NewInt(1))
	suite.Require().Equal(suite.calculator.Factorial(ctx, 10), big.NewInt(3628800))

}

func TestNewCalculator(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))

}
