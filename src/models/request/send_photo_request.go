package request

type SendPhotoRequest struct {
	ChatId                string `json:"chat_id"`
	PhotoBase64Encode     string `json:"photo"`
	Caption               string `json:"caption"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ProtectContent        bool   `json:"protect_content"`
}
