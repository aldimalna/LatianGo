package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the Homepage!")
	fmt.Println("endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("The server is On")
	handleRequests()

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world")
}

func userAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("this is user authentication apge")
}
