package repositories

import (
	"context"
	"errors"

	"my-go-server/model/users"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection 
}

var _ UserRepository= (*MongoUserRepository)(nil)

func NewMongoUserRepository(client *mongo.Client, dbName string) *MongoUserRepository {
	collection := client.Database(dbName).Collection("users")
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) Insert(user users.User) (users.User, error) {
	_, err := r.collection.InsertOne(context.Background(), user)
	return user, err
}

func (r *MongoUserRepository) GetByID(id string) (users.User, error) {
	var user users.User
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (r *MongoUserRepository) Update(user users.User) (users.User, error) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return user, err
}

func (r *MongoUserRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoUserRepository) GetByEmail(email string) (users.User, error) {
	var user users.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}
