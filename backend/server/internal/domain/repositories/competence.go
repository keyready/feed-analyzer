package repositories

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
)

type CompetenceRepository interface {
	GetAllCompetencies() (httpCode int, repErr error, competencies []models.CompetenceModel)
	GetAllBodyCompetencies(allBodyCompetenceRequest request.AllBodyCompetenceRequest) (httpCode int, repErr error, allBodyCompetencies []response.AllBodyCompetenceResponse)
}

type CompetenceRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCompetenceRepository(mongoDB *mongo.Database) *CompetenceRepositoryImpl {
	return &CompetenceRepositoryImpl{mongoDB: mongoDB}
}

func (compR *CompetenceRepositoryImpl) GetAllBodyCompetencies(allBodyCompetenceRequest request.AllBodyCompetenceRequest) (
	httpCode int, repErr error, allBodyCompetencies []response.AllBodyCompetenceResponse) {
	if allBodyCompetenceRequest.Type == "" && allBodyCompetenceRequest.Name == "" {
		cursor, mongoErr := compR.mongoDB.Collection("body_competencies").
			Find(ctx, bson.M{})
		if mongoErr != nil {
			repErr = fmt.Errorf("Ошибка извлечения из коллекции body_competencies: %w", mongoErr.Error())
			return http.StatusInternalServerError, repErr, allBodyCompetencies
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var bodyCompetence response.AllBodyCompetenceResponse
			if decodeErr := cursor.Decode(&bodyCompetence); decodeErr != nil {
				repErr = fmt.Errorf("Ошибка анмаршлинга одной записи body_competencies: %w", decodeErr)
				return http.StatusInternalServerError, repErr, allBodyCompetencies
			}
			allBodyCompetencies = append(allBodyCompetencies, bodyCompetence)
		}
		return http.StatusOK, nil, allBodyCompetencies
	}

	filterOr := bson.M{"$or": []bson.M{{"type": allBodyCompetenceRequest.Type}, {"name": allBodyCompetenceRequest.Name}}}

	cursor, mongoErr := compR.mongoDB.Collection("body_competencies").
		Find(ctx, filterOr)
	if mongoErr != nil {
		repErr = fmt.Errorf("Ошибка извлечения из коллекции с body_competencies по type %w или name %w : %w",
			allBodyCompetenceRequest.Type, allBodyCompetenceRequest.Name, mongoErr.Error())
		return http.StatusInternalServerError, repErr, allBodyCompetencies
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bodyCompetence response.AllBodyCompetenceResponse
		if decodeErr := cursor.Decode(&bodyCompetence); decodeErr != nil {
			repErr = fmt.Errorf("Ошибка анмаршалинга записи из body_competencies: %w", decodeErr.Error())
			return http.StatusInternalServerError, repErr, allBodyCompetencies
		}
		allBodyCompetencies = append(allBodyCompetencies, bodyCompetence)
	}

	return http.StatusOK, nil, allBodyCompetencies
}

func (compR *CompetenceRepositoryImpl) GetAllCompetencies() (httpCode int, repErr error, competences []models.CompetenceModel) {
	cursor, mongoErr := compR.mongoDB.Collection("competencies").Find(ctx, bson.M{})
	if mongoErr != nil {
		repErr = fmt.Errorf("Ошибка извлечения всех записей из коллекции competences: %v", mongoErr.Error())
		return http.StatusInternalServerError, repErr, competences
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var competence models.CompetenceModel
		if decodeErr := cursor.Decode(&competence); decodeErr != nil {
			repErr = fmt.Errorf("Ошибка анмаршалинга одной записи компетенции: %v", decodeErr.Error())
			return http.StatusInternalServerError, repErr, competences
		}
		competences = append(competences, competence)
	}

	return http.StatusOK, nil, competences
}
