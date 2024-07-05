package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"runtime"

	"github.com/Tynukua/factorial-online/database"
	"github.com/Tynukua/factorial-online/math"

	"github.com/julienschmidt/httprouter"
)

type CalculateRequest struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type Handler struct {
	DB database.FactorialDatabase
}
type ContentKey string

const CalculateDataKey ContentKey = "CalculateData"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (handler Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.Context().Value(CalculateDataKey).(CalculateRequest)
	var a, b int = *params.A, *params.B
	af, bf := handler.DoubleFactorial(a, b)
	response := map[string]*big.Int{"a!": af, "b!": bf}
	responsedata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusInternalServerError)
		return
	}
	w.Write(responsedata)
}

func (handler Handler) DoubleFactorial(a int, b int) (*big.Int, *big.Int) {
	var swapped bool
	if a > b {
		a, b = b, a
		swapped = true
	}
	var af, bf *big.Int
	var ac, bc int
	var acf, bcf *big.Int
	ac, acf, _ = handler.DB.GetClosestFactorial(a)
	bc, bcf, _ = handler.DB.GetClosestFactorial(b)
	af = big.NewInt(1)
	bf = big.NewInt(1)

	af.Mul(acf, math.MulRangeParallel(ac+1, a, runtime.NumCPU()))
	if a > bc {
		bc = a
		bcf = af
	}
	bf.Mul(bcf, math.MulRangeParallel(bc+1, b, runtime.NumCPU()))

	handler.DB.SaveFactorial(a, af)
	handler.DB.SaveFactorial(b, bf)
	if swapped {
		af, bf = bf, af
	}
	return af, bf
}
