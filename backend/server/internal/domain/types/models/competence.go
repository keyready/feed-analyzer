package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompetenceModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Body primitive.ObjectID `bson:"body" json:"body"`

	Value float64 `bson:"value" json:"value"` //есть среднее арифметическое всех body-компетишинов
	Score float64 `bson:"score" json:"score"` //есть произведение value на
}
