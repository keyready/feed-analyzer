package usecases

import (
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/response"
)

type CompetenceUsecase interface {
	GetAllTypeOfNames() (httpCode int, usecaseErr error, types response.TypeOfNames)
	GetAllSkillsCompetencies() (httpCode int, usecaseErr error, allBodyCompetenciesResponse []response.AllSkills)
}

type CompetenceUsecaseImpl struct {
	compRepo repositories.CompetenceRepository
}

func NewCompetenciesUsecase(compRepo repositories.CompetenceRepository) *CompetenceUsecaseImpl {
	return &CompetenceUsecaseImpl{compRepo: compRepo}
}

func (compUsecase *CompetenceUsecaseImpl) GetAllSkillsCompetencies() (httpCode int, usecaseErr error, allBodyCompetencies []response.AllSkills) {
	httpCode, usecaseErr, allBodyCompetencies = compUsecase.compRepo.GetAllSkillsCompetencies()
	if usecaseErr != nil {
		return httpCode, usecaseErr, nil
	}
	return httpCode, nil, allBodyCompetencies
}

func (compUsecase *CompetenceUsecaseImpl) GetAllTypeOfNames() (httpCode int, usecaseErr error, types response.TypeOfNames) {
	httpCode, usecaseErr, types = compUsecase.compRepo.GetAllTypeOfNames()
	return http.StatusOK, nil, types
}
