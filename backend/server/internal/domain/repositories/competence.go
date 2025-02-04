package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/response"
)

var (
	ctx context.Context
)

type CompetenceRepository interface {
	GetAllCompetencies() (httpCode int, repErr error, competencies []models.CompetenceModel)
	GetAllBodyCompetencies() (httpCode int, repErr error, allBodyCompetencies []response.AllBodyCompetence)
	//GetOneCompetenceByType(t enum.TypeCompetence) (httpCode int, repErr error, competence models.CompetenceModel)
}

type CompetenceRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCompetenceRepository(mongoDB *mongo.Database) *CompetenceRepositoryImpl {
	return &CompetenceRepositoryImpl{mongoDB: mongoDB}
}

func (compR *CompetenceRepositoryImpl) GetAllBodyCompetencies(
	httpCode int,
	repErr error,
	allBodyCompetencies []response.AllBodyCompetence) {
	//cursor, mongoErr := compR.mongoDB.Collection("body_competencies").Find(ctx, bson.D{})
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
