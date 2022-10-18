package base

import "github.com/daewu14/telegram-bot-go/src/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
