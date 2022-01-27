package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MrTj458/scorecard/controllers"
	"github.com/MrTj458/scorecard/db"
	"github.com/MrTj458/scorecard/middleware"
	"github.com/MrTj458/scorecard/models"
	"github.com/MrTj458/scorecard/views"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	dbUrl := os.Getenv("DB_URL")
	if len(dbUrl) == 0 {
		log.Fatal("DB_URL environment variable must be set")
	}
	splitUrl := strings.Split(strings.Split(dbUrl, "?")[0], "/")
	dbName := splitUrl[len(splitUrl)-1]

	db := db.Connect(dbUrl, dbName)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.User)

	usersController := controllers.NewUsers(models.NewUserStore(db))
	scorecardsController := controllers.NewScorecards(models.NewScorecardStore(db))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", usersController.Routes())
		r.Mount("/scorecards", scorecardsController.Routes())
	})

	// Static files should be served from the build directory
	r.Handle("/*", http.FileServer(http.Dir("static")))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		views.Error(w, http.StatusNotFound, "route not found")
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		views.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, r)
}
