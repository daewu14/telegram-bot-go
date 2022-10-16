package tests

import (
	"telegram_bot/src/base"
	"telegram_bot/src/models/response"
	"telegram_bot/src/repositories"
	"telegram_bot/src/services"
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