package controllers

import (
	"github.com/MrTj458/scorecard/models"
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

func (u *Users) Routes() *chi.Mux {
	r := chi.NewRouter()

	return r
}
