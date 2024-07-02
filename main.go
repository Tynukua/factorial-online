package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
		params := CalculateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			w.Write([]byte("Invid JSON resourse"))
			return
		}
		if params.A == nil || params.B == nil {
			w.Write([]byte("Params a and b must be provided"))
			return
		}
		if *params.A < 0 || *params.B < 0 {
			w.Write([]byte("Value must be non-negative"))
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

	a, b = doubleFactorial(a, b)

	var response map[string]int = map[string]int{"a!": a, "b!": b}
	responsedata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Write(responsedata)
}

func doubleFactorial(a int, b int) (int, int) {
	if a <= 1 && b <= 1 {
		return 1, 1
	}
	channel := make(chan int)
	period := func(start int, end int) {
		if start > end {
			end, start = start, end
		}
		answer := 1
		for i := start; i <= end; i++ {
			answer *= i
		}
		channel <- answer
	}
	go period(min(a, b)+1, max(a, b))
	// from zero to min(a, b)
	go period(1, min(a, b))

	m, n := <-channel, <-channel
	if a > b {
		return n * m, m
	} else if a == b {
		return m, m
	} else {
		return m, n * m
	}
}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/calculate", CalculateCheckInputMiddleware(Calculate))
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
