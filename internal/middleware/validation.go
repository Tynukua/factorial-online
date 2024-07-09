package middleware

import (
	"context"
	"encoding/json"
	"github.com/Tynukua/factorial-online/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CalculateCheckInputMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		params := handlers.CalculateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}
		if params.A == nil || params.B == nil || *params.A < 0 || *params.B < 0 {
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}

		ctx := context.WithValue(r.Context(), handlers.CalculateDataKey, params)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}
