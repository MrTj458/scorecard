package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/MrTj458/scorecard/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

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

type UserStore struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewUserStore(db *mongo.Database) *UserStore {
	return &UserStore{
		db:   db,
		coll: db.Collection("users"),
	}
}

// Add creates a new User in the database and returns the newly generated ID
func (us *UserStore) Add(user UserIn) (User, error) {
	// Check if username or email is already in use
	existingUsers, err := us.findExistingUsers(user.Email, user.Username)
	if err != nil {
		return User{}, err
	}

	if len(existingUsers) > 0 {
		for _, u := range existingUsers {
			if u.Email == user.Email {
				return User{}, ErrEmailInUse
			}

			if u.Username == user.Username {
				return User{}, ErrUsernameInUse
			}
		}
	}

	// Create new user
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	u := User{
		ID:        primitive.NewObjectID(),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: time.Now().UTC(),
		Password:  string(hashed),
	}

	_, err = us.coll.InsertOne(db.Ctx, u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

// FindAll returns all users found in the database
func (us *UserStore) FindAll() ([]User, error) {
	cur, err := us.coll.Find(db.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var users []User
	if err = cur.All(db.Ctx, &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]User, 0)
	}

	return users, nil
}

// FindById returns all users in the database with the given id
func (us *UserStore) FindByID(id string) (User, error) {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	var u User
	err = us.coll.FindOne(db.Ctx, bson.D{{"_id", oId}}).Decode(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

// FindByEmail returns all users in the database with the given email
func (us *UserStore) FindByEmail(email string) (User, error) {
	var u User
	err := us.coll.FindOne(db.Ctx, bson.D{{"email", email}}).Decode(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

// FindExistingUsers returns all the users in the database that have either
// the given email or username
func (us *UserStore) findExistingUsers(email, username string) ([]User, error) {
	filter := bson.D{
		{"$or", bson.A{bson.D{{"email", email}}, bson.D{{"username", username}}}},
	}

	cur, err := us.coll.Find(db.Ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var users []User
	if err = cur.All(db.Ctx, &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]User, 0)
	}

	return users, nil
}

func (us *UserStore) SearchByUsername(username string) ([]User, error) {
	filter := bson.D{
		{"username", username},
	}

	cur, err := us.coll.Find(db.Ctx, filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cur.Close(db.Ctx)

	var users []User
	if err = cur.All(db.Ctx, &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]User, 0)
	}

	return users, nil
}
