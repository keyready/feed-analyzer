package usecases

import (
	"server/internal/domain/repositories"
	"server/internal/domain/types/models"
)

type CandidateUsecase interface {
	GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel)
}

type CandidateUsecaseImpl struct {
	candidateRepo repositories.CandidateRepository
}

func NewCandidateUsecase(candidateRepo repositories.CandidateRepository) *CandidateUsecaseImpl {
	return &CandidateUsecaseImpl{candidateRepo: candidateRepo}
}

func (candUsecase *CandidateUsecaseImpl) GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel) {
	httpCode, usecaseErr, candidates = candUsecase.candidateRepo.GetAllCandidates()
	if usecaseErr != nil {
		return httpCode, usecaseErr, candidates
	}
	return httpCode, nil, candidates
}

//func (candUsecase *CandidateUsecaseImpl)
