package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QualificationModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Value      string  `bson:"value" json:"value"`           //1 из 5 квалов(от high до low)
	Assessment float64 `bson:"assessment" json:"assessment"` //оценка deepSeek
	Notes      string  `bson:"notes" json:"notes"`
}
