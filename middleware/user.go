package middleware

import (
	"context"
	"net/http"

	"github.com/MrTj458/scorecard/views"
)

func User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		c, err := r.Cookie("Auth")
		if err != nil {
			ctx = context.WithValue(r.Context(), "user", "")
		} else {
			ctx = context.WithValue(r.Context(), "user", c.Value)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uId := r.Context().Value("user").(string)

		if len(uId) == 0 {
			views.Error(w, http.StatusUnauthorized, "you must be signed in to access this route")
			return
		}

		next.ServeHTTP(w, r)
	})
}
