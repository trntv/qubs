package http

type httpMessage struct {
	Payload   []byte `json:"payload"`
	Broadcast bool   `json:"broadcast"`
}
