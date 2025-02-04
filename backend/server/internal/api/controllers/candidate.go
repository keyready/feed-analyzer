package controllers

import (
	"github.com/gin-gonic/gin"
	"server/internal/domain/usecases"
)

type CandidateController struct {
	candidateUsecase usecases.CandidateUsecase
}

func NewCandidateController(candUsecase usecases.CandidateUsecase) *CandidateController {
	return &CandidateController{candidateUsecase: candUsecase}
}

func (candCont *CandidateController) GetAllCandidates(ctx *gin.Context) {
	httpCode, usecaseErr, candidates := candCont.candidateUsecase.GetAllCandidates()
	if usecaseErr != nil {
		ctx.AbortWithStatusJSON(httpCode, usecaseErr)
		return
	}
	ctx.JSON(httpCode, candidates)
}
