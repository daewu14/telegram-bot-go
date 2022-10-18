package services

import (
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/response"
)

type GetUpdatesService struct{ TeleRepo repositories.TelegramInterface }

func (s GetUpdatesService) Do() response.ServiceResponse {
	err, data := s.TeleRepo.GetUpdates()
	if err != nil {
		return response.Service().Error(err, nil)
	}

	return response.Service().Success("success get updates data", data)
}
