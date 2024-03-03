package main

import (
	"github.com/discoriver/gueue/pkg/timeconv"
	"time"
)

type Gueue struct {
	Queue chan Message
}

type MessageTimestamp time.Time

func (t *MessageTimestamp) UnmarshalJSON(b []byte) error {
	tt, err := timeconv.Btot(b)
	if err != nil {
		return err
	}

	*t = MessageTimestamp(tt)

	return nil
}

func (t *MessageTimestamp) MarshalJSON() ([]byte, error) {
	return timeconv.Ttob(time.Time(*t)), nil
}

type Message struct {
	Timestamp MessageTimestamp `json:"timestamp"`
	Topic     string           `json:"topic"`
	Payload   any              `json:"payload"`
}
