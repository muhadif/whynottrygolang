package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Use(CoreMiddleWare)

	errMiddleWare := NewErrorHandle()
	r.Handle("/hello", errMiddleWare.Handle(GetHello)).Methods(http.MethodGet)
	r.Handle("/hello", errMiddleWare.Handle(CreateHello)).Methods(http.MethodPost)
	r.Handle("/hello", errMiddleWare.Handle(UpdateHello)).Methods(http.MethodPut)

	// Starting the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
