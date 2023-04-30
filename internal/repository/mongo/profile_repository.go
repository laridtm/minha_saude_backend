package mongo

import (
	"context"
	"errors"
	"log"

	"github.com/laridtm/minha_saude/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository struct {
	collection *mongo.Collection
}

func NewProfileRepository(database *mongo.Database) ProfileRepository {
	return ProfileRepository{
		collection: database.Collection("profile"),
	}
}

func (pr *ProfileRepository) Insert(profile model.Profile) error {

	result, err := pr.collection.InsertOne(context.Background(), profile)

	log.Printf("Profile inserted with id %s", result.InsertedID)

	return err
}

func (pr *ProfileRepository) Update(profile model.Profile) error {
	filter := bson.M{"cpf": profile.CPF}
	toUpdate := bson.M{"$set": profile}

	_, err := pr.collection.UpdateOne(context.Background(), filter, toUpdate)
	log.Printf("Updated profile with cpf %s", profile.CPF)

	return err
}

func (pr *ProfileRepository) Find(CPF string) (*model.Profile, error) {
	var profile model.Profile

	filter := bson.M{"cpf": CPF}

	cursor, err := pr.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return nil, errors.New("profile does not exist")
	}

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&profile); err != nil {
			return nil, err
		}
	}

	return &profile, nil
}
