package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/Tynukua/factorial-online/config"
	"github.com/Tynukua/factorial-online/services"

	"github.com/julienschmidt/httprouter"
)

type CalculateRequest struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type Handler struct {
	service services.FactorialService
}

func NewCalculateHandler(cfg config.Config) Handler {
	return Handler{service: services.NewFactorialService(cfg)}
}

type ContentKey string

const CalculateDataKey ContentKey = "CalculateData"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (handler Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.Context().Value(CalculateDataKey).(CalculateRequest)
	var a, b int = *params.A, *params.B
	af, bf := handler.service.DoubleFactorial(a, b)
	response := map[string]*big.Int{"a!": af, "b!": bf}
	responsedata, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusInternalServerError)
		return
	}
	w.Write(responsedata)
}
