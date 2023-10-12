package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string
}

const (
	userKey = "user"
)

func GetHello(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userKey, "adif")
	*req = *req.WithContext(ctx)

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
