package contextime

import (
	"context"
	"time"
)

// ResetFunc resets the context timeout timer
type ResetFunc func()

// WithTimeoutReset returns a child context which is canceled after the provided duration elapses.
// The returned ResetFunc may be called before the context is canceled to restart the timeout timer.
// Unlike context.WithTimeout, the returned context will not report the correct deadline.
func WithTimeoutReset(parent context.Context, d time.Duration) (context.Context, context.CancelFunc, ResetFunc) {
	ctx, cancel0 := context.WithCancel(parent)
	timer := time.AfterFunc(d, cancel0)
	cancel := func() {
		cancel0()
		timer.Stop()
	}
	reset := func() { timer.Reset(d) }
	return ctx, cancel, reset
}
