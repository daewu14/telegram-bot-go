package tests

import (
	"github.com/daewu14/telegram-bot-go/src/base"
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/models/response"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/services"
	"testing"
)

func TestSetWebhookService_Do(t *testing.T) {
	base.SetToken(tokenTest)
	service := new(services.SetWebhookService)
	service.TeleRepo = repositories.TelegramRepository{}
	service.Data = request.SetWebhookRequest{
		Url: "daewu.com",
	}
	data := service.Do()

	println("success set webhook:", data.Data.(response.SetWebhookResponse).Description)
	println("description set webhook:", data.Message)
}
