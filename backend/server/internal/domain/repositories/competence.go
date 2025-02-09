package repositories

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"net/http"
	"server/internal/domain/types/enum/competence"
	"server/internal/domain/types/response"
)

type CompetenceRepository interface {
	GetAllSkillsCompetencies() (httpCode int, repErr error, allSkills []response.AllSkills)
	GetAllTypeOfNames() (httpCode int, repErr error, typesOfNames response.TypeOfNames)
}

type CompetenceRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewCompetenceRepository(mongoDB *mongo.Database) *CompetenceRepositoryImpl {
	return &CompetenceRepositoryImpl{mongoDB: mongoDB}
}

func (compR *CompetenceRepositoryImpl) GetAllTypeOfNames() (httpCode int, repErr error, typesOfNames response.TypeOfNames) {
	typesOfNames.Types = []string{
		competence.AcademicAchievements,
		competence.SoftSkills,
		competence.DomainKnowledge,
		competence.TechnicalSkills,
		competence.MilitaryTraining,
	}

	for _, t := range typesOfNames.Types {
		project := bson.D{{"name", 1}}

		cur, mongoErr := compR.mongoDB.Collection("skills").
			Find(ctx, bson.D{{"type", t}}, options.Find().SetProjection(project))
		defer cur.Close(ctx)

		if mongoErr != nil {
			repErr = fmt.Errorf("Ошибка извлечения названия скиллов одного типа компетенции: %s", mongoErr.Error())
			return http.StatusInternalServerError, repErr, typesOfNames
		}

		var names []string
		for cur.Next(ctx) {
			var name string
			if err := cur.Decode(&name); err != nil {
				repErr = fmt.Errorf("Ошибка анмаршалинга одного имени: %s", mongoErr.Error())
				return http.StatusInternalServerError, repErr, typesOfNames
			}
			names = append(names, name)
		}
		typesOfNames.Names = append(typesOfNames.Names, names)
	}

	return http.StatusOK, repErr, typesOfNames
}

func (compR *CompetenceRepositoryImpl) GetAllSkillsCompetencies() (httpCode int, repErr error, allSkills []response.AllSkills) {
	types := []string{
		competence.AcademicAchievements,
		competence.SoftSkills,
		competence.DomainKnowledge,
		competence.TechnicalSkills,
		competence.MilitaryTraining,
	}

	var oneTypeSkills response.AllSkills
	for _, t := range types {
		oneTypeSkills.Type = t

		project := bson.D{{"name", 1}, {"description", 1}}

		cur, mongoErr := compR.mongoDB.Collection("skills").
			Find(ctx, bson.D{{"type", t}}, options.Find().SetProjection(project))
		defer cur.Close(ctx)

		if mongoErr != nil {
			repErr = fmt.Errorf("Ошибка извлечения одного скилла: %s", mongoErr.Error())
			return http.StatusInternalServerError, repErr, nil
		}

		var skills []response.Skill
		for cur.Next(ctx) {
			var skill response.Skill
			if err := cur.Decode(&skill); err != nil {
				repErr = fmt.Errorf("Ошибка анмаршалинга одного скилла: %s", mongoErr.Error())
				return http.StatusInternalServerError, repErr, nil
			}
			skills = append(skills, skill)
		}
		oneTypeSkills.Skills = skills

		allSkills = append(allSkills, oneTypeSkills)
	}

	return http.StatusOK, nil, allSkills
}
