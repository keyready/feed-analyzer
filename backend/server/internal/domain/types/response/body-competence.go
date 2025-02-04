package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllBodyCompetenceResponse struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"type" bson:"type"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}
