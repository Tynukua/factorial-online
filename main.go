package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

type CalculateRequest struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type Handler struct {
	db *sql.DB
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
func (handler Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.Context().Value("CalculateData").(CalculateRequest)
	var a, b int = *params.A, *params.B
	var swapped bool
	if a > b {
		a, b = b, a
		swapped = true
	}
	var af, bf *big.Int
	var ac, bc int
	var acf, bcf *big.Int
	var err error
	ac, acf, err = GetClosestFactorial(handler.db, a)
	bc, bcf, err = GetClosestFactorial(handler.db, b)
	af = big.NewInt(1)
	bf = big.NewInt(1)

	af.Mul(acf, MulRangeParallel(ac, a, 2))
	if a > bc {
		bc = a
		bcf = af
	}
	bf.Mul(bcf, MulRangeParallel(bc+1, b, 2))

	SaveFactorialToDatabase(handler.db, a, af)
	SaveFactorialToDatabase(handler.db, b, bf)
	if swapped {
		af, bf = bf, af
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
	h := Handler{}
	var err error = nil
	h.db, err = sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
		return
	}
	InitDatabase(h.db)
	router.POST("/calculate", CalculateCheckInputMiddleware(h.Calculate))

	log.Println("Server started on port 8989")
	log.Fatal(http.ListenAndServe(":8989", router))
}
