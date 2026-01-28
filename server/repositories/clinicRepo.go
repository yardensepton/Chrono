package repositories

import (
	"context"
	"errors"

	"my-go-server/model/clinics"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoClinicRepository struct {
	collection *mongo.Collection
}

var _ Repository[clinics.Clinic] = (*MongoClinicRepository)(nil)

func NewMongoClinicRepository(client *mongo.Client, dbName string) *MongoClinicRepository {
	collection := client.Database(dbName).Collection("clinics")
	return &MongoClinicRepository{collection: collection}
}

func (r *MongoClinicRepository) Insert(clinic clinics.Clinic) (clinics.Clinic, error) {
	_, err := r.collection.InsertOne(context.Background(), clinic)
	return clinic, err
}

func (r *MongoClinicRepository) GetByID(id string) (clinics.Clinic, error) {
	var clinic clinics.Clinic
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&clinic)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return clinic, errors.New("clinic not found")
		}
		return clinic, err
	}
	return clinic, nil
}

func (r *MongoClinicRepository) Update(clinic clinics.Clinic) (clinics.Clinic, error) {
	filter := bson.M{"_id": clinic.ID}
	update := bson.M{"$set": clinic}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return clinic, err
}

func (r *MongoClinicRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoClinicRepository) GetByEmail(email string) (clinics.Clinic, error) {
	// Clinics don't have emails, so this method doesn't make sense for clinics
	return clinics.Clinic{}, errors.New("clinics do not have email addresses")
}
