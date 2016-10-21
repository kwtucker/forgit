package lib

import (
	"time"
)

// Ticker is for delaying time for future.
func Ticker(ptime int) {
	if ptime > 0 {
		ticker := time.NewTimer(time.Minute * time.Duration(ptime))
		<-ticker.C
		// ticker := time.NewTicker(time.Minute * time.Duration(ptime))
		// go func() {
		// 	for i := range ticker.C {
		// 		fmt.Println("Tick at", i)
		// 	}
		// }()
		// time.Sleep(time.Minute * time.Duration(ptime))
		// ticker.Stop()
	}
}
