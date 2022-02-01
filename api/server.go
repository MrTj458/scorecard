package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/MrTj458/scorecard/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

const (
	AuthCookieMaxAge = 86400
)

type Server struct {
	mux  *chi.Mux
	port string

	UserService      models.UserService
	ScorecardService models.ScorecardService
	DiscService      models.DiscService
}

func NewServer(port string) *Server {
	s := &Server{
		mux:  chi.NewRouter(),
		port: port,
	}

	// Global middleware
	s.mux.Use(middleware.Logger)
	s.mux.Use(addUserToContextMiddleware)

	// Register Routes
	s.registerUserRoutes()
	s.registerScorecardRoutes()
	s.registerDiscRoutes()

	// Set up file server
	s.mux.Handle("/*", http.FileServer(http.Dir("static")))

	// Set 404 and 405 error responses
	s.mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		renderError(w, http.StatusNotFound, "route not found")
	})

	s.mux.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		renderError(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+s.port, s.mux)
}

// newCookie returns a pointer to a new http.Cookie with default values.
func newCookie(maxAge int, name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
}

// addUserToContextMiddleware checks for an auth cookie and adds the user
// ID to the context if it exists.
func addUserToContextMiddleware(next http.Handler) http.Handler {
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

// requireLoginMiddleware checks if there is an auth cookie and returns a
// 401 response if it doesn't exist.
func requireLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uId := r.Context().Value("user").(string)

		if len(uId) == 0 {
			renderError(w, http.StatusUnauthorized, "you must be signed in to access this route")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// decodeJSON decodes the JSON contents of r into the data interface given.
func decodeJSON(r io.Reader, out interface{}) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&out); err != nil {
		return err
	}
	return nil
}

// renderJSON writes the given status code and interface to w as a JSON object.
func renderJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.Encode(data)
}

// renderError writes the given status code and detail string to w as JSON.
func renderError(w http.ResponseWriter, status int, detail string) {
	res := models.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     make([]*models.ErrorField, 0),
	}
	renderJSON(w, status, res)
}

// renderErrorWithFields writes the given status code and detail string, along with a list
// of `models.ErrorField`s to w as JSON.
func renderErrorWithFields(w http.ResponseWriter, status int, detail string, fields []*models.ErrorField) {
	res := models.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     fields,
	}
	renderJSON(w, status, res)
}

// validateStruct validates the given struct with the validator package.
// It returns an models.ErrorField slice full of errors, along with weather
// or not There were any errors.
func validateStruct(s interface{}) ([]*models.ErrorField, bool) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(s)
	if err != nil {
		errors := []*models.ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &models.ErrorField{
				Location: err.Field(),
				Type:     err.Type().String(),
				Detail:   err.ActualTag(),
			})
		}

		return errors, false
	}

	return nil, true
}
