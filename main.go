package main

import (
	"fmt"
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets"
)

func main() {
	var db any
	pets := pets.New(db)

	// other startup code

	s := http.NewServeMux()
	for name, handler := range pets.ListHandlers() {
		pattern := fmt.Sprintf("%s %s", handler.Method, handler.Path)
		fmt.Println("registering", name, "as", pattern)
		s.HandleFunc(pattern, handler.Ref)
	}

	_ = http.ListenAndServe(":3000", s)
}
