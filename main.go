package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Tynukua/factorial-online/database"
	"github.com/Tynukua/factorial-online/handlers"
	"github.com/Tynukua/factorial-online/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	h := handlers.Handler{}
	h.DB = database.NewMySQLFactorialDatabase(os.Getenv("MYSQL_DSN"))
	// h.db = NewMemoryFactorialDatabase()
	h.DB.InitDatabase()
	router.POST("/calculate", middleware.CalculateCheckInputMiddleware((h.Calculate)))

	log.Println("Server started on port 8989")
	log.Fatal(http.ListenAndServe(":8989", router))
}
