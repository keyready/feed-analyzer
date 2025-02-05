package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PointModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Value float64 `bson:"value" json:"value"` //программирование - 90
	Score float64 `bson:"score" json:"score"` //value * weight
}
