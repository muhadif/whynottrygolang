package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	errprMiddlerWare := NewErrorHandle()
	mux.Handle("/hello", errprMiddlerWare.Handle(GetHello))

	// Starting the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		return
	}
}
