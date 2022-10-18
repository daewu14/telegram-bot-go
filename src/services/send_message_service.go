package services

import (
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/response"
)

type SendMessageService struct{
	TeleRepo repositories.TelegramInterface
	Data request.SendMessageRequest
}

func (s SendMessageService) Do() response.ServiceResponse {

	err, data := s.TeleRepo.SendMessage(s.Data)
	if err != nil {
		return response.Service().Error(err, nil)
	}

	return response.Service().Success("success send message", data)
}
