package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain/usecases"
)

type CompetenceController struct {
	compUsecase usecases.CompetenceUsecase
}

func NewCompetenceControllers(compUsecase usecases.CompetenceUsecase) *CompetenceController {
	return &CompetenceController{compUsecase: compUsecase}
}

func (compContr *CompetenceController) GetAllSkillsCompetencies(ctx *gin.Context) {
	httpCode, usecaseErr, allBodyCompetenciesResponse := compContr.compUsecase.GetAllSkillsCompetencies()
	if usecaseErr != nil {
		ctx.AbortWithStatusJSON(httpCode, usecaseErr)
		return
	}

	ctx.JSON(http.StatusOK, allBodyCompetenciesResponse)
}

func (compContr *CompetenceController) GetAllTypeOfNames(ctx *gin.Context) {
	httpCode, _, types := compContr.compUsecase.GetAllTypeOfNames()
	ctx.JSON(httpCode, types)
}
