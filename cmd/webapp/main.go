package main

import (
	"github.com/Tynukua/factorial-online/internal/config"
	"github.com/Tynukua/factorial-online/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	r := router.SetupRouter(cfg)

	log.Println("Server started on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
