package main

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string
}

func GetHello(w http.ResponseWriter, req *http.Request) error {
	if req.URL.Query().Get("isNeedError") == "true" {
		err := CustomError{
			Code:    http.StatusInternalServerError,
			Message: "error happen like you want",
		}

		return err
	}

	hello := &Hello{Message: "Hi, Im Here"}

	return json.NewEncoder(w).Encode(hello)
}
