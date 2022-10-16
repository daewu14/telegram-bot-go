package services

import (
	"telegram_bot/src/models/request"
	"telegram_bot/src/repositories"
	"telegram_bot/src/response"
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
