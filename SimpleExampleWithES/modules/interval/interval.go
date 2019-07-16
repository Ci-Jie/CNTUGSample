package interval

import "time"

// Create ...
func Create(second int, intervalHandler func()) {
	ticker := time.NewTicker(time.Second * time.Duration(second))
	go func() {
		for {
			<-ticker.C
			intervalHandler()
		}
	}()
}
