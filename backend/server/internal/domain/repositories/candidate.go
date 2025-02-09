package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
)

var (
	ctx context.Context
)

type CandidateRepository interface {
	GetAllCandidates() (httpCode int, repoErr error, candidates []models.CandidateModel)
	GetOneCandidate(getOneCandidateRequest request.GetOneCandidateRequest) (httpCode int, repoErr error, candidate models.CandidateModel)
	AssessmentCandidate(candidateData request.CandidateData) (httpCode int, repoErr error, candidate models.CandidateModel)
}

type CandidateRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCandidateRepository(mongoDB *mongo.Database) *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{mongoDB: mongoDB}
}

func (candRepo *CandidateRepositoryImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, repoErr error, candidate models.CandidateModel) {

	return http.StatusOK, nil, candidate

}

func (candRepo *CandidateRepositoryImpl) GetAllCandidates() (httpCode int, repoErr error, candidates []models.CandidateModel) {
	cursor, _ := candRepo.mongoDB.Collection("candidates").Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		candidate := models.CandidateModel{}
		cursor.Decode(&candidate)
		candidates = append(candidates, candidate)
	}

	return http.StatusOK, nil, candidates
}

func (candRepo *CandidateRepositoryImpl) GetOneCandidate(getOneCandidateRequest request.GetOneCandidateRequest) (
	httpCode int, repoErr error, candidate models.CandidateModel) {

	candidateId, primitiveErr := primitive.ObjectIDFromHex(getOneCandidateRequest.ID)
	if primitiveErr != nil {
		repoErr = fmt.Errorf("Ошибка декодирования candidateId: %w", primitiveErr.Error())
		return http.StatusInternalServerError, repoErr, candidate
	}

	mongoErr := candRepo.mongoDB.Collection("candidates").
		FindOne(ctx, bson.M{"_id": candidateId}).Decode(&candidate)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения одного кандидата: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, candidate
	}

	var qual models.QualificationModel
	candRepo.mongoDB.Collection("qualifications").FindOne(ctx, bson.M{"_id": candidate.Qualification}).Decode(&qual)

	var comp models.CandidateModel
	candRepo.mongoDB.Collection("competencies").FindOne(ctx, bson.M{"_id": candidate.Competencies[0]}).Decode(&comp)

	return http.StatusOK, nil, candidate
}
