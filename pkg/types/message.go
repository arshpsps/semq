package types

import "time"

type Message struct {
	Offset    int64
	Payload   []byte
	Timestamp int64
}

func NewMessage(offset int64, payload []byte) Message {
	return Message{
		Offset:    offset,
		Payload:   payload,
		Timestamp: time.Now().UnixNano(),
	}
}

func (m Message) Size() int {
	return len(m.Payload)
}
