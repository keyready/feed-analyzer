package controllers

import (
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, bindErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (candCont *CandidateController) GetOneCandidate(ctx *gin.Context) {
	getOneCandidate := request.GetOneCandidateRequest{
		ID: ctx.Param("id"),
	}
	httpCode, usecaseErr, candidate := candCont.candidateUsecase.GetOneCandidate(getOneCandidate)
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
