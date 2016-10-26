package lib

import (
	"time"
)

// Ticker is for delaying time for future.
func Ticker(ptime int) {
	if ptime > 0 {
		ticker := time.NewTimer(time.Minute * time.Duration(ptime))
		<-ticker.C
	}
}
