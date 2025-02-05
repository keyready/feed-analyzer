package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BodyCompetenceModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`

	Weight float64 `bson:"weight" json:"weight"` //программирование - 0.9
}
