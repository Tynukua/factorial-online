package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

type CalculateRequest struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

func CalculateCheckInputMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		params := CalculateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}
		if params.A == nil || params.B == nil || *params.A < 0 || *params.B < 0 {
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}

		ctx := context.WithValue(r.Context(), "CalculateData", params)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}
func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.Context().Value("CalculateData").(CalculateRequest)
	var a, b int = *params.A, *params.B
	var af, bf *big.Int
	if a < b {
		af = MulRangeParallel(1, a, 2)
		bf = big.NewInt(1)
		bf.Mul(af, MulRangeParallel(a+1, b, 2))
	} else if a == b {
		af = MulRangeParallel(1, a, 2)
		bf = af
	} else {
		bf = MulRangeParallel(1, b, 2)
		af = big.NewInt(1)
		af.Mul(bf, MulRangeParallel(b+1, a, 2))
	}

	response := map[string]*big.Int{"a!": af, "b!": bf}
	responsedata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusInternalServerError)
		return
	}
	w.Write(responsedata)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/calculate", CalculateCheckInputMiddleware(Calculate))
	log.Println("Server started on port 8989")
	log.Fatal(http.ListenAndServe(":8989", router))
}
