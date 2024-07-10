package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type AsyncServiceTestSuite struct {
	suite.Suite
	s *AsyncService
}

func (suite *AsyncServiceTestSuite) SetupSuite() {
	suite.s = New()
}

func TestCalculatorService(t *testing.T) {
	suite.Run(t, new(AsyncServiceTestSuite))
}

func (suite *AsyncServiceTestSuite) TestDo() {
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

	err := suite.s.Do(ctx, fs)
	suite.Require().NoError(err)
	suite.Require().Equal(i, len(fs))
}
