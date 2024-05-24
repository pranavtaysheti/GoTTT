package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pranavtaysheti/GoTTT/internal/templating"
)

func LoginCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := getClientFromRequest(r)
		if err != nil {
			templating.Render(w, "Not Permitted",nil, "notpermitted.html")
		}

		ctx := context.WithValue(r.Context(), "client", c)
		log.Println("Added client to context successfully")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
