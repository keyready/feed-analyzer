package usecases

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
)

type CandidateUsecase interface {
	GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel)
	GetOneCandidate(candidateId string) (httpCode int, usecaseErr error, candidate response.CandidateResponse)
	AssessmentCandidate(candidateData request.CandidateData) (httpCode int, usecaseErr error, qualification models.QualificationModel)
}

type CandidateUsecaseImpl struct {
	candidateRepo repositories.CandidateRepository
}

func NewCandidateUsecase(candidateRepo repositories.CandidateRepository) *CandidateUsecaseImpl {
	return &CandidateUsecaseImpl{candidateRepo: candidateRepo}
}

func (candUsecase *CandidateUsecaseImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, usecaseErr error, qualification models.QualificationModel) {

	httpCode, usecaseErr, qualification = candUsecase.candidateRepo.AssessmentCandidate(candidateData)
	if usecaseErr != nil {
		return httpCode, usecaseErr, qualification
	}

	return http.StatusOK, nil, qualification
}

func (candUsecase *CandidateUsecaseImpl) GetAllCandidates() (httpCode int, usecaseErr error, candidates []models.CandidateModel) {
	httpCode, usecaseErr, candidates = candUsecase.candidateRepo.GetAllCandidates()
	if usecaseErr != nil {
		return httpCode, usecaseErr, candidates
	}
	return httpCode, nil, candidates
}

func (candUsecase *CandidateUsecaseImpl) GetOneCandidate(candidateId string) (
	httpCode int, usecaseErr error, candidate response.CandidateResponse) {

	candidateIdPrimitive, _ := primitive.ObjectIDFromHex(candidateId)

	httpCode, usecaseErr, candidate = candUsecase.candidateRepo.GetOneCandidate(candidateIdPrimitive)
	if usecaseErr != nil {
		return httpCode, usecaseErr, candidate
	}

	return httpCode, nil, candidate
}
