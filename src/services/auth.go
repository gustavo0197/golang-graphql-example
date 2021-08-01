package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gustavo0197/graphql/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	MongoService
}

func (a *AuthService) Login(credentials *model.UserCredentials) (token string, err error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	user := model.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{"email": credentials.Email}
	a.MongoService.GetCollection(a.MongoService.Collections.Users).FindOne(ctx, filter).Decode(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(credentials.Password))

	// Password is invalid
	if err != nil {
		return "", errors.New("invalid password")
	}

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	signedToken, err := authToken.SignedString([]byte(JWT_SECRET))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return signedToken, nil
}