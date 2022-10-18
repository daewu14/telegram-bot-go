package tests

import (
	"github.com/daewu14/telegram-bot-go/src/base"
	"github.com/daewu14/telegram-bot-go/src/models/response"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/services"
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