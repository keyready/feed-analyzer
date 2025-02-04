package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllBodyCompetence struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
