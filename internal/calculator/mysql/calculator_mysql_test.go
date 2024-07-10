package mysql

import (
	"context"
	"database/sql"
	"github.com/Tynukua/factorial-online/internal/calculator/mathematics"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type CalclualtorSuite struct {
	suite.Suite
	mc *Calculator
}

func (s *CalclualtorSuite) SetupSuite() {

	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/testdb")
	s.Require().NoError(err)
	s.mc = NewCalculator(db, mathematics.NewCalculator())
}

func (s *CalclualtorSuite) TestFactorial() {
	ctx := context.TODO()
	s.Require().Equal(s.mc.Factorial(ctx, 5), big.NewInt(120))
	s.Require().Equal(s.mc.Factorial(ctx, 6), big.NewInt(720))

}

func TestNewCalculator(t *testing.T) {
	suite.Run(t, new(CalclualtorSuite))

}
