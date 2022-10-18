package tests

import (
	"github.com/daewu14/telegram-bot-go/src/base"
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/models/response"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/services"
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
