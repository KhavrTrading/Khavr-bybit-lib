package bybit_ws

// wsAuthRequest is the authentication message sent after connecting.
// Bybit format: {"op": "auth", "args": [apiKey, expires, signature]}
type wsAuthRequest struct {
	Op   string `json:"op"`
	Args []any  `json:"args"`
}

// wsSubscribeRequest is the subscription message for private topics.
// Bybit format: {"op": "subscribe", "args": ["position", "order", ...]}
type wsSubscribeRequest struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

// wsOpResponse is the generic response to auth/subscribe/ping operations.
type wsOpResponse struct {
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	Op      string `json:"op"`
	ConnID  string `json:"conn_id"`
}
