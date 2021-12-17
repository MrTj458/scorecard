package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Password string             `json:"-"`
}

type UserService struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		db:   db,
		coll: db.Collection("users"),
	}
}
