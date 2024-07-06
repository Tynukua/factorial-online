package main

import (
	"log"
	"net/http"

	"github.com/Tynukua/factorial-online/config"
	"github.com/Tynukua/factorial-online/handlers"
	"github.com/Tynukua/factorial-online/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg := config.NewConfig()
	router := httprouter.New()
	router.GET("/", handlers.Index)
	h := handlers.NewCalculateHandler(cfg)
	router.POST("/calculate", middleware.CalculateCheckInputMiddleware((h.Calculate)))

	log.Println("Server started on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
