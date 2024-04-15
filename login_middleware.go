package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pranavtaysheti/goTTT/templating"
)

func LoginCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Passing logincookiemiddleware")
		c, err := getClientFromRequest(r)
		if err != nil {
			templating.ExecuteLayout(w, nil, "notpermitted.html")
		}

		ctx := context.WithValue(r.Context(), "client", c)
		log.Println("Added client to context successfully")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
