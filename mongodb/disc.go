package mongodb

import (
	"context"

	"github.com/MrTj458/scorecard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ models.DiscService = (*DiscService)(nil)

type DiscService struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewDiscService(db *mongo.Database) *DiscService {
	return &DiscService{
		db:   db,
		coll: db.Collection("discs"),
	}
}

func (ds *DiscService) Add(disc models.DiscIn) (models.Disc, error) {
	d := models.Disc{
		ID:           primitive.NewObjectID(),
		CreatedBy:    disc.CreatedBy,
		Name:         disc.Name,
		Type:         disc.Type,
		Manufacturer: disc.Manufacturer,
		Plastic:      disc.Plastic,
		Weight:       disc.Weight,
		Speed:        disc.Speed,
		Glide:        disc.Glide,
		Turn:         disc.Turn,
		Fade:         disc.Fade,
		InBag:        disc.InBag,
	}

	_, err := ds.coll.InsertOne(context.Background(), d)
	if err != nil {
		return models.Disc{}, err
	}

	return d, nil
}

func (ds *DiscService) FindAllByUserId(id string) ([]models.Disc, error) {
	cur, err := ds.coll.Find(context.Background(), bson.D{{"createdby", id}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var discs []models.Disc
	err = cur.All(context.Background(), &discs)
	if err != nil {
		return nil, err
	}

	if len(discs) == 0 {
		discs = make([]models.Disc, 0)
	}

	return discs, nil
}

func (ds *DiscService) FindOneById(id string) (models.Disc, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Disc{}, err
	}

	var d models.Disc
	err = ds.coll.FindOne(context.Background(), bson.M{"_id": oID}).Decode(&d)
	if err != nil {
		return models.Disc{}, err
	}

	return d, nil
}
