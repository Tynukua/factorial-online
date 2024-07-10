package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type AsyncServiceTestSuite struct {
	suite.Suite
	as *AsyncService
}

func (s *AsyncServiceTestSuite) SetupSuite() {
	s.as = New()
}

func TestCalculatorService(t *testing.T) {
	suite.Run(t, new(AsyncServiceTestSuite))
}

func (s *AsyncServiceTestSuite) TestDo() {
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
