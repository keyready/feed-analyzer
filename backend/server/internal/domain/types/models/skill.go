package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SkillModel struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Type string `bson:"type" json:"type"`

	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`

	Weight float64 `bson:"weight" json:"weight"` //программирование - 0.9
}
