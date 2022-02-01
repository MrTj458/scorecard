package api

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerScorecardRoutes() {
	s.mux.Route("/api/scorecards", func(r chi.Router) {
		r.Use(requireLoginMiddleware)

		r.Post("/", s.createScorecard)
		r.Get("/", s.getAllScorecards)
		r.Get("/{id}", s.getScorecardByID)
		r.Post("/{cardId}/hole", s.addHoleToScorecard)
		r.Post("/{cardId}/complete", s.completeScorecard)
		r.Delete("/{cardId}", s.deleteScorecard)
	})
}

func (s *Server) createScorecard(w http.ResponseWriter, r *http.Request) {
	var sc models.ScorecardIn
	if err := decodeJSON(r.Body, &sc); err != nil {
		renderError(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	if errors, ok := validateStruct(sc); !ok {
		renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid scorecard object received", errors)
		return
	}

	newCard, err := s.ScorecardService.Add(sc)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "unable to create new user")
		return
	}

	renderJSON(w, http.StatusCreated, newCard)
}

func (s *Server) getAllScorecards(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(string)

	scorecards, err := s.ScorecardService.FindAllByUserId(userId)

	if err != nil {
		renderError(w, http.StatusInternalServerError, "error retrieving scorecards")
		return
	}

	renderJSON(w, http.StatusOK, scorecards)
}

func (s *Server) getScorecardByID(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(string)
	id := chi.URLParam(r, "id")

	sc, err := s.ScorecardService.FindByID(id)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("scorecard with id '%s' not found", id))
		return
	}

	permission := false
	for _, player := range sc.Players {
		if player.ID.Hex() == userId {
			permission = true
			break
		}
	}
	if !permission {
		renderError(w, http.StatusForbidden, "you don't have access to this scorecard")
		return
	}

	renderJSON(w, http.StatusOK, sc)
}

func (s *Server) addHoleToScorecard(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "cardId")
	userId := r.Context().Value("user").(string)

	card, err := s.ScorecardService.FindByID(cardId)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("scorecard with id '%s' not found", cardId))
		return
	}

	if card.CreatedBy.Hex() != userId {
		renderError(w, http.StatusForbidden, "you can't add holes to this scorecard")
		return
	}

	var h models.Hole
	if err := decodeJSON(r.Body, &h); err != nil {
		renderError(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	// Validate hole object
	if errors, ok := validateStruct(h); !ok {
		renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid hole object received", errors)
		return
	}

	// Validate nested score objects
	for _, s := range h.Scores {
		if errors, ok := validateStruct(s); !ok {
			renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid score object received", errors)
			return
		}
	}

	card, err = s.ScorecardService.AddHole(cardId, h)
	if err != nil {
		renderError(w, http.StatusBadRequest, "unable to find scorecard to add hole")
		return
	}

	renderJSON(w, http.StatusCreated, card)
}

func (s *Server) completeScorecard(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "cardId")
	userId := r.Context().Value("user").(string)

	card, err := s.ScorecardService.FindByID(cardId)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("scorecard with id '%s' not found", cardId))
		return
	}
	if card.CreatedBy.Hex() != userId {
		renderError(w, http.StatusForbidden, "you can't mark this card as complete")
		return
	}

	card, err = s.ScorecardService.Complete(cardId)
	if err != nil {
		renderError(w, http.StatusBadRequest, "unable to find scorecard to complete")
		return
	}

	renderJSON(w, http.StatusOK, card)
}

func (s *Server) deleteScorecard(w http.ResponseWriter, r *http.Request) {
	cardId := chi.URLParam(r, "cardId")
	userId := r.Context().Value("user").(string)

	card, err := s.ScorecardService.FindByID(cardId)
	if err != nil {
		renderError(w, http.StatusNotFound, "unable to find scorecard")
		return
	}

	if userId != card.CreatedBy.Hex() {
		renderError(w, http.StatusForbidden, "you do not have permission to delete this scorecard")
		return
	}

	err = s.ScorecardService.Delete(cardId)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "unable to delete scorecard")
		return
	}

	renderJSON(w, http.StatusOK, map[string]interface{}{"ok": true})
}
