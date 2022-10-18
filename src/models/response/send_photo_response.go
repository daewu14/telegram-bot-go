package response

import "github.com/daewu14/telegram-bot-go/src/base"

type SendPhotoResponse struct {
	base.ResponseBase
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date  int `json:"date"`
		Photo []struct {
			FileId       string `json:"file_id"`
			FileUniqueId string `json:"file_unique_id"`
			FileSize     int    `json:"file_size"`
			Width        int    `json:"width"`
			Height       int    `json:"height"`
		} `json:"photo"`
	} `json:"result"`
}
