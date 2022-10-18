package tests

import (
	"github.com/daewu14/telegram-bot-go/src/base"
	"github.com/daewu14/telegram-bot-go/src/models/response"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/services"
	"testing"
)

func TestGetMeService_Do(t *testing.T) {
	base.SetToken(tokenTest)
	service := new(services.GetMeService)
	service.TeleRepo = repositories.TelegramRepository{}
	data := service.Do()

	if data.Status {
		println("success get user:", data.Data.(response.UserResponse).Result.FirstName)
	}
}