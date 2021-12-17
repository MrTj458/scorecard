package models

import (
	"log"
	"time"

	"github.com/MrTj458/scorecard/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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

// Add creates a new User in the database and returns the newly generated ID
func (us *UserService) Add(user UserIn) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
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
		return "", err
	}

	return u.ID.Hex(), nil
}

// FindAll returns all users found in the database
func (us *UserService) FindAll() ([]User, error) {
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
func (us *UserService) FindByID(id string) (User, error) {
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
func (us *UserService) FindByEmail(email string) (User, error) {
	var u User
	err := us.coll.FindOne(db.Ctx, bson.D{{"email", email}}).Decode(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

// FindExistingUsers returns all the users in the database that have either
// the given email or username
func (us *UserService) FindExistingUsers(email, username string) ([]User, error) {
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
