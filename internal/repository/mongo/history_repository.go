package mongo

import (
	"context"
	"log"

	"github.com/laridtm/minha_saude/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HistoryRepository struct {
	collection *mongo.Collection
}

func NewHistoryRepository(database *mongo.Database) HistoryRepository {
	return HistoryRepository{
		collection: database.Collection("records"),
	}
}

func (hr *HistoryRepository) Insert(record model.MedicalRecord) error {
	record.ID = primitive.NewObjectID().Hex()

	result, err := hr.collection.InsertOne(context.Background(), record)

	log.Printf("Record inserted with id %s", result.InsertedID)

	return err
}

func (hr *HistoryRepository) Update(recordId string, record model.MedicalRecord) error {
	filter := bson.M{"_id": recordId}
	toUpdate := bson.M{"$set": record}

	_, err := hr.collection.UpdateOne(context.Background(), filter, toUpdate)

	log.Printf("Updated record with id %s", recordId)

	return err
}

func (hr *HistoryRepository) FindAll(userId string, filter *model.MedicalRecordType) ([]model.MedicalRecord, error) {
	where := bson.M{"userId": userId}

	if filter != nil {
		where["type"] = string(*filter)
	}

	cursor, err := hr.collection.Find(context.Background(), where)
	if err != nil {
		return nil, err
	}

	results := make([]model.MedicalRecord, 0)

	cursor.All(context.Background(), &results)

	return results, nil
}

func (hr *HistoryRepository) Delete(recordId string) error {
	filter := bson.M{"_id": recordId}

	_, err := hr.collection.DeleteOne(context.Background(), filter)

	log.Printf("Deleted record with id %s", recordId)

	return err
}
