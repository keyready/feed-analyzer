package response

import (
	"server/internal/domain/types/models"
)

type CandidateWithQualification struct {
	Candidate          models.CandidateModel `bson:",inline"`
	QualificationValue string                `bson:"qualificationValue"`
	Assessment         float64               `bson:"assessment"`
}
