package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Add(user UserIn) (User, error)
	FindAll() ([]User, error)
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	FindExistingUsers(email, username string) ([]User, error)
	SearchByUsername(username string) ([]User, error)
}

var (
	ErrEmailInUse    = errors.New("email already in use")
	ErrUsernameInUse = errors.New("username already in use")
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Username  string             `json:"username" bson:"username"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Password  string             `json:"-" bson:"password"`
}

type UserIn struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required,max=30"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
