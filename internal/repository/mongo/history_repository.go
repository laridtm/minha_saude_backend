package mongo

import (
	"context"
	"log"

	"github.com/laridtm/minha_saude/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (hr *HistoryRepository) FindAll(userId string, filter *model.Filter) ([]model.MedicalRecord, error) {
	opts := options.Find()

	if filter.Size != 0 {
		opts.SetLimit(int64(filter.Size))
	}

	where := bson.M{"userId": userId}

	for k, v := range filter.Fields {
		where[k] = v
	}

	if filter.FromDate != nil {
		where["date"] = bson.M{"$gte": filter.FromDate}
	}

	cursor, err := hr.collection.Find(context.Background(), where, opts)
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
