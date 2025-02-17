package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CompetenceModel struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Type        string          `bson:"type" json:"type"`
	Points      []bson.ObjectID `bson:"points" json:"points"`
	TotalPoints float64         `bson:"totalPoints" json:"totalPoints"` //среднее арифметическое всех score в points конкретного типа
}
