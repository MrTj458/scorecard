package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/MrTj458/scorecard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ models.ScorecardService = (*ScorecardService)(nil)

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

func (ss *ScorecardService) Add(scorecard models.ScorecardIn) (models.Scorecard, error) {
	s := models.Scorecard{
		ID:          primitive.NewObjectID(),
		CreatedBy:   scorecard.CreatedBy,
		StartTime:   time.Now().UTC(),
		EndTime:     nil,
		CourseName:  scorecard.CourseName,
		CourseState: scorecard.CourseState,
		Players:     scorecard.Players,
		Holes:       make([]models.Hole, 0),
	}

	_, err := ss.coll.InsertOne(context.Background(), s)
	if err != nil {
		return models.Scorecard{}, err
	}

	return s, nil
}

func (ss *ScorecardService) FindAll() ([]models.Scorecard, error) {
	cur, err := ss.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var scorecards []models.Scorecard
	if err = cur.All(context.Background(), &scorecards); err != nil {
		return nil, err
	}

	if scorecards == nil {
		scorecards = make([]models.Scorecard, 0)
	}

	return scorecards, nil
}

func (ss *ScorecardService) FindByID(id string) (models.Scorecard, error) {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Scorecard{}, err
	}

	var s models.Scorecard
	err = ss.coll.FindOne(context.Background(), bson.D{{"_id", oId}}).Decode(&s)
	if err != nil {
		return models.Scorecard{}, err
	}

	return s, nil
}

func (ss *ScorecardService) FindAllByUserId(userId string) ([]models.Scorecard, error) {
	uId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{"players.id", uId},
	}

	cur, err := ss.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var scorecards []models.Scorecard
	if err = cur.All(context.Background(), &scorecards); err != nil {
		return nil, err
	}

	if scorecards == nil {
		scorecards = make([]models.Scorecard, 0)
	}

	return scorecards, nil
}

func (ss *ScorecardService) AddHole(scorecardId string, hole models.Hole) (models.Scorecard, error) {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return models.Scorecard{}, err
	}

	_, err = ss.coll.UpdateByID(context.Background(), cardId, bson.D{{"$push", bson.D{{"holes", hole}}}})
	if err != nil {
		return models.Scorecard{}, err
	}

	return ss.FindByID(scorecardId)
}

func (ss *ScorecardService) Complete(scorecardId string) (models.Scorecard, error) {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return models.Scorecard{}, err
	}

	_, err = ss.coll.UpdateByID(context.Background(), cardId, bson.D{{"$set", bson.D{{"end_time", time.Now().UTC()}}}})
	if err != nil {
		fmt.Println(err.Error())
		return models.Scorecard{}, err
	}

	return ss.FindByID(scorecardId)
}

func (ss *ScorecardService) Delete(scorecardId string) error {
	cardId, err := primitive.ObjectIDFromHex(scorecardId)
	if err != nil {
		return err
	}

	_, err = ss.coll.DeleteOne(context.Background(), bson.D{{"_id", cardId}})
	if err != nil {
		return err
	}

	return nil
}
