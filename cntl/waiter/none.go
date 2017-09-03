package waiter

// None is a waiter that does nothing
type None struct{}

// NewNone creates a new None waiter
func NewNone() *None {
	return &None{}
}

// Wait waits for a specific event to happen. In this case, nothing.
func (t *None) Wait(done chan struct{}, cancel chan struct{}, err chan error) error {
	done <- struct{}{}
	return nil
}
