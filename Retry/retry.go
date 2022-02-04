package retry
import(
	"time"
	"errors"
	// "fmt"
)

type Klausa func() (interface{}, error)

var ErrRetryExceedMax = errors.New("Retry reach Maximum, and still error")



func Retry(Klausul Klausa, retryCount int, delay time.Duration) Klausa {
	return func() (interface{}, error) {
		for {
			ret, err := Klausul()
			if err == nil && retryCount > 0 {
				return ret, nil
			}
			retryCount--
			if retryCount < 1 {
				return nil, ErrRetryExceedMax
			}
			time.Sleep(delay)
		}
	}
}