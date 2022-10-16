package tests

import (
	"telegram_bot/src/base"
	"telegram_bot/src/models/request"
	"telegram_bot/src/models/response"
	"telegram_bot/src/repositories"
	"telegram_bot/src/services"
	"testing"
)

func TestSendMessageService_Do(t *testing.T) {
	base.SetToken(tokenTest)
	service := new(services.SendMessageService)
	service.TeleRepo = repositories.TelegramRepository{}
	service.Data = request.SendMessageRequest{
		ChatId:                "433657762",
		Text:                  "Halooo",
		ParseMode:             "",
		DisableWebPagePreview: false,
		DisableNotification:   false,
		ProtectContent:        false,
	}
	data := service.Do()

	if data.Status {
		println("success send message user:", data.Data.(response.SendMessageResponse).Result.MessageId)
	}
}
