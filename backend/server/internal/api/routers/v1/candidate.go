package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
)

func NewCandidateRouters(candContr *controllers.CandidateController, router *gin.Engine) {
	candidatesRouters := router.Group("/api/v1/candidates")

	candidatesRouters.GET("/get-all", candContr.GetAllCandidates)
}
