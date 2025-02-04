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
	GetOneBodyCompetence(t, name string) (bodyCompetence models.BodyCompetenceModel)
}

type CandidateRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCandidateRepository(mongoDB *mongo.Database) *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{mongoDB: mongoDB}
}

func (candRepo *CandidateRepositoryImpl) GetOneBodyCompetence(t, name string) (bodyCompetence models.BodyCompetenceModel) {
	candRepo.mongoDB.Collection("body_competencies").
		FindOne(ctx, bson.M{"name": name, "type": t}).
		Decode(&bodyCompetence)
	return bodyCompetence
}

func (candRepo *CandidateRepositoryImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, repoErr error, candidate models.CandidateModel) {

	return http.StatusOK, nil, candidate

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

	return http.StatusOK, nil, candidate
}

func (candRepo *CandidateRepositoryImpl) GetAllCandidates() (httpCode int, repoErr error, candidates []models.CandidateModel) {
	cursor, mongoErr := candRepo.mongoDB.Collection("candidates").
		Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения записей из коллекции candidates: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, candidates
	}

	for cursor.Next(ctx) {
		var candidate models.CandidateModel
		if decodeErr := cursor.Decode(&candidate); decodeErr != nil {
			repoErr = fmt.Errorf("Ошибка анмаршалинга записи из коллекции candidates: %w", decodeErr.Error())
			return http.StatusInternalServerError, repoErr, candidates
		}
		candidates = append(candidates, candidate)
	}

	return http.StatusOK, nil, candidates
}
