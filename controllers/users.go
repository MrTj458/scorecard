package controllers

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/middleware"
	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	store *models.UserStore
}

func NewUsers(service *models.UserStore) *Users {
	return &Users{
		store: service,
	}
}

func (uc *Users) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", uc.create)
	r.With(middleware.RequireLogin).Get("/", uc.getAll)
	r.With(middleware.RequireLogin).Get("/{id}", uc.GetByID)
	r.With(middleware.RequireLogin).Get("/me", uc.me)
	r.Post("/login", uc.logIn)
	r.Post("/logout", uc.LogOut)

	return r
}

func (uc *Users) create(w http.ResponseWriter, r *http.Request) {
	// Decode JSON from request body
	var user models.UserIn
	if err := views.DecodeJSON(r.Body, &user); err != nil {
		views.Error(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate received JSON
	if vErrors := views.Validate(user); vErrors != nil {
		views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid user object received", vErrors)
		return
	}

	// Create new user
	newUser, err := uc.store.Add(user)
	if err != nil {
		switch err {
		case models.ErrEmailInUse:
			views.Error(w, http.StatusUnprocessableEntity, fmt.Sprintf("email '%s' already in use", user.Email))
		case models.ErrUsernameInUse:
			views.Error(w, http.StatusUnprocessableEntity, fmt.Sprintf("username '%s' already in use", user.Username))
		default:
			views.Error(w, http.StatusInternalServerError, "error creating new user")
		}
		return
	}

	// Set auth cookie
	cookie := &http.Cookie{
		Name:     "Auth",
		Value:    newUser.ID.Hex(),
		MaxAge:   86400,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	views.JSON(w, http.StatusCreated, newUser)
}

func (uc *Users) getAll(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	var users []models.User
	var err error

	if len(username) == 0 {
		views.Error(w, http.StatusForbidden, "must include username to query for")
		return
	} else {
		users, err = uc.store.SearchByUsername(username)
	}

	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error retrieving users")
		return
	}

	views.JSON(w, http.StatusOK, users)
}

func (uc *Users) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := uc.store.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusNotFound, fmt.Sprintf("user with id '%s' not found", id))
		return
	}

	views.JSON(w, http.StatusOK, u)
}

func (uc *Users) me(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(string)

	u, err := uc.store.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusNotFound, fmt.Sprintf("user with id '%s' not found", id))
		return
	}

	views.JSON(w, http.StatusOK, u)
}

func (uc *Users) logIn(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// Decode login object
	var login form
	if err := views.DecodeJSON(r.Body, &login); err != nil {
		views.Error(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate login object
	if vErrors := views.Validate(login); vErrors != nil {
		views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid login object received", vErrors)
		return
	}

	// Find a user with the login email
	u, err := uc.store.FindByEmail(login.Email)
	if err != nil {
		views.Error(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(login.Password))
	if err != nil {
		views.Error(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Set auth cookie
	cookie := &http.Cookie{
		Name:     "Auth",
		Value:    u.ID.Hex(),
		MaxAge:   86400,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	views.JSON(w, http.StatusOK, u)
}

func (uc *Users) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "Auth",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	views.JSON(w, http.StatusOK, views.M{"ok": true})
}
