package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompetenceModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Type       string               `bson:"type" json:"type"`
	Body       []primitive.ObjectID `bson:"body" json:"body"`
	Points     []primitive.ObjectID `bson:"points" json:"points"`
	TotalScore float64              `bson:"totalScore" json:"totalScore"` //среднее арифметическое всех score в points конкретного типа
}
