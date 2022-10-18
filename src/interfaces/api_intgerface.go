package interfaces

import "github.com/daewu14/telegram-bot-go/src/base"

type ApiInterface interface {
	Get(endpoint string) base.NetClient
	Post(endpoint string) base.NetClient
	Put(endpoint string) base.NetClient
	Delete(endpoint string) base.NetClient
}
