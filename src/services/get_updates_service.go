package services

import (
	"telegram_bot/src/repositories"
	"telegram_bot/src/response"
)

type GetUpdatesService struct{ TeleRepo repositories.TelegramInterface }

func (s GetUpdatesService) Do() response.ServiceResponse {
	err, data := s.TeleRepo.GetUpdates()
	if err != nil {
		return response.Service().Error(err, nil)
	}

	return response.Service().Success("success get updates data", data)
}
