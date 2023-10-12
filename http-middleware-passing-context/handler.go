package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type CreateHelloRequest struct {
	IsNeedError bool
}

// CreateHello Use io.ReadAll
func CreateHello(w http.ResponseWriter, req *http.Request) error {
	var payload *CreateHelloRequest
	if err := decodeData(req, &payload); err != nil {
		err := CustomError{
			Code:    http.StatusInternalServerError,
			Message: "error decode data",
		}
		return err
	}

	if payload.IsNeedError {
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

// UpdateHello Use json decoder
func UpdateHello(w http.ResponseWriter, req *http.Request) error {
	var payload *CreateHelloRequest
	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(req.Body, b)
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&payload); err != nil {
		err := CustomError{
			Code:    http.StatusInternalServerError,
			Message: "error decode data",
		}
		return err
	}
	req.Body = io.NopCloser(b)

	if payload.IsNeedError {
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

func decodeData(req *http.Request, resp interface{}) error {
	bb, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	req.Body = io.NopCloser(io.NopCloser(bytes.NewBuffer(bb)))

	if err := json.Unmarshal(bb, &resp); err != nil {
		return err
	}

	return nil
}
