package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Key struct {
	ID        primitive.ObjectID `bson:"_id"`
	KeyID     string             `bson:"key_id"`
	Key       string             `bson:"key"`
	Directory string             `bson:"directory"`
	Filename  string             `bson:"filename"`
}
