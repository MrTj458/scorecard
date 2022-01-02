package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MrTj458/scorecard/controllers"
	"github.com/MrTj458/scorecard/db"
	"github.com/MrTj458/scorecard/middleware"
	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file, continuing anyway")
	}

	db := db.Connect(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.User)

	usersController := controllers.NewUsers(models.NewUserStore(db))
	scorecardsController := controllers.NewScorecards(models.NewScorecardStore(db))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", usersController.Routes())
		r.Mount("/scorecards", scorecardsController.Routes())
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		views.Error(w, http.StatusNotFound, "route not found")
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		views.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	log.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
