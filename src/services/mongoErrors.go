package services

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoErrors struct {}

func (*MongoErrors) handleError(err error) error {
	message := err.Error()

	if (mongo.IsDuplicateKeyError(err)) {
		return errors.New("duplicated key")
	} else if (strings.Contains(message, "failed validation")) {
		return errors.New("failed validation")
	}

	return err
}