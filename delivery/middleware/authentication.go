package middleware

import (
	"fmt"
	"net/http"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

func CheckAuth(next httpHandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Auth-String")
		fmt.Println(header)
		if header != "mockPass" {
			fmt.Fprint(w, "Invalid password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
