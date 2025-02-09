package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PointResponse struct {
	Value float64 `bson:"value" json:"value"`
	Score float64 `bson:"score" json:"score"`
}

type CompetenceResponse struct {
	Type        string          `bson:"type" json:"type"`
	Points      []PointResponse `bson:"points" json:"points"`
	TotalPoints float64         `bson:"totalPoints" json:"totalPoints"`
}

type QualificationResponse struct {
	Value      string  `bson:"value" json:"value"`
	Assessment float64 `bson:"assessment" json:"assessment"`
	Notes      string  `bson:"notes" json:"notes"`
}

type CandidateResponse struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string             `bson:"firstname" json:"firstname"`
	Lastname   string             `bson:"lastname" json:"lastname"`
	Middlename string             `bson:"middlename" json:"middlename"`
	Age        int64              `bson:"age" json:"age"`
	Rank       string             `bson:"rank" json:"rank"`
	Avatar     string             `bson:"avatar" json:"avatar"`
	Documents  []string           `bson:"documents" json:"documents"`

	Qualification QualificationResponse `json:"qualification"`
	Competences   []CompetenceResponse  `json:"competences"`
}
