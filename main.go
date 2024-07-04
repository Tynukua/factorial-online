package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type CalculateRequest struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type Handler struct {
	db FactorialDatabase
}
type ContentKey string

const CalculateDataKey ContentKey = "CalculateData"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
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

		ctx := context.WithValue(r.Context(), CalculateDataKey, params)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}
func (handler Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.Context().Value(CalculateDataKey).(CalculateRequest)
	var a, b int = *params.A, *params.B
	af, bf := DoubleFactorial(handler.db, a, b)
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
	h.db = NewMySQLFactorialDatabase(os.Getenv("MYSQL_DSN"))

	h.db.InitDatabase()
	router.POST("/calculate", CalculateCheckInputMiddleware(h.Calculate))

	log.Println("Server started on port 8989")
	log.Fatal(http.ListenAndServe(":8989", router))
}
