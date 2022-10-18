package services

import (
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/response"
)

type SendPhotoService struct{
	TeleRepo repositories.TelegramInterface
	Data request.SendPhotoRequest
}

func (s SendPhotoService) Do() response.ServiceResponse {

	err, data := s.TeleRepo.SendPhoto(s.Data)
	if err != nil {
		return response.Service().Error(err, nil)
	}

	return response.Service().Success("success send photo", data)
}
