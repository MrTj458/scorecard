package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Scorecard struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedBy   primitive.ObjectID `json:"created_by"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	NumHoles    int                `json:"num_holes"`
	CourseName  string             `json:"course_name"`
	CourseState string             `json:"course_state"`
	Players     []struct {
		Username string             `json:"username"`
		ID       primitive.ObjectID `json:"id"`
	} `json:"players"`
	Holes []struct {
		Hole     int `json:"hole"`
		Par      int `json:"par"`
		Distance int `json:"distance"`
		Scores   []struct {
			Username string             `json:"username"`
			ID       primitive.ObjectID `json:"id"`
			Strokes  int                `json:"strokes"`
		} `json:"scores"`
	} `json:"holes"`
}

type ScorecardService struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewScorecardService(db *mongo.Database) *ScorecardService {
	return &ScorecardService{
		db:   db,
		coll: db.Collection("scorecards"),
	}
}
