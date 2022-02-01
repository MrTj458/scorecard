package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MrTj458/scorecard/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var _ models.UserService = (*UserService)(nil)

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

// Add creates a new User in the database and returns it.
func (us *UserService) Add(user models.UserIn) (models.User, error) {
	// Check if username or email is already in use
	existingUsers, err := us.FindExistingUsers(user.Email, user.Username)
	if err != nil {
		return models.User{}, err
	}

	if len(existingUsers) > 0 {
		for _, u := range existingUsers {
			if u.Email == user.Email {
				return models.User{}, models.ErrEmailInUse
			}

			if u.Username == user.Username {
				return models.User{}, models.ErrUsernameInUse
			}
		}
	}

	// Create new user
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	u := models.User{
		ID:        primitive.NewObjectID(),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: time.Now().UTC(),
		Password:  string(hashed),
	}

	_, err = us.coll.InsertOne(context.Background(), u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// FindAll returns all users found in the database.
func (us *UserService) FindAll() ([]models.User, error) {
	cur, err := us.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]models.User, 0)
	}

	return users, nil
}

// FindById returns all users in the database with the given id.
func (us *UserService) FindByID(id string) (models.User, error) {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	var u models.User
	err = us.coll.FindOne(context.Background(), bson.D{{"_id", oId}}).Decode(&u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// FindByEmail returns all users in the database with the given email.
func (us *UserService) FindByEmail(email string) (models.User, error) {
	var u models.User
	err := us.coll.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// FindExistingUsers returns all the users in the database that have either
// the given email or username.
func (us *UserService) FindExistingUsers(email, username string) ([]models.User, error) {
	filter := bson.D{
		{"$or", bson.A{bson.D{{"email", email}}, bson.D{{"username", username}}}},
	}

	cur, err := us.coll.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]models.User, 0)
	}

	return users, nil
}

func (us *UserService) SearchByUsername(username string) ([]models.User, error) {
	filter := bson.D{
		{"username", username},
	}

	cur, err := us.coll.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	if err = cur.All(context.Background(), &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]models.User, 0)
	}

	return users, nil
}
