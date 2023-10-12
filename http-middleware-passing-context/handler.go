package main

import (
	"encoding/json"
	"fmt"
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

	user := getContextValueByKey(req.Context(), userKey, "string")
	hello := &Hello{Message: fmt.Sprintf("Hi, %s", user)}

	return json.NewEncoder(w).Encode(hello)
}
