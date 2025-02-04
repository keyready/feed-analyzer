package request

import (
	"mime/multipart"
	"server/internal/domain/types/dto"
)

type CandidateData struct {
	Firstname  string `bson:"firstname" json:"firstname"`
	Middlename string `bson:"middlename" json:"middlename"`
	Lastname   string `bson:"lastname" json:"lastname"`

	Age        int             `bson:"age" json:"age"`
	Rank       string          `bson:"rank" json:"rank"`
	Avatar     string          `bson:"avatar"`
	AvatarFile *multipart.File `json:"avatar"`

	DocumentsNames []string                      `bson:"documentsName"`
	Documents      []*multipart.File             `json:"documents"`
	Competences    []dto.CandidateCompetenceData `json:"competences"`
}

type GetOneCandidateRequest struct {
	ID string `json:"id"`
}

type FilteredCandidatesRequest struct {
}
