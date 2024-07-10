package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type CalclualtorSuite struct {
	suite.Suite
	as *AsyncService
}

func (s *CalclualtorSuite) SetupSuite() {
	s.as = NewAsyncService()
}

func TestCalculatorService(t *testing.T) {
	suite.Run(t, new(CalclualtorSuite))
}

func (s *CalclualtorSuite) TestDo() {
	ctx := context.Background()
	var i int
	m := sync.Mutex{}
	f := func() {
		m.Lock()
		i += 1
		m.Unlock()
	}
	g := func() {
		m.Lock()
		i += 1
		m.Unlock()
	}
	fs := []func(){f, g}

	err := s.as.Do(ctx, fs)
	s.Require().NoError(err)
	s.Require().Equal(i, len(fs))
}
