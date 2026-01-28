package repositories

import (
	"context"
	"my-go-server/model/users"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRefreshTokenRepository struct {
	collection *mongo.Collection
}

var _ RefreshTokenRepository = (*MongoRefreshTokenRepository)(nil)


func NewMongoRefreshTokenRepository(client *mongo.Client, dbName string) *MongoRefreshTokenRepository {
	collection := client.Database(dbName).Collection("refresh_tokens")
	return &MongoRefreshTokenRepository{collection: collection}
}

func (r *MongoRefreshTokenRepository) Insert(token users.RefreshToken) (users.RefreshToken, error) {
	_, err := r.collection.InsertOne(context.Background(), token)
	return token, err
}

func (r *MongoRefreshTokenRepository) GetByID(id string) (users.RefreshToken, error) {
	var token users.RefreshToken
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&token)
	return token, err
}

func (r *MongoRefreshTokenRepository) GetByToken(token string) (users.RefreshToken, error) {
	var storedToken users.RefreshToken
	filter := bson.M{"token_hash": token}
	err := r.collection.FindOne(context.Background(), filter).Decode(&storedToken)
	return storedToken, err
}

func (r *MongoRefreshTokenRepository) Update(token users.RefreshToken) (users.RefreshToken, error) {
	filter := bson.M{"_id": token.ID}
	update := bson.M{"$set": token}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return token, err
}

func (r *MongoRefreshTokenRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

