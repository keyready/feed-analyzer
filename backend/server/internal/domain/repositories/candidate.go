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
	"server/pkg/ds"
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
	mongoDB        *mongo.Database
	deepSeekClient *ds.DeepSeek
}

func NewCandidateRepository(mongoDB *mongo.Database, ds *ds.DeepSeek) *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{mongoDB: mongoDB, deepSeekClient: ds}
}

func (candRepo *CandidateRepositoryImpl) AssessmentCandidate(candidateData request.CandidateData) (
	httpCode int, repoErr error, qualification models.QualificationModel) {

	deepSeekRequest := make(map[string]float64)

	types := []string{
		competence.AcademicAchievements,
		competence.SoftSkills,
		competence.TechnicalSkills,
		competence.MilitaryTraining,
		competence.DomainKnowledge,
	}

	var competencies []bson.ObjectID

	for _, t := range types {
		var totalScore float64
		var totalPoints float64
		var points []bson.ObjectID

		lenSkills, _ := candRepo.mongoDB.Collection("skills").CountDocuments(ctx, bson.D{{"type", t}})

		for _, c := range candidateData.Competences {
			if c.Type == t {
				var skill models.SkillModel
				_ = candRepo.mongoDB.Collection("skills").
					FindOne(ctx, bson.D{
						{"type", t},
						{"name", c.Name}},
					).
					Decode(&skill)

				score := skill.Weight * c.Value
				totalScore += score

				pointId := bson.NewObjectID()
				point := models.PointModel{
					ID:    pointId,
					Value: c.Value,
					Score: score,
				}
				_, _ = candRepo.mongoDB.Collection("points").
					InsertOne(ctx, point)

				points = append(points, pointId)
			}
		}

		compId := bson.NewObjectID()
		totalPoints = totalScore / float64(lenSkills)

		deepSeekRequest[t] = totalScore

		_, _ = candRepo.mongoDB.Collection("competences").
			InsertOne(ctx, models.CompetenceModel{
				ID:          compId,
				Type:        t,
				Points:      points,
				TotalPoints: totalPoints,
			})
		competencies = append(competencies, compId)
	}

	deepSeekResult := candRepo.deepSeekClient.Assessment(deepSeekRequest)

	fmt.Println(deepSeekResult)

	_, _ = candRepo.mongoDB.Collection("candidates").
		InsertOne(ctx, models.CandidateModel{
			Firstname:    candidateData.Firstname,
			Lastname:     candidateData.Lastname,
			Middlename:   candidateData.Middlename,
			Age:          candidateData.Age,
			Rank:         candidateData.Rank,
			Avatar:       candidateData.Avatar,
			Documents:    candidateData.DocumentsNames,
			Competencies: competencies,
		})

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
