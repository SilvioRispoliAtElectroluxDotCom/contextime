package contextime

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithTimeoutReset(t *testing.T) {
	assert := assert.New(t)

	// Set up a watchdog
	watchdog, cancelWatchDog, kickWatchdog := WithTimeoutReset(context.Background(), 2*time.Second)
	defer cancelWatchDog()

	// Doing some heavy lifting
	time.Sleep(750 * time.Millisecond)

	// Done! Kick the watchdog
	kickWatchdog()

	// .. However if we do something heavier ..

	select {
	case <-time.After(10 * time.Second):
		assert.Fail("the watchdog did not trigger")

	case <-watchdog.Done():
		assert.True(true)
	}
}
