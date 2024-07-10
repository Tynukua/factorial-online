package test

import (
	"context"
	"github.com/Tynukua/factorial-online/internal/calculator/mathematics"
	"github.com/Tynukua/factorial-online/internal/calculator/mysql"
	"github.com/Tynukua/factorial-online/internal/calculator/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type CalclualtorSuite struct {
	suite.Suite
	as service.AsyncService
	mc mysql.MysqlCalculator
}

func (s *CalclualtorSuite) SetupSuite() {
	s.mc = mysql.NewMysqlCalculator("root:example@tcp(localhost:3306)/testdb", mathematics.MathCalculator{})
}

func TestCalculatorService(t *testing.T) {
	suite.Run(t, new(CalclualtorSuite))
}

func (s *CalclualtorSuite) TestDo() {
	ctx := context.Background()
	f := func() {
		log.Println(s.mc.Factorial(ctx, 5555))
	}
	g := func() {
		log.Println(s.mc.Factorial(ctx, 6666))
	}
	fs := []func(){f, g}

	err := s.as.Do(ctx, fs)
	s.Require().NoError(err)
}
