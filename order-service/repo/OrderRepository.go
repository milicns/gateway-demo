package repo

import (
	"context"
	"order-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "order"
	COLLECTION = "orders"
)

type OrderRepository struct {
	orders *mongo.Collection
}

func NewOrderRepository(client *mongo.Client) *OrderRepository {
	ordersCol := client.Database(DATABASE).Collection(COLLECTION)
	return &OrderRepository{
		orders: ordersCol,
	}
}

func (repo *OrderRepository) Create(order *model.Order) error {
	_, err := repo.orders.InsertOne(context.TODO(), order)
	return err
}

func (repo *OrderRepository) Delete(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.orders.DeleteOne(context.TODO(), filter)
	return err
}

func (repo *OrderRepository) GetOne(user string) (*model.Order, error) {
	filter := bson.D{{Key: "user", Value: user}}

	var result model.Order
	err := repo.orders.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}
