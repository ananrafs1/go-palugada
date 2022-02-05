package retry

import (
	"time"
	"fmt"
)
type WatcherObject interface{}

type Klausas func(*interface{}) error

func RecurseTry(Klausul Klausas, LoopingControl func(*interface{}) bool, retryCount int, delay time.Duration) Klausas {
	return func(w *interface{}) error {
		for {
			if LoopingControl(w) {
				return nil
			}
	
			err := Klausul(w)
			if err == nil && retryCount > 0 {
				continue
			}
			retryCount--
			if retryCount < 1 {
				return ErrRetryExceedMax
			}
			time.Sleep(delay)
		}
	}
}

