package controllers

import (
	"github.com/gin-gonic/gin"
	"server/internal/domain/usecases"
)

type CompetenceController struct {
	compUsecase usecases.CompetenceUsecase
}

func NewCompetenceControllers(compUsecase usecases.CompetenceUsecase) *CompetenceController {
	return &CompetenceController{compUsecase: compUsecase}
}

func (compContr *CompetenceController) GetAllTypesCompetencies(ctx *gin.Context) {
	httpCode, _, types := compContr.compUsecase.GetAllTypesCompetencies()
	ctx.JSON(httpCode, types)
}

func (compContr *CompetenceController) GetAllCompetencies(ctx *gin.Context) {
	httpCode, usecaseErr, competencies := compContr.compUsecase.GetAllCompetencies()
	if usecaseErr != nil {
		//handlers.ErrorHandler(ctx, usecaseErr)
		ctx.AbortWithStatusJSON(httpCode, usecaseErr)
		return
	}

	ctx.JSON(httpCode, competencies)
}
