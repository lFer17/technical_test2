package Utils

import (
	"fmt"
	"time"
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func TimerDelay() {
	const timerDefined time.Duration = 2000 * time.Millisecond
	time.Sleep(timerDefined)
}
