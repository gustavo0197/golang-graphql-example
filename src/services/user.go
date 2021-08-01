package services

import (
	"fmt"

	"context"
	"time"

	"github.com/gustavo0197/graphql/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// TODO hash password

	document := bson.M{"name": data.Name, "password": data.Password, "email": data.Email}

	defer cancel()

	documentId, insertError := u.GetCollection("users").InsertOne(ctx, document)

	if insertError != nil {
		fmt.Println("User service error: ", insertError.Error())
		return nil, u.MongoErrors.handleError(insertError)
	}

	userError := u.GetCollection("users").FindOne(ctx, bson.D{{"_id", documentId.InsertedID}}).Decode(&user)

	if userError != nil {
		fmt.Println("Error fetching user: ", userError)
		return nil, userError
	}

	return user, nil
}

func (u *UserService) GetUser(id string) (*model.User, error) {
	user := &model.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	userId, userIdError := primitive.ObjectIDFromHex(id)

	if userIdError != nil {
		return nil, userIdError
	}

	userError := u.GetCollection(u.MongoService.Collections.Users).FindOne(ctx, bson.D{{"_id", userId}}).Decode(&user)

	if userError != nil {
		fmt.Println(userError)
		panic("Get user error")
	}

	return user, nil
}