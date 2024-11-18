package main

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
	Age   uint8
}

func main() {
	fmt.Println("Hello")
}
