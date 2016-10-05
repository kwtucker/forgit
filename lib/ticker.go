package lib

import "time"

// Ticker is for delaying time for future.
func Ticker(ptime int) {
	ticker := time.NewTicker(time.Second * time.Duration(ptime))
	go func() {
		for range ticker.C {
		}
	}()
	time.Sleep(time.Second * time.Duration(ptime))
	ticker.Stop()
}
