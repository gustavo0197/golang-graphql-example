package services

import (
	"fmt"

	"context"
	"time"

	"github.com/gustavo0197/graphql/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name string `bson:"name"`
	Password string `bson:"password"`
	Email string `bson:"email"`
}

type UserService struct {
	MongoService
	MongoErrors
}

func (u *UserService) CreateUser(data model.NewUser) (*model.User, error) {
	user := &model.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		fmt.Println("Error hashing password: ", err)
		return nil, err
	}

	document := bson.M{"name": data.Name, "password": string(password), "email": data.Email}
	documentId, err := u.GetCollection("users").InsertOne(ctx, document)

	if err != nil {
		fmt.Println("User service error: ", err.Error())
		return nil, u.MongoErrors.handleError(err)
	}

	err = u.GetCollection("users").FindOne(ctx, bson.D{{"_id", documentId.InsertedID}}).Decode(&user)

	if err != nil {
		fmt.Println("Error fetching user: ", err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUser(id string) (*model.User, error) {
	user := &model.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = u.GetCollection(u.MongoService.Collections.Users).FindOne(ctx, bson.D{{"_id", userId}}).Decode(&user)

	if err != nil {
		fmt.Println("Get user error: ", err)
		return nil, err
	}

	return user, nil
}