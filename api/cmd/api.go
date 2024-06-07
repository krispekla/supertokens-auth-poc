package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Get("/another-route", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is another route"))
	})

	http.ListenAndServe(":8080", r)
}
