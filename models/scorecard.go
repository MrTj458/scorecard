package models

import (
	"time"

	"github.com/MrTj458/scorecard/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Scorecard struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedBy   primitive.ObjectID `json:"created_by" bson:"created_by"`
	StartTime   time.Time          `json:"start_time" bson:"start_time"`
	EndTime     *time.Time         `json:"end_time" bson:"end_time"`
	NumHoles    int                `json:"num_holes" bson:"num_holes"`
	CourseName  string             `json:"course_name" bson:"course_name"`
	CourseState string             `json:"course_state" bson:"course_state"`
	Players     []player           `json:"players" bson:"players"`
	Holes       []hole             `json:"holes" bson:"holes"`
}

type player struct {
	Username string             `json:"username" validate:"required" bson:"username"`
	ID       primitive.ObjectID `json:"id" validate:"required" bson:"id"`
}

type hole struct {
	Hole     int     `json:"hole" validate:"required" bson:"hole"`
	Par      int     `json:"par" validate:"required" bson:"par"`
	Distance int     `json:"distance" validate:"required" bson:"distance"`
	Scores   []score `json:"scores" validate:"required" bson:"scores"`
}

type score struct {
	Username string             `json:"username" validate:"required" bson:"username"`
	ID       primitive.ObjectID `json:"id" validate:"required" bson:"id"`
	Strokes  int                `json:"strokes" validate:"required" bson:"strokes"`
}

type ScorecardIn struct {
	CreatedBy   primitive.ObjectID `json:"created_by" validate:"required"`
	NumHoles    int                `json:"num_holes" validate:"required"`
	CourseName  string             `json:"course_name" validate:"required"`
	CourseState string             `json:"course_state" validate:"required"`
	Players     []player           `json:"players" validate:"required"`
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

func (ss *ScorecardService) Add(scorecard ScorecardIn) (string, error) {
	s := Scorecard{
		ID:          primitive.NewObjectID(),
		CreatedBy:   scorecard.CreatedBy,
		StartTime:   time.Now().UTC(),
		EndTime:     nil,
		NumHoles:    scorecard.NumHoles,
		CourseName:  scorecard.CourseName,
		CourseState: scorecard.CourseState,
		Players:     scorecard.Players,
		Holes:       make([]hole, 0),
	}

	_, err := ss.coll.InsertOne(db.Ctx, s)
	if err != nil {
		return "", err
	}

	return s.ID.Hex(), nil
}

func (ss *ScorecardService) FindAll() ([]Scorecard, error) {
	cur, err := ss.coll.Find(db.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var scorecards []Scorecard
	if err = cur.All(db.Ctx, &scorecards); err != nil {
		return nil, err
	}

	if scorecards == nil {
		scorecards = make([]Scorecard, 0)
	}

	return scorecards, nil
}

func (ss *ScorecardService) FindByID(id string) (Scorecard, error) {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Scorecard{}, err
	}

	var s Scorecard
	err = ss.coll.FindOne(db.Ctx, bson.D{{"_id", oId}}).Decode(&s)
	if err != nil {
		return Scorecard{}, err
	}

	return s, nil
}

func (ss *ScorecardService) FindAllByUserId(userId string) ([]Scorecard, error) {
	uId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"players.id", uId},
	}

	cur, err := ss.coll.Find(db.Ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var scorecards []Scorecard
	if err = cur.All(db.Ctx, &scorecards); err != nil {
		return nil, err
	}

	if scorecards == nil {
		scorecards = make([]Scorecard, 0)
	}

	return scorecards, nil
}
