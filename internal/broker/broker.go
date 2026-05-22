package broker

import (
	"errors"
	"strings"

	"github.com/arshpsps/semq/internal/log"
	"github.com/arshpsps/semq/pkg/types"
)

type Broker struct {
	topics map[string]*log.Log
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]*log.Log),
	}
}

func (b *Broker) Produce(topic string, payload []byte) (int64, error) {
	if strings.TrimSpace(topic) == "" {
		return 0, errors.New("no topic name")
	}

	l := b.getOrCreateLog(topic)
	
	offset := l.Append(payload)

	return offset, nil
}

func (b *Broker) Consume(topic string, offset int64, max int) ([]types.Message, error) {
	if strings.TrimSpace(topic) == "" {
		return nil, errors.New("no topic name")
	}

	if offset < 0 {
		return nil, errors.New("invalid offset")
	}

	if max <= 0 {
		return nil, errors.New("invalid max")
	}

	l, exists := b.topics[topic]
	if !exists {
		return nil, errors.New("topic not found")
	}

	return l.Read(offset, max), nil
}

func (b *Broker) getOrCreateLog(topic string) *log.Log {
	if l, exists := b.topics[topic]; exists {
		return l
	}
	l := log.NewLog()
	b.topics[topic] = l
	return l
}
