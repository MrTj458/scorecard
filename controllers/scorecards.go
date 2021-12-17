package controllers

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
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

	r.Post("/", s.create)
	r.Get("/", s.findAll)
	r.Get("/{id}", s.findByID)

	return r
}

func (sc *Scorecards) create(w http.ResponseWriter, r *http.Request) {
	var s models.ScorecardIn
	if err := views.DecodeJSON(r.Body, &s); err != nil {
		views.Error(w, http.StatusUnprocessableEntity, "invalid scorecard object received")
		return
	}

	if vErrors := views.Validate(s); vErrors != nil {
		views.FieldError(w, http.StatusUnprocessableEntity, "invalid scorecard object received", vErrors)
		return
	}

	id, err := sc.s.Add(s)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to create new user")
		return
	}

	ret, err := sc.s.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to find created scorecard")
		return
	}

	views.JSON(w, http.StatusCreated, ret)
}

func (sc *Scorecards) findAll(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	var scorecards []models.Scorecard
	var err error

	if len(user) == 0 {
		scorecards, err = sc.s.FindAll()
	} else {
		scorecards, err = sc.s.FindAllByUserId(user)
	}

	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error retrieving scorecards")
		return
	}

	views.JSON(w, http.StatusOK, scorecards)
}

func (sc *Scorecards) findByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	s, err := sc.s.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusNotFound, fmt.Sprintf("scorecard with id '%s' not found", id))
		return
	}

	views.JSON(w, http.StatusOK, s)
}
