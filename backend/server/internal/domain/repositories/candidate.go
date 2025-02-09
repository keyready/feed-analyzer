package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/enum/competence"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
)

var (
	ctx context.Context
)

type CandidateRepository interface {
	GetAllCandidates() (httpCode int, repoErr error, candidates []models.CandidateModel)
	GetOneCandidate(candidateId primitive.ObjectID) (httpCode int, repoErr error, candidate response.CandidateResponse)
	AssessmentCandidate(candidateData request.CandidateData) (httpCode int, repoErr error, qualification models.QualificationModel)
}

type CandidateRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCandidateRepository(mongoDB *mongo.Database) *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{mongoDB: mongoDB}
}

func (candRepo *CandidateRepositoryImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, repoErr error, qualification models.QualificationModel) {

	types := []string{
		competence.AcademicAchievements,
		competence.SoftSkills,
		competence.TechnicalSkills,
		competence.MilitaryTraining,
		competence.DomainKnowledge,
	}

	for _, t := range types {
		skillsLen, _ := candRepo.mongoDB.Collection("skills").CountDocuments(ctx, bson.D{{"type", t}})
		var pointIds []primitive.ObjectID
		var totalValue float64
		for _, c := range candidateData.Competences {
			if c.Type == t {
				var skill models.SkillModel

				candRepo.mongoDB.Collection("skills").
					FindOne(ctx, bson.D{{"name", c.Name}, {"type", c.Type}}).Decode(&skill)

				score := skill.Weight * c.Value
				totalValue += c.Value

				point := models.PointModel{
					Score: score,
					Value: c.Value,
				}

				resInsert, _ := candRepo.mongoDB.Collection("points").InsertOne(ctx, &point)
				id := resInsert.InsertedID.(primitive.ObjectID)

				pointIds = append(pointIds, id)
			}
		}
		totalPoints := totalValue / float64(skillsLen)
	}

	newCandidate := models.CandidateModel{
		Firstname:  candidateData.Firstname,
		Lastname:   candidateData.Lastname,
		Middlename: candidateData.Middlename,
		Age:        candidateData.Age,
		Rank:       candidateData.Rank,
		Avatar:     candidateData.Avatar,
		Documents:  candidateData.DocumentsNames,
	}

	return http.StatusOK, nil, qualification
}

func (candRepo *CandidateRepositoryImpl) GetOneCandidate(candidateId primitive.ObjectID) (
	httpCode int, repoErr error, candidateResponse response.CandidateResponse) {

	var candidate models.CandidateModel
	mongoErr := candRepo.mongoDB.Collection("candidates").
		FindOne(ctx, bson.M{}).
		Decode(&candidate)

	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения одного кандидата: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, candidateResponse
	}
	candidateResponse.ID = candidate.ID
	candidateResponse.Firstname = candidate.Firstname
	candidateResponse.Lastname = candidate.Lastname
	candidateResponse.Middlename = candidate.Middlename
	candidateResponse.Age = candidate.Age
	candidateResponse.Rank = candidate.Rank
	candidateResponse.Avatar = candidate.Avatar
	candidateResponse.Documents = candidate.Documents

	var q models.QualificationModel
	mongoErr = candRepo.mongoDB.Collection("qualifications").
		FindOne(ctx, bson.M{"_id": candidate.Qualification}).
		Decode(&q)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения одной квалификации: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, candidateResponse
	}
	candidateResponse.Qualification.Value = q.Value
	candidateResponse.Qualification.Notes = q.Notes
	candidateResponse.Qualification.Assessment = q.Assessment

	var competencies []response.CompetenceResponse
	for _, competenceId := range candidate.Competencies {
		var oneC response.CompetenceResponse

		var c models.CompetenceModel
		candRepo.mongoDB.Collection("competences").
			FindOne(ctx, bson.M{"_id": competenceId}).
			Decode(&c)

		oneC.Type = c.Type

		var points []response.PointResponse
		for _, pointId := range c.Points {
			var p models.PointModel
			candRepo.mongoDB.Collection("points").
				FindOne(ctx, bson.M{"_id": pointId}).
				Decode(&p)
			var pResponse response.PointResponse
			pResponse.Value = p.Value
			pResponse.Score = p.Score
			points = append(points, pResponse)
		}
		oneC.Points = points

		oneC.TotalPoints = c.TotalPoints

		competencies = append(competencies, oneC)
	}

	candidateResponse.Competences = competencies

	return http.StatusOK, nil, candidateResponse
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
