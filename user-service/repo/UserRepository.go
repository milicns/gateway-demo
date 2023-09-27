package repo

import (
	"context"
	"user-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "users"
)

type UserRepository struct {
	users *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	usersCol := client.Database(DATABASE).Collection(COLLECTION)
	return &UserRepository{
		users: usersCol,
	}
}

func (repo *UserRepository) Create(user *model.User) error {
	_, err := repo.users.InsertOne(context.TODO(), user)
	return err
}

func (repo *UserRepository) Delete(name string) error {
	filter := bson.D{{Key: "name", Value: name}}

	_, err := repo.users.DeleteOne(context.TODO(), filter)
	return err
}

func (repo *UserRepository) GetOne(name string) (*model.User, error) {
	filter := bson.D{{Key: "name", Value: name}}

	var result model.User
	err := repo.users.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}
