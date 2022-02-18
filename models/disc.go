package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiscService interface {
	Add(disc DiscIn) (Disc, error)
	FindAllByUserId(id string) ([]Disc, error)
	FindOneById(id string) (Disc, error)
}

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
	Name         string `json:"name" validate:"required,max=30"`
	Type         string `json:"type" validate:"required,max=30"`
	Manufacturer string `json:"manufacturer" validate:"required,max=30"`
	Plastic      string `json:"plastic" validate:"max=30"`
	Weight       int    `json:"weight" validate:"min=100,max=200"`
	Speed        int    `json:"speed" validate:"min=1,max=14"`
	Glide        int    `json:"glide" validate:"min=1,max=10"`
	Turn         int    `json:"turn" validate:"min=-5,max=5"`
	Fade         int    `json:"fade" validate:"min=0,max=10"`
	InBag        bool   `json:"in_bag"`
}
