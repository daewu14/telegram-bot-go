package repositories

import (
	"telegram_bot/src/base"
	"telegram_bot/src/models/request"
	"telegram_bot/src/models/response"
	"telegram_bot/src/utils"
)

type TelegramInterface interface {
	GetMe() (error, response.UserResponse)
	SendMessage(data request.SendMessageRequest) (error, response.SendMessageResponse)
	SendPhoto(data request.SendPhotoRequest) (error, response.SendPhotoResponse)
	SetWebhook(data request.SetWebhookRequest) (error, response.SetWebhookResponse)
	GetUpdates() (error, response.GetUpdatesResponse)
}

type TelegramRepository struct{}

// GetMe is method to getting user data
func (t TelegramRepository) GetMe() (error, response.UserResponse) {
	var model response.UserResponse
	call, err := new(base.Api).Get("getMe").Call()
	if err != nil {
		return err, response.UserResponse{}
	}
	errParse := utils.HandleResponseAndParseToModel(call, &model)
	if errParse != nil {
		return errParse, response.UserResponse{}
	}
	return nil, model
}

// SendMessage is method to send message user data
func (t TelegramRepository) SendMessage(data request.SendMessageRequest) (error, response.SendMessageResponse) {
	var model response.SendMessageResponse
	call, err := new(base.Api).Post("sendMessage").Bodys(data).Call()
	if err != nil {
		return err, response.SendMessageResponse{}
	}

	errParse := utils.HandleResponseAndParseToModel(call, &model)
	if err != nil {
		return errParse, response.SendMessageResponse{}
	}
	return nil, model
}

// SendPhoto is method to get updates message of telegram bot
func (t TelegramRepository) SendPhoto(data request.SendPhotoRequest) (error, response.SendPhotoResponse) {
	var model response.SendPhotoResponse
	call, err := new(base.Api).Post("sendPhoto").ClearAllHeader().AddHeader("Accept","application/json").AddHeader("Content-Type","multipart/form-data").Bodys(data).Call()
	if err != nil {
		return err, response.SendPhotoResponse{}
	}

	errParse := utils.HandleResponseAndParseToModel(call, &model)
	if err != nil {
		return errParse, response.SendPhotoResponse{}
	}
	return nil, model
}

// SetWebhook is method to get updates message of telegram bot
func (t TelegramRepository) SetWebhook(data request.SetWebhookRequest) (error, response.SetWebhookResponse) {
	var model response.SetWebhookResponse
	call, err := new(base.Api).Post("setWebhook").Bodys(data).Call()
	if err != nil {
		return err, response.SetWebhookResponse{}
	}

	errParse := utils.HandleResponseAndParseToModel(call, &model)
	if err != nil {
		return errParse, response.SetWebhookResponse{}
	}
	return nil, model
}

// GetUpdates is method to get updates message of telegram bot
func (t TelegramRepository) GetUpdates() (error, response.GetUpdatesResponse) {
	var model response.GetUpdatesResponse
	call, err := new(base.Api).Get("getUpdates").Call()
	if err != nil {
		return err, response.GetUpdatesResponse{}
	}

	errParse := utils.HandleResponseAndParseToModel(call, &model)
	if err != nil {
		return errParse, response.GetUpdatesResponse{}
	}
	return nil, model
}
