package mongoose

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"io/ioutil"
	"log"
	"os"
	"server/internal/domain/types/enum/competence"
	"server/internal/domain/types/models"
	"sync"
	"time"
)

type JsonSkill struct {
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Weight      float64 `bson:"weight" json:"weight"`
}

type JsonCompetence struct {
	TechnicalSkills     []JsonSkill `json:"Технические навыки"`
	SoftSkills          []JsonSkill `json:"Софт-скиллы"`
	SubjectKnowledge    []JsonSkill `json:"Знания предметной области"`
	Achievements        []JsonSkill `json:"Академические достижения"`
	MilitaryPreparation []JsonSkill `json:"Общевоенная подготовка"`
}

var (
	clientInstance *mongo.Client
	clientError    error
	mongoOnce      sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().
			ApplyURI(os.Getenv("MONGO_URI")).
			SetAppName("dashboard")

		clientInstance, clientError = mongo.Connect(clientOptions)
		if clientError != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %s", clientError)
		}

		if err := clientInstance.Ping(ctx, nil); err != nil {
			log.Fatalf("Ошибка пингования MongoDB: %s", err)
		}

		data, readErr := ioutil.ReadFile("/app/data/data.json")
		if readErr != nil {
			log.Fatalf("Ошибка чтения данных: %s", readErr.Error())
		}

		var jsonData []JsonCompetence
		decodeErr := json.Unmarshal(data, &jsonData)
		if decodeErr != nil {
			log.Fatalf("Ошибка анмаршалинга данны: %s", decodeErr.Error())
		}

		for _, c := range jsonData {
			for _, s := range c.TechnicalSkills {
				_, _ = clientInstance.Database("dashboard").
					Collection("skills").
					InsertOne(ctx,
						models.SkillModel{
							Type:        competence.TechnicalSkills,
							Name:        s.Name,
							Description: s.Description,
							Weight:      s.Weight,
						})
			}
			for _, s := range c.SubjectKnowledge {
				_, _ = clientInstance.Database("dashboard").
					Collection("skills").
					InsertOne(ctx, models.SkillModel{
						Type:        competence.DomainKnowledge,
						Name:        s.Name,
						Description: s.Description,
						Weight:      s.Weight,
					})
			}
			for _, s := range c.Achievements {
				_, _ = clientInstance.Database("dashboard").
					Collection("skills").
					InsertOne(ctx, models.SkillModel{
						Type:        competence.AcademicAchievements,
						Name:        s.Name,
						Description: s.Description,
						Weight:      s.Weight,
					})
			}
			for _, s := range c.MilitaryPreparation {
				_, _ = clientInstance.Database("dashboard").
					Collection("skills").
					InsertOne(ctx, models.SkillModel{
						Type:        competence.MilitaryTraining,
						Name:        s.Name,
						Description: s.Description,
						Weight:      s.Weight,
					})
			}
			for _, s := range c.SoftSkills {
				_, _ = clientInstance.Database("dashboard").
					Collection("skills").
					InsertOne(ctx, models.SkillModel{
						Type:        competence.SoftSkills,
						Name:        s.Name,
						Description: s.Description,
						Weight:      s.Weight,
					})
			}
		}

		log.Println("Успешное подключение к MongoDB")
	})

	return clientInstance, clientError
}
