package log

import (
	"sync"

	"github.com/arshpsps/semq/pkg/types"
)

type Log struct {
	messages []types.Message
	mu       sync.RWMutex
}
