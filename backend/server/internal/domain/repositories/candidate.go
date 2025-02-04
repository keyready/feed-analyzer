package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/models"
)

var (
	ctx context.Context
)

type CandidateRepository interface {
	GetAllCandidates() (httpCode int, repoErr error, candidates []models.CandidateModel)
	//GetOneCandidate() (httpCode int, repoErr error, candidate models.CandidateModel)
}

type CandidateRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCandidateRepository(mongoDB *mongo.Database) *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{mongoDB: mongoDB}
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
