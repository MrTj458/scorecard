package controllers

import (
	"net/http"

	"github.com/MrTj458/scorecard/middleware"
	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
)

type Discs struct {
	store *models.DiscStore
}

func NewDiscs(service *models.DiscStore) *Discs {
	return &Discs{
		store: service,
	}
}

func (dc *Discs) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequireLogin)

	r.Post("/", dc.create)
	r.Get("/", dc.getAll)

	return r
}

func (dc *Discs) create(w http.ResponseWriter, r *http.Request) {
	// uId := r.Context().Value("user").(string)
	var disc models.DiscIn
	if err := views.DecodeJSON(r.Body, &disc); err != nil {
		views.Error(w, http.StatusBadRequest, "invalid JSON object received")
		return
	}

	if errors := views.Validate(disc); errors != nil {
		views.ErrorWithFields(w, http.StatusUnprocessableEntity, "invalid disc object received", errors)
		return
	}

	discs, err := dc.store.Add(disc)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "error fetching discs")
		return
	}

	views.JSON(w, http.StatusOK, discs)
}

func (dc *Discs) getAll(w http.ResponseWriter, r *http.Request) {
	uId := r.Context().Value("user").(string)

	discs, err := dc.store.FindAllByUserId(uId)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, "unable to fetch discs")
		return
	}

	views.JSON(w, http.StatusOK, discs)
}
