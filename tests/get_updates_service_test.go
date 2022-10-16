package tests

import (
	"telegram_bot/src/base"
	"telegram_bot/src/models/response"
	"telegram_bot/src/repositories"
	"telegram_bot/src/services"
	"testing"
)

func TestGetUpdatesService_Do(t *testing.T) {
	base.SetToken(tokenTest)
	service := new(services.GetUpdatesService)
	service.TeleRepo = repositories.TelegramRepository{}
	data := service.Do()

	if data.Status {
		println("success get updates:", data.Data.(response.GetUpdatesResponse).Result[0].Message.Chat.FirstName)
	}
}