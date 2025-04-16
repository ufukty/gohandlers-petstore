package main

import (
	"fmt"
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets"
)

func main() {
	s := http.NewServeMux()

	pets := pets.Pets{}
	for name, handler := range pets.ListHandlers() {
		pattern := fmt.Sprintf("%s %s", handler.Method, handler.Path)
		fmt.Println("registering", name, "as", pattern)
		s.HandleFunc(pattern, handler.Ref)
	}
}
