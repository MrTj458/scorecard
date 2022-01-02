package controllers

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
)

type Scorecards struct {
	store *models.ScorecardStore
}

func NewScorecards(service *models.ScorecardStore) *Scorecards {
	return &Scorecards{
		store: service,
	}
}

func (s *Scorecards) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", s.create)
	r.Get("/", s.getAll)
	r.Get("/{id}", s.getByID)
	r.Post("/{cardId}/hole", s.addHole)
	r.Post("/{cardId}/complete", s.complete)

	return r
}

func (sc *Scorecards) create(w http.ResponseWriter, r *http.Request) {
	var s models.ScorecardIn
	if err := views.DecodeJSON(r.Body, &s); err != nil {
		views.Error(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	if vErrors := views.Validate(s); vErrors != nil {
		views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid scorecard object received", vErrors)
		return
	}

	newCard, err := sc.store.Add(s)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to create new user")
		return
	}

	views.JSON(w, http.StatusCreated, newCard)
}

func (sc *Scorecards) getAll(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	var scorecards []models.Scorecard
	var err error

	if len(user) == 0 {
		scorecards, err = sc.store.FindAll()
	} else {
		scorecards, err = sc.store.FindAllByUserId(user)
	}

	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error retrieving scorecards")
		return
	}

	views.JSON(w, http.StatusOK, scorecards)
}

func (sc *Scorecards) getByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	s, err := sc.store.FindByID(id)
	if err != nil {
		views.Error(w, http.StatusNotFound, fmt.Sprintf("scorecard with id '%s' not found", id))
		return
	}

	views.JSON(w, http.StatusOK, s)
}

func (sc *Scorecards) addHole(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "cardId")

	var h models.Hole
	if err := views.DecodeJSON(r.Body, &h); err != nil {
		views.Error(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate hole object
	if vErrors := views.Validate(h); vErrors != nil {
		views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid hole object received", vErrors)
		return
	}

	// Validate nested score objects
	for _, s := range h.Scores {
		if vErrors := views.Validate(s); vErrors != nil {
			views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid score object received", vErrors)
			return
		}
	}

	card, err := sc.store.AddHole(cardId, h)
	if err != nil {
		views.Error(w, http.StatusBadRequest, "unable to find scorecard to add hole")
		return
	}

	views.JSON(w, http.StatusCreated, card)
}

func (sc *Scorecards) complete(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "cardId")

	card, err := sc.store.Complete(cardId)
	if err != nil {
		views.Error(w, http.StatusBadRequest, "unable to find scorecard to complete")
		return
	}

	views.JSON(w, http.StatusOK, card)
}
