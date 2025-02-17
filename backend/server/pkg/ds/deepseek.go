package ds

import (
	"context"
	"fmt"
	"github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
	"log"
	"server/internal/domain/types/enum/competence"
)

type DeepSeek struct {
	client *deepseek.Client
}

func (ds *DeepSeek) Init(token string) {
	ds.client = deepseek.NewClient(token)
}

func (ds *DeepSeek) Assessment(deepseekRequest map[string]float64) string {
	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekReasoner,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: constants.ChatMessageRoleSystem, Content: fmt.Sprintf(`
	Ты являешься опытным HR-менеджером. Тебе необходимо оценить выпускника кафедры IT-технологий.
	На вход ты получишь 5 различных типов компетенций и 5 значений, которые имеет выпускник в соответсвии с каждым типом.
	Тебе необходимо оценить данного выпускника значением от 0 до 100 и дать ему краткую характеристику.
	
	Входные данные: 
		%s - %s
		%s - %s
		%s - %s
		%s - %s
		%s - %s

	Формат ответа:
		Оценка - [значение от 0 до 100]
		Характеристика - [краткая характеристика]
`, competence.TechnicalSkills, deepseekRequest[competence.TechnicalSkills],
				competence.SoftSkills, deepseekRequest[competence.SoftSkills],
				competence.AcademicAchievements, deepseekRequest[competence.AcademicAchievements],
				competence.MilitaryTraining, deepseekRequest[competence.MilitaryTraining],
				competence.DomainKnowledge, deepseekRequest[competence.DomainKnowledge])},
		},
	}

	response, err := ds.client.CreateChatCompletion(context.TODO(), request)
	if err != nil {
		log.Fatalf("Ошибка запроса к ds: %v", err.Error())
	}

	return response.Choices[0].Message.Content
}
