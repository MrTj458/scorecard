package api

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) registerUserRoutes() {
	s.mux.Route("/api/users", func(r chi.Router) {
		r.Post("/", s.createUser)
		r.With(requireLoginMiddleware).Get("/", s.getAllUsers)
		r.With(requireLoginMiddleware).Get("/{id}", s.getUserByID)
		r.With(requireLoginMiddleware).Get("/me", s.me)
		r.Post("/login", s.logInUser)
		r.Post("/logout", s.logOutUser)
	})
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	// Decode JSON from request body
	var user models.UserIn
	if err := decodeJSON(r.Body, &user); err != nil {
		renderError(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate received JSON
	if errors, ok := validateStruct(user); !ok {
		renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid user object received", errors)
		return
	}

	// Create new user
	newUser, err := s.UserService.Add(user)
	if err != nil {
		switch err {
		case models.ErrEmailInUse:
			renderError(w, http.StatusUnprocessableEntity, fmt.Sprintf("email '%s' already in use", user.Email))
		case models.ErrUsernameInUse:
			renderError(w, http.StatusUnprocessableEntity, fmt.Sprintf("username '%s' already in use", user.Username))
		default:
			renderError(w, http.StatusInternalServerError, "error creating new user")
		}
		return
	}

	cookie := newCookie(AuthCookieMaxAge, "Auth", newUser.ID.Hex())
	http.SetCookie(w, cookie)

	renderJSON(w, http.StatusCreated, newUser)
}

func (s *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	var users []models.User
	var err error

	if len(username) == 0 {
		renderError(w, http.StatusForbidden, "must include username to query for")
		return
	} else {
		users, err = s.UserService.SearchByUsername(username)
	}

	if err != nil {
		renderError(w, http.StatusInternalServerError, "error retrieving users")
		return
	}

	renderJSON(w, http.StatusOK, users)
}

func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := s.UserService.FindByID(id)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("user with id '%s' not found", id))
		return
	}

	renderJSON(w, http.StatusOK, u)
}

func (s *Server) me(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(string)

	u, err := s.UserService.FindByID(id)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("user with id '%s' not found", id))
		return
	}

	renderJSON(w, http.StatusOK, u)
}

func (s *Server) logInUser(w http.ResponseWriter, r *http.Request) {
	// Decode login object
	var login models.UserLogin
	if err := decodeJSON(r.Body, &login); err != nil {
		renderError(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate login object
	if errors, ok := validateStruct(login); !ok {
		renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid login object received", errors)
		return
	}
	fmt.Println("Value of user service: ", s.UserService)
	// Find a user with the login email
	u, err := s.UserService.FindByEmail(login.Email)
	if err != nil {
		renderError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(login.Password))
	if err != nil {
		renderError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	cookie := newCookie(AuthCookieMaxAge, "Auth", u.ID.Hex())
	http.SetCookie(w, cookie)

	renderJSON(w, http.StatusOK, u)
}

func (s *Server) logOutUser(w http.ResponseWriter, r *http.Request) {
	cookie := newCookie(-1, "Auth", "")
	http.SetCookie(w, cookie)

	renderJSON(w, http.StatusOK, map[string]interface{}{"ok": true})
}
