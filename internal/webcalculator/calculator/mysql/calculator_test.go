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

type CalculatorSuite struct {
	suite.Suite
	calculator *Calculator
}

func (s *CalculatorSuite) SetupSuite() {

	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/testdb")
	s.Require().NoError(err)
	s.calculator = New(db, mathematics.New())
}

func (s *CalculatorSuite) TestFactorial() {
	ctx := context.TODO()
	s.Require().Equal(s.calculator.Factorial(ctx, 5), big.NewInt(120))
	s.Require().Equal(s.calculator.Factorial(ctx, 6), big.NewInt(720))

}

func TestNewCalculator(t *testing.T) {
	suite.Run(t, new(CalculatorSuite))

}
