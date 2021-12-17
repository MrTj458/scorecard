package controllers

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
)

type Users struct {
	s *models.UserService
}

func NewUsers(service *models.UserService) *Users {
	return &Users{
		s: service,
	}
}

func (uc *Users) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", uc.findAll)
	r.Post("/", uc.create)

	return r
}

func (uc *Users) create(w http.ResponseWriter, r *http.Request) {
	// Decode JSON from request body
	var u models.UserIn
	if err := views.DecodeJSON(r.Body, &u); err != nil {
		views.Error(w, http.StatusUnprocessableEntity, "invalid user object received")
		return
	}

	// Validate received JSON
	if vErrors := views.Validate(u); vErrors != nil {
		views.FieldError(w, http.StatusUnprocessableEntity, "invalid user object received", vErrors)
		return
	}

	// Check if username or email already exists
	existingUsers, err := uc.s.FindExistingUsers(u.Email, u.Username)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error checking for existing users")
		return
	}

	if len(existingUsers) != 0 {
		// Build error fields
		var errors []views.ErrorField
		for _, eUser := range existingUsers {
			if u.Email == eUser.Email {
				// Email already in use
				errors = append(errors, views.ErrorField{
					Location: "email",
					Type:     "string",
					Detail:   fmt.Sprintf("a user already exists with the email '%s'", u.Email),
				})
			}

			if u.Username == eUser.Username {
				// Username already in use
				errors = append(errors, views.ErrorField{
					Location: "username",
					Type:     "string",
					Detail:   fmt.Sprintf("a user already exists with the username '%s'", u.Username),
				})
			}
		}

		views.FieldError(w, http.StatusBadRequest, "email or username already in use", errors)
		return
	}

	// Create new user
	id, err := uc.s.Add(u)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to create new user")
		return
	}

	// Find newly created user
	ret, err := uc.s.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to find created user")
		return
	}

	views.JSON(w, http.StatusCreated, ret)
}

func (uc *Users) findAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.s.FindAll()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error retrieving users")
		return
	}

	views.JSON(w, http.StatusOK, users)
}
