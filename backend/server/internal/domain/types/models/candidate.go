package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CandidateModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Firstname  string `bson:"firstname" json:"firstname"`
	Lastname   string `bson:"lastname" json:"lastname"`
	Middlename string `bson:"middlename" json:"middlename"`

	Age    int64  `bson:"age" json:"age"`
	Rank   string `bson:"rank" json:"rank"`
	Avatar string `bson:"avatar" json:"avatar"`

	Documents []string `bson:"documents" json:"documents"`

	Competencies  []primitive.ObjectID `bson:"competencies" json:"competencies"`
	Qualification primitive.ObjectID   `bson:"qualification" json:"qualification"`
}
