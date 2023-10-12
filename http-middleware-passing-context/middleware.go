package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	userKey   = "username"
	userIDKey = "userID"
)

func CoreMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		username := r.Header.Get(userKey)
		userid := r.Header.Get(userIDKey)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userIDKey, userid)
		ctx = context.WithValue(ctx, userKey, username)
		*r = *r.WithContext(ctx)

		next.ServeHTTP(w, r)

		fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
