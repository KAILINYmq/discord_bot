package helper

import (
	"fmt"
	"github.com/hedzr/assert"
	"github.com/levigross/grequests"
	"testing"
	"time"
)

var (
	url = "http://www.baidu.com"
)

func TestRetryTimeout(t *testing.T) {
	err := RetryDo(func() error {

		time.Sleep(time.Second * 2)
		_, err := grequests.Get(url, nil)
		fmt.Println("e", err)
		return err
	}, SetAttempts(3), SetTimeout(time.Second), SetStepDelay(500))

	if err != nil {
		fmt.Println(err)
	}
}

func TestRetry(t *testing.T) {
	err := RetryDo(func() error {
		_, err := grequests.Get(url, nil)
		return err
	}, SetDelay(time.Second), SetAttempts(2))

	assert.Equal(t, err, nil)
}
