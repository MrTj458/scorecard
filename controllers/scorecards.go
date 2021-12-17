package controllers

import (
	"github.com/MrTj458/scorecard/models"
	"github.com/go-chi/chi/v5"
)

type Scorecards struct {
	s *models.ScorecardService
}

func NewScorecards(service *models.ScorecardService) *Scorecards {
	return &Scorecards{
		s: service,
	}
}

func (s *Scorecards) Routes() *chi.Mux {
	r := chi.NewRouter()

	return r
}
