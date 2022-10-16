package base

import "telegram_bot/src/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
