package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScorecardService interface {
	Add(scorecard ScorecardIn) (Scorecard, error)
	FindAll() ([]Scorecard, error)
	FindByID(id string) (Scorecard, error)
	FindAllByUserId(userId string) ([]Scorecard, error)
	AddHole(scorecardId string, hole Hole) (Scorecard, error)
	Complete(scorecardId string) (Scorecard, error)
	Delete(scorecardId string) error
}

type Scorecard struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
	StartTime   time.Time          `json:"start_time" bson:"start_time"`
	EndTime     *time.Time         `json:"end_time" bson:"end_time"`
	CourseName  string             `json:"course_name" bson:"course_name"`
	CourseState string             `json:"course_state" bson:"course_state"`
	Players     []Player           `json:"players" bson:"players"`
	Holes       []Hole             `json:"holes" bson:"holes"`
}

type Player struct {
	Username string             `json:"username" validate:"required" bson:"username"`
	ID       primitive.ObjectID `json:"id" validate:"required" bson:"id"`
}

type Hole struct {
	Number   int     `json:"number" validate:"required" bson:"number"`
	Par      int     `json:"par" validate:"required" bson:"par"`
	Distance int     `json:"distance" validate:"required" bson:"distance"`
	Scores   []Score `json:"scores" validate:"required" bson:"scores"`
}

type Score struct {
	Username string             `json:"username" validate:"required" bson:"username"`
	ID       primitive.ObjectID `json:"id" validate:"required" bson:"id"`
	Strokes  int                `json:"strokes" validate:"required" bson:"strokes"`
}

type ScorecardIn struct {
	CreatedBy   primitive.ObjectID `json:"created_by" validate:"required"`
	CourseName  string             `json:"course_name" validate:"required,max=50"`
	CourseState string             `json:"course_state" validate:"required,max=2"`
	Players     []Player           `json:"players" validate:"required"`
}
