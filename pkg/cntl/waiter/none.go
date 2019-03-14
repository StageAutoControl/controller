package waiter

import "github.com/StageAutoControl/controller/pkg/internal/logging"

// None is a waiter that does nothing
type None struct {
	logger logging.Logger
}

// NewNone creates a new None waiter
func NewNone(logger logging.Logger) *None {
	return &None{logger}
}

// Wait waits for a specific event to happen. In this case, nothing.
func (t *None) Wait(done chan struct{}, cancel chan struct{}, err chan error) error {
	t.logger.Info("Not waiting")
	done <- struct{}{}
	return nil
}
