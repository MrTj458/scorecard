package models

import (
	"github.com/MrTj458/scorecard/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Disc struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreatedBy    string             `json:"created_by"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	Manufacturer string             `json:"manufacturer"`
	Plastic      string             `json:"plastic"`
	Weight       int                `json:"weight"`
	Speed        int                `json:"speed"`
	Glide        int                `json:"glide"`
	Turn         int                `json:"turn"`
	Fade         int                `json:"fade"`
	InBag        bool               `json:"in_bag"`
}

type DiscIn struct {
	CreatedBy    string `json:"created_by" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Plastic      string `json:"plastic"`
	Weight       int    `json:"weight"`
	Speed        int    `json:"speed"`
	Glide        int    `json:"glide"`
	Turn         int    `json:"turn"`
	Fade         int    `json:"fade"`
	InBag        bool   `json:"in_bag"`
}

type DiscStore struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewDiscStore(db *mongo.Database) *DiscStore {
	return &DiscStore{
		db:   db,
		coll: db.Collection("discs"),
	}
}

func (ds *DiscStore) Add(disc DiscIn) (Disc, error) {
	d := Disc{
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

	_, err := ds.coll.InsertOne(db.Ctx, d)
	if err != nil {
		return Disc{}, err
	}

	return d, nil
}

func (ds *DiscStore) FindAllByUserId(id string) ([]Disc, error) {
	cur, err := ds.coll.Find(db.Ctx, bson.D{{"createdby", id}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var discs []Disc
	err = cur.All(db.Ctx, &discs)
	if err != nil {
		return nil, err
	}

	if len(discs) == 0 {
		discs = make([]Disc, 0)
	}

	return discs, nil
}
