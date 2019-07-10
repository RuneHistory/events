package events

import (
	"encoding/json"
	"time"
)

type NewAccountEvent struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *NewAccountEvent) Type() string {
	return "queue.account.new"
}
func (e *NewAccountEvent) Payload() ([]byte, error) {
	return json.Marshal(e)
}
func (e *NewAccountEvent) WithPayload(payload []byte) error {
	return json.Unmarshal(payload, e)
}
