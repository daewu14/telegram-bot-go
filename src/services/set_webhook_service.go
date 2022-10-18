package services

import (
	"errors"
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/response"
)

type SetWebhookService struct{
	TeleRepo repositories.TelegramInterface
	Data request.SetWebhookRequest
}

func (s SetWebhookService) Do() response.ServiceResponse {

	err, data := s.TeleRepo.SetWebhook(s.Data)
	if err != nil {
		return response.Service().Error(err, nil)
	}

	if !data.Ok {
		return response.Service().Error(errors.New(data.Description), nil)
	}

	if !data.Result {
		return response.Service().Error(errors.New(data.Description), nil)
	}

	return response.Service().Success(data.Description, data)
}
