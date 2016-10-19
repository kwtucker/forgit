package lib

import "time"

// Ticker is for delaying time for future.
func Ticker(ptime int) {
	if ptime > 0 {
		ticker := time.NewTicker(time.Minute * time.Duration(ptime))
		go func() {
			for range ticker.C {
			}
		}()
		time.Sleep(time.Minute * time.Duration(ptime))
		ticker.Stop()
	}
}
