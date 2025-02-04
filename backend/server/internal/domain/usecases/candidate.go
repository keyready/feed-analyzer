package usecases

import (
	"server/internal/domain/repositories"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
)

type CandidateUsecase interface {
	GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel)
	GetOneCandidate(getOneCandidate request.GetOneCandidateRequest) (httpCode int, usecaseErr error, candidate models.CandidateModel)
	AssessmentCandidate(candidateData request.CandidateData) (httpCode int, usecaseErr error, candidate models.CandidateModel)
}

type CandidateUsecaseImpl struct {
	candidateRepo repositories.CandidateRepository
}

func NewCandidateUsecase(candidateRepo repositories.CandidateRepository) *CandidateUsecaseImpl {
	return &CandidateUsecaseImpl{candidateRepo: candidateRepo}
}

func (candUsecase *CandidateUsecaseImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, usecaseErr error, candidate models.CandidateModel) {

	var weightedAssessment float64

	for _, competence := range candidateData.Competences {
		bodyCompetence := candUsecase.candidateRepo.GetOneBodyCompetence(
			competence.Type,
			competence.Name,
		)
		oneCompetenceScore := bodyCompetence.Weight * competence.Value
		//Запись скора по одному из типов компетенций
		weightedAssessment += oneCompetenceScore
	}

}

func (candUsecase *CandidateUsecaseImpl) GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel) {
	httpCode, usecaseErr, candidates = candUsecase.candidateRepo.GetAllCandidates()
	if usecaseErr != nil {
		return httpCode, usecaseErr, candidates
	}
	return httpCode, nil, candidates
}

func (candUsecase *CandidateUsecaseImpl) GetOneCandidate(getOneCandidate request.GetOneCandidateRequest) (httpCode int, usecaseErr error, candidate models.CandidateModel) {
	httpCode, usecaseErr, candidate = candUsecase.candidateRepo.GetOneCandidate(getOneCandidate)
	if usecaseErr != nil {
		return httpCode, usecaseErr, candidate
	}
	return httpCode, nil, candidate
}
