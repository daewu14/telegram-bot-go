package tests

import (
	"telegram_bot/src/base"
	"telegram_bot/src/models/request"
	"telegram_bot/src/models/response"
	"telegram_bot/src/repositories"
	"telegram_bot/src/services"
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
