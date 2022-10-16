package services

import (
	"telegram_bot/src/repositories"
	"telegram_bot/src/response"
)

type GetMeService struct{
	TeleRepo repositories.TelegramInterface
}

func (s GetMeService) Do() response.ServiceResponse {

	err, data := s.TeleRepo.GetMe()
	if err != nil {
		return response.Service().Error(err, nil)
	}

	return response.Service().Success("success get data user", data)
}
