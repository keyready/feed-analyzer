package usecases

import (
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/enum"
	"server/internal/domain/types/models"
)

type CompetenceUsecase interface {
	GetAllCompetencies() (httpCode int, usecaseErr error, competencies []models.CompetenceModel)
	GetAllTypesCompetencies() (httpCode int, usecaseErr error, types []enum.TypeCompetence)
}

type CompetenceUsecaseImpl struct {
	compRepo repositories.CompetenceRepository
}

func NewCompetenciesUsecase(compRepo repositories.CompetenceRepository) *CompetenceUsecaseImpl {
	return &CompetenceUsecaseImpl{compRepo: compRepo}
}

func (compUsecase *CompetenceUsecaseImpl) GetAllTypesCompetencies() (httpCode int, usecaseErr error, types []enum.TypeCompetence) {
	types = []enum.TypeCompetence{
		enum.SoftSkills,
		enum.TeamMethodicalSkills,
		enum.TechnicalSkills,
		enum.AcademicAchievements,
		enum.DomainKnowledge,
	}
	return http.StatusOK, nil, types
}

func (compUsecase *CompetenceUsecaseImpl) GetAllCompetencies() (httpCode int, usecaseErr error, competencies []models.CompetenceModel) {
	httpCode, usecaseErr, competencies = compUsecase.compRepo.GetAllCompetencies()
	if usecaseErr != nil {
		return httpCode, usecaseErr, competencies
	}
	return httpCode, nil, competencies
}
