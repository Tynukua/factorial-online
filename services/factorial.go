package services

import (
	"math/big"
	"runtime"

	"github.com/Tynukua/factorial-online/config"
	"github.com/Tynukua/factorial-online/database"
	"github.com/Tynukua/factorial-online/util"
)

type FactorialService struct {
	db database.FactorialDatabase
}

func NewFactorialService(cfg config.Config) FactorialService {
	switch expression := cfg.DBType; expression {
	case config.MySQL:
		return FactorialService{db: database.NewMySQLFactorialDatabase(cfg.DSN)}
	default:
		return FactorialService{db: database.NewMemoryFactorialDatabase()}
	}
}

func (s FactorialService) DoubleFactorial(a int, b int) (*big.Int, *big.Int) {
	var swapped bool
	if a > b {
		a, b = b, a
		swapped = true
	}
	var af, bf *big.Int
	var ac, bc int
	var acf, bcf *big.Int
	ac, acf, _ = s.db.GetClosestFactorial(a)
	bc, bcf, _ = s.db.GetClosestFactorial(b)
	af = big.NewInt(1)
	bf = big.NewInt(1)

	af.Mul(acf, util.MulRangeParallel(ac+1, a, runtime.NumCPU()))
	if a > bc {
		bc = a
		bcf = af
	}
	bf.Mul(bcf, util.MulRangeParallel(bc+1, b, runtime.NumCPU()))

	s.db.SaveFactorial(a, af)
	s.db.SaveFactorial(b, bf)
	if swapped {
		af, bf = bf, af
	}
	return af, bf
}
