package controllers

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
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
	r.Get("/", uc.getAll)
	r.Get("/{id}", uc.GetByID)

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

	views.JSON(w, http.StatusCreated, newUser)
}

func (uc *Users) getAll(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	var users []models.User
	var err error

	if len(username) == 0 {
		users, err = uc.store.FindAll()
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
