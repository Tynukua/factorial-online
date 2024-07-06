package main

import (
	"log"
	"net/http"

	"github.com/Tynukua/factorial-online/config"
	"github.com/Tynukua/factorial-online/router"
)

func main() {
	cfg := config.NewConfig()
	r := router.SetupRouter(cfg)

	log.Println("Server started on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
