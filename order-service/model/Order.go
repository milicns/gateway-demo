package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	User string
}
