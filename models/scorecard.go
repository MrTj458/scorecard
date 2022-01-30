package models

import (
	"fmt"
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

type ScorecardStore struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewScorecardStore(db *mongo.Database) *ScorecardStore {
	return &ScorecardStore{
		db:   db,
		coll: db.Collection("scorecards"),
	}
}

func (ss *ScorecardStore) Add(scorecard ScorecardIn) (Scorecard, error) {
	s := Scorecard{
		ID:          primitive.NewObjectID(),
		CreatedBy:   scorecard.CreatedBy,
		StartTime:   time.Now().UTC(),
		EndTime:     nil,
		CourseName:  scorecard.CourseName,
		CourseState: scorecard.CourseState,
		Players:     scorecard.Players,
		Holes:       make([]Hole, 0),
	}

	_, err := ss.coll.InsertOne(db.Ctx, s)
	if err != nil {
		return Scorecard{}, err
	}

	return s, nil
}

func (ss *ScorecardStore) FindAll() ([]Scorecard, error) {
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

func (ss *ScorecardStore) FindByID(id string) (Scorecard, error) {
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

func (ss *ScorecardStore) FindAllByUserId(userId string) ([]Scorecard, error) {
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

func (ss *ScorecardStore) AddHole(scorecardId string, hole Hole) (Scorecard, error) {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return Scorecard{}, err
	}

	_, err = ss.coll.UpdateByID(db.Ctx, cardId, bson.D{{"$push", bson.D{{"holes", hole}}}})
	if err != nil {
		return Scorecard{}, err
	}

	return ss.FindByID(scorecardId)
}

func (ss *ScorecardStore) Complete(scorecardId string) (Scorecard, error) {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return Scorecard{}, err
	}

	_, err = ss.coll.UpdateByID(db.Ctx, cardId, bson.D{{"$set", bson.D{{"end_time", time.Now().UTC()}}}})
	if err != nil {
		fmt.Println(err.Error())
		return Scorecard{}, err
	}

	return ss.FindByID(scorecardId)
}

func (ss *ScorecardStore) Delete(scorecardId string) error {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return err
	}

	_, err = ss.coll.DeleteOne(db.Ctx, bson.D{{"_id", cardId}})
	if err != nil {
		return err
	}

	return nil
}
