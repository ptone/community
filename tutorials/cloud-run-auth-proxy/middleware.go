package main

import (
	"fmt"
	"net/http"
)

func authGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			fmt.Println("skipping favicon")
			next.ServeHTTP(w, r)
		}

		_, err := readCookie(r)
		if err == nil {
			next.ServeHTTP(w, r)
			return
		}

		if err != nil {
			fmt.Printf("cookie get error: %s\n", err)
			http.Redirect(w, r, fmt.Sprintf("/_login?next=%s", r.URL.Path), 302)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}
