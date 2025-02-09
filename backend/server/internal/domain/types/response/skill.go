package response

type Skill struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type AllSkills struct {
	Type   string  `json:"type" bson:"type"`
	Skills []Skill `json:"skills" bson:"skills"`
}
