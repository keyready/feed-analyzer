package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QualificationModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Type       string  `bson:"type" json:"type"`
	TotalScore float64 `bson:"totalScore" json:"totalScore"`
}
