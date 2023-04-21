package mongo

import (
	"context"
	"log"

	"github.com/laridtm/minha_saude/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReminderRepository struct {
	collection *mongo.Collection
}

func NewReminderRepository(database *mongo.Database) ReminderRepository {
	return ReminderRepository{
		collection: database.Collection("reminders"),
	}
}

func (rr *ReminderRepository) Insert(reminder model.Reminder) error {
	reminder.ID = primitive.NewObjectID().Hex()

	result, err := rr.collection.InsertOne(context.Background(), reminder)

	log.Printf("Reminder inserted with id %s", result.InsertedID)

	return err
}

func (rr *ReminderRepository) Update(reminderId string, reminder model.Reminder) error {
	filter := bson.M{"_id": reminderId}
	toUpdate := bson.M{"$set": reminder}

	_, err := rr.collection.UpdateOne(context.Background(), filter, toUpdate)

	log.Printf("Updated reminder with id %s", reminderId)

	return err
}

func (rr *ReminderRepository) FindAll(userId string) ([]model.Reminder, error) {
	where := bson.M{"userid": userId}

	cursor, err := rr.collection.Find(context.Background(), where)
	if err != nil {
		return nil, err
	}

	results := make([]model.Reminder, 0)

	cursor.All(context.Background(), &results)

	return results, nil
}

func (rr *ReminderRepository) Delete(reminderId string) error {
	filter := bson.M{"_id": reminderId}

	_, err := rr.collection.DeleteOne(context.Background(), filter)

	log.Printf("Deleted reminder with id %s", reminderId)

	return err
}
