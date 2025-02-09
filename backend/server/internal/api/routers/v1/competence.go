package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
)

func NewCompetenceRouters(compControllers *controllers.CompetenceController, router *gin.Engine) {
	competenceRouters := router.Group("/api/v1/competence")

	competenceRouters.GET("/type-of-names", compControllers.GetAllTypeOfNames)
	competenceRouters.GET("/skills", compControllers.GetAllSkillsCompetencies)
}
