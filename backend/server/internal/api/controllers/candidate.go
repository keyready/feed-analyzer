package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain/types/request"
	"server/internal/domain/usecases"
)

type CandidateController struct {
	candidateUsecase usecases.CandidateUsecase
}

func NewCandidateController(candUsecase usecases.CandidateUsecase) *CandidateController {
	return &CandidateController{candidateUsecase: candUsecase}
}

func (candCont *CandidateController) AssessmentCandidate(ctx *gin.Context) {
	var candidateData request.CandidateData
	if bindErr := ctx.ShouldBindJSON(&candidateData); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Ошибка получения данных с клиента: %s", bindErr.Error())})
		return
	}

	httpCode, contrErr, qualification := candCont.candidateUsecase.AssessmentCandidate(candidateData)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, qualification)
}

func (candCont *CandidateController) GetOneCandidate(ctx *gin.Context) {
	candidateId := ctx.Param("candidateId")

	httpCode, usecaseErr, candidate := candCont.candidateUsecase.GetOneCandidate(candidateId)
	if usecaseErr != nil {
		ctx.AbortWithStatusJSON(httpCode, usecaseErr)
		return
	}

	ctx.JSON(httpCode, candidate)
}

func (candCont *CandidateController) GetAllCandidates(ctx *gin.Context) {
	httpCode, usecaseErr, candidates := candCont.candidateUsecase.GetAllCandidates()
	if usecaseErr != nil {
		ctx.AbortWithStatusJSON(httpCode, usecaseErr)
		return
	}

	ctx.JSON(httpCode, candidates)
}
