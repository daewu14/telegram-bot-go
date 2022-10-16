package services

import (
	"telegram_bot/src/models/request"
	"telegram_bot/src/repositories"
	"telegram_bot/src/response"
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
