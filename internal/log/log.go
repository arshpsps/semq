package log

import (
	"sync"

	"github.com/arshpsps/semq/pkg/types"
)

type Log struct {
	messages []types.Message
	mu       sync.RWMutex
}

func NewLog() *Log {
	return &Log{
		messages: make([]types.Message, 0),
	}
}

func (l *Log) Append(payload []byte) int64 {
	l.mu.Lock()
	defer l.mu.Unlock()

	offset := int64(len(l.messages))

	msg := types.NewMessage(offset, payload)
	l.messages = append(l.messages, msg)

	return offset
}

func (l *Log) Read(from int64, max int) []types.Message {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if from >= int64(len(l.messages)) {
		return nil
	}

	end := min(from+int64(max), int64(len(l.messages)))

	return l.messages[from:end]
}
