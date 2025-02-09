package usecases

import (
	"net/http"
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
	return http.StatusOK, nil, candidate
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
