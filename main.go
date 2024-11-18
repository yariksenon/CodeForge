package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
	Age   uint8
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
