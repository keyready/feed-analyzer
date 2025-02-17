package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CandidateModel struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Firstname  string `bson:"firstname" json:"firstname"`
	Lastname   string `bson:"lastname" json:"lastname"`
	Middlename string `bson:"middlename" json:"middlename"`

	Age    int64  `bson:"age" json:"age"`
	Rank   string `bson:"rank" json:"rank"`
	Avatar string `bson:"avatar" json:"avatar"`

	Documents []string `bson:"documents" json:"documents"`

	Competencies  []bson.ObjectID `bson:"competencies" json:"competencies"`
	Qualification bson.ObjectID   `bson:"qualification"`
}
