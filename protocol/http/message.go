package http

type httpMessage struct {
	Tags      []string `json:"tags"`
	Payload   []byte   `json:"payload"`
	Broadcast bool     `json:"payload"`
}
