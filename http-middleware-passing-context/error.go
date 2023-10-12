package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (o CustomError) Error() string {
	return fmt.Sprintf("CustomError code = %v desc - %v errors = %v", o.Code, o.Message)
}

type ErrHandler func(http.ResponseWriter, *http.Request) error

type ErrHandle struct {
	fn ErrHandler
}

func NewErrorHandle() ErrHandle {
	return ErrHandle{}
}
func (o ErrHandle) Handle(fn ErrHandler) ErrHandle {
	o.fn = fn
	return o
}

func (o ErrHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := o.fn(w, r); err != nil {
		var errResponse CustomError
		ok := errors.As(err, &errResponse)
		if !ok {
			errResponse = CustomError{
				Code:    http.StatusInternalServerError,
				Message: "error decode custom error",
			}
		}
		w.WriteHeader(errResponse.Code)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
}
