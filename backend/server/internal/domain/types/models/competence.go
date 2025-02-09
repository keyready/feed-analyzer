package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompetenceModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Type        string               `bson:"type" json:"type"`
	Points      []primitive.ObjectID `bson:"points" json:"points"`
	TotalPoints float64              `bson:"totalPoints" json:"totalPoints"` //среднее арифметическое всех score в points конкретного типа
}
