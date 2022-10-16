package request

type SetWebhookRequest struct {
	Url                string   `json:"url"`
	Certificate        string   `json:"certificate"`
	IpAddress          string   `json:"ip_address"`
	MaxConnections     int      `json:"max_connections"`
	AllowedUpdates     []string `json:"allowed_updates"`
	DropPendingUpdates bool     `json:"drop_pending_updates"`
	SecretToken        string   `json:"secret_token"`
}