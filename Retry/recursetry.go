package retry

import (
	"time"
)

type Klausas func(*interface{}) error

func RecurseTry(Klausul Klausas, LoopingControl func(*interface{}) bool, retryCount int, delay time.Duration) Klausas {
	filler := retryCount
	return func(w *interface{}) error {
		for {
			if LoopingControl(w) {
				return nil
			}
	
			err := Klausul(w)
			if err == nil && retryCount > 0 {
				retryCount = filler //fill the retry because current process is success
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

