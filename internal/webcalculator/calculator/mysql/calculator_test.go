package calculatormysql

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
	suite.calculator = New(db, calculatormathematics.New())
}

type FactorialCase struct {
	n      int
	result *big.Int
}

func (suite *CalculatorTestSuite) TestFactorial() {
	ctx := context.Background()
	cases := []FactorialCase{{
		n:      0,
		result: big.NewInt(1),
	}, {
		n:      1,
		result: big.NewInt(1),
	}, {
		n:      10,
		result: big.NewInt(3628800),
	},
	}
	for _, c := range cases {
		suite.Require().Equal(suite.calculator.Factorial(ctx, c.n), c.result)
	}
}

func TestNewCalculator(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))

}
