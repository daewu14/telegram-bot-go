package telegram_bot

import (
	"github.com/daewu14/telegram-bot-go/src/base"
	"github.com/daewu14/telegram-bot-go/src/models/request"
	"github.com/daewu14/telegram-bot-go/src/repositories"
	"github.com/daewu14/telegram-bot-go/src/response"
	"github.com/daewu14/telegram-bot-go/src/services"
)

// teleRepo is method to simplify repositories.TelegramRepository getter
func teleRepo() repositories.TelegramRepository {
	return repositories.TelegramRepository{}
}

// do is method to call/doing the service called
func do(service base.ServiceInterface) response.ServiceResponse {
	return service.Do()
}

// SetToken is method to set the awesome telegram bot token
func SetToken(token string) { base.SetToken(token) }

// GetToken is method to get the awesome telegram bot token
func GetToken() string { return base.GetToken() }

// GetMe is method to get information of the awesome telegram bot registered
func GetMe() response.ServiceResponse {
	return do(services.GetMeService{TeleRepo: teleRepo()})
}

// GetUpdates is method to get all update from telegram bot registered
func GetUpdates() response.ServiceResponse {
	return do(services.GetUpdatesService{TeleRepo: teleRepo()})
}

// SendMessage is method to send message by telegram
func SendMessage(data request.SendMessageRequest) response.ServiceResponse {
	return do(services.SendMessageService{TeleRepo: teleRepo(), Data: data})
}

// SetWebhook is method to set telegram bot webhook configuration
func SetWebhook(data request.SetWebhookRequest) response.ServiceResponse {
	return do(services.SetWebhookService{TeleRepo: teleRepo(), Data: data})
}