package router

import (
	"github.com/Tynukua/factorial-online/internal/config"
	"github.com/Tynukua/factorial-online/internal/handlers"
	"github.com/Tynukua/factorial-online/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func SetupRouter(cfg config.Config) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	h := handlers.NewCalculateHandler(cfg)
	router.POST("/calculate", middleware.CalculateCheckInputMiddleware(h.Calculate))
	return router
}
