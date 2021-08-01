package services

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	Users string
}

type MongoService struct {
	client *mongo.Client
	Collections Collections
}


func (m *MongoService) Connect() {
	DB_URI := os.Getenv("DB_URI")
	m.Collections = Collections{Users: "users"}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	
	clientOptions := options.Client().ApplyURI(DB_URI)
	client, error := mongo.Connect(ctx, clientOptions)
	
	if error != nil {
		fmt.Println("error: ", error)
		panic(error)
	}
	
	m.client = client

	defer cancel()
}

func (m *MongoService) GetCollection(collectionName string) (collection *mongo.Collection) {
	DB_NAME := os.Getenv("DB_NAME")
	collection = m.client.Database(DB_NAME).Collection(collectionName)
	return
}

// Create the collections that we need
func (m *MongoService) CreateCollections() {
	go m.createUsersCollection()
}

// Create the users collection
func (m *MongoService) createUsersCollection() {
	DB_NAME := os.Getenv("DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	schema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "email", "password"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType": "string",
				"description": "User's name",
			},
			"email": bson.M{
				"bsonType": "string",
				"description": "User's email",
			},
			"password": bson.M{
				"bsonType": "string",
				"description": "User's password",
			},
		},
	}

	validator := bson.M{
		"$jsonSchema": schema,
	}
	createColOpts := options.CreateCollection().SetValidator(validator)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	collectionError := m.client.Database(DB_NAME).CreateCollection(ctx, m.Collections.Users, createColOpts)
	
	if collectionError != nil {
		if !strings.Contains(collectionError.Error(), "NamespaceExists") {
			fmt.Println("Error: ", collectionError.Error())
			panic("Error creating the users collection")
		}
	}
	
	_, indexError := m.client.Database(DB_NAME).Collection(m.Collections.Users).Indexes().CreateOne(ctx, indexModel)

	if indexError != nil {
		fmt.Println(indexError)
		panic("Error creating users index")
	}

	defer cancel()
}