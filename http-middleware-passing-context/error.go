package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

		if r.Method != http.MethodGet {
			bb, err := io.ReadAll(r.Body)
			if err != nil {
				errResponse = CustomError{
					Code:    http.StatusInternalServerError,
					Message: "error decode custom error",
				}
			}
			defer r.Body.Close()
			fmt.Println(string(bb))
		}

		user := getContextValueByKey(r.Context(), userKey, "string")
		userID := getContextValueByKey(r.Context(), userIDKey, "int")
		errResponse.Message = errResponse.Message + fmt.Sprintf("happened for %d : %s", userID, user)

		w.WriteHeader(errResponse.Code)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func getContextValueByKey(ctx context.Context, key string, returnType string) any {
	var value any
	if value = ctx.Value(key); value == nil {
		return nil
	}
	switch returnType {
	case "string":
		return value.(string)
	case "int":
		number, _ := strconv.Atoi(value.(string))
		return number
	}

	return ""
}
