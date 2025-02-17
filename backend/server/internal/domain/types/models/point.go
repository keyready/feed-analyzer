package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PointModel struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Value float64 `bson:"value" json:"value"` //программирование - 90
	Score float64 `bson:"score" json:"score"` //value * weight
}
