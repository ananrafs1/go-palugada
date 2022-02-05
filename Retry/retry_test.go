package retry

import(
	"errors"
	"time"
	"testing"
	// "fmt"
)

var(
	testStub = 3
)

func Test_Retry_Success(t *testing.T) {
	testFunction := func() (interface{}, error) {
		return 1, nil
	}
	r := Retry(testFunction, 5, time.Duration(5*time.Second))
	ret, _ := r()
	if ret != 1 {
		t.Error("Should return 1, got ", ret)
	}
}

func Test_Retry_SuccessAfterFail(t *testing.T) {
	testFunction := func() (interface{}, error) {
		testStub++
		if testStub < 5 {
			return nil, errors.New("out of range")
		}
		return 1, nil
	}
	r := Retry(testFunction, 5, time.Duration(2*time.Second))
	ret, _ := r()
	if ret != 1 {
		t.Error("Should return 1, got ", ret)
	}
}

func Test_Retry_FailWhenExceedMax( t *testing.T) {
	testFunction := func() (interface{}, error) {
		testStub++
		if testStub < 50 {
			return nil, errors.New("out of range")
		}
		return 1, nil
	}
	r := Retry(testFunction, 5, time.Duration(2*time.Second))
	ret, err := r()
	if ret != nil{
		t.Error("Should error with max attempt, got ", err)
	}
}