package main

import (
	"log"
	"os"

	"github.com/MrTj458/scorecard/api"
	"github.com/MrTj458/scorecard/mongodb"
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

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		log.Fatal("DB_NAME environment variable must be set")
	}

	db := mongodb.Connect(dbUrl, dbName)

	s := api.NewServer(port)

	s.UserService = mongodb.NewUserService(db)
	s.ScorecardService = mongodb.NewScorecardService(db)
	s.DiscService = mongodb.NewDiscService(db)

	log.Fatal(s.Run())
}
