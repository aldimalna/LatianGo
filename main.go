package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Server is starting.......")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello-world", getHandler)
	r.Get("/api/auth", userAuth)
	fmt.Println("Server is listening on port 10000...")
	log.Fatal(http.ListenAndServe(":10000", r))

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world")
}

func userAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("this is user authentication apge")
}
