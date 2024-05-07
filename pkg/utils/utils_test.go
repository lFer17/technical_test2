package Utils_test

import (
	Utils "lfer17/user/test/pkg/utils"
	"testing"
	"time"
)

func TestTimerDelay(t *testing.T) {
	start := time.Now()

	Utils.TimerDelay()
	// flag for time delay
	duration := time.Since(start)

	// verify if duration is less than 2000 milliseconds
	if duration < 2000*time.Millisecond {
		t.Errorf("El retraso fue menor de lo esperado. Duración: %s", duration)
	}
}
