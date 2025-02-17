package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/internal/api/controllers"
	v1 "server/internal/api/routers/v1"
	"server/internal/domain/repositories"
	"server/internal/domain/usecases"
	"server/pkg/ds"
)

func AppRouters(mongoDB *mongo.Database, ds *ds.DeepSeek) *gin.Engine {
	r := gin.New()

	candRepo := repositories.NewCandidateRepository(mongoDB, ds)
	candUsecase := usecases.NewCandidateUsecase(candRepo)
	candContr := controllers.NewCandidateController(candUsecase)
	v1.NewCandidateRouters(candContr, r)

	compRepo := repositories.NewCompetenceRepository(mongoDB)
	compUsecase := usecases.NewCompetenciesUsecase(compRepo)
	compContr := controllers.NewCompetenceControllers(compUsecase)
	v1.NewCompetenceRouters(compContr, r)

	return r
}
