package api

import (
	"fmt"
	"net/http"

	"github.com/MrTj458/scorecard/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerDiscRoutes() {
	s.mux.Route("/api/discs", func(r chi.Router) {
		r.Use(requireLoginMiddleware)

		r.Post("/", s.createDisc)
		r.Get("/", s.getAllDiscs)
		r.Get("/{id}", s.getOneDisc)
	})
}

func (s *Server) createDisc(w http.ResponseWriter, r *http.Request) {
	var disc models.DiscIn
	if err := decodeJSON(r.Body, &disc); err != nil {
		renderError(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	if errors, ok := validateStruct(disc); !ok {
		renderErrorWithFields(w, http.StatusUnprocessableEntity, "invalid disc object received", errors)
	}

	discs, err := s.DiscService.Add(disc)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "error fetching discs")
		return
	}

	renderJSON(w, http.StatusOK, discs)
}

func (s *Server) getAllDiscs(w http.ResponseWriter, r *http.Request) {
	uId := r.Context().Value("user").(string)

	discs, err := s.DiscService.FindAllByUserId(uId)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "unable to fetch discs")
		return
	}

	renderJSON(w, http.StatusOK, discs)
}

func (s *Server) getOneDisc(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(string)
	id := chi.URLParam(r, "id")

	d, err := s.DiscService.FindOneById(id)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("disc with id '%s' not found", id))
		return
	}

	if d.CreatedBy != userId {
		renderError(w, http.StatusForbidden, "you don't have access to this disc")
		return
	}

	renderJSON(w, http.StatusOK, d)
}
