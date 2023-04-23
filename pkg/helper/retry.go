package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type config struct {
	attempts uint
	delay    time.Duration
	timeout  time.Duration
	step     int64
}

var (
	TimeoutError = errors.New("timeout")
)

type Option func(*config)
type DoFunc func() error

type Error []error

func (e Error) Error() string {
	return e[lenNotNil(e)-1].Error()
}

func (e Error) FullError() string {
	fullErrors := make([]string, lenNotNil(e))
	//return e[lenNotNil(e)].Error()

	for i, err := range e {
		if err != nil {
			fullErrors[i] = fmt.Sprintf("#%d: %s", i+1, err.Error())
		}
	}
	return fmt.Sprintf("full errors: \n%s", strings.Join(fullErrors, "\n"))
}

func lenNotNil(errorSlice []error) int {
	count := 0
	for _, v := range errorSlice {
		if v != nil {
			count += 1
		}
	}
	return count
}

func SetAttempts(attempts uint) Option {
	return func(c *config) {
		c.attempts = attempts
	}
}

func SetDelay(delay time.Duration) Option {
	return func(c *config) {
		c.delay = delay
	}
}

func SetTimeout(timeout time.Duration) Option {
	return func(c *config) {
		c.timeout = timeout
	}
}

func SetStepDelay(step int64) Option {
	return func(c *config) {
		c.step = step
	}
}

func newConf() *config {
	return &config{
		attempts: 3,
		delay:    time.Millisecond * 500,
		timeout:  time.Second * 10,
	}
}

func RetryDo(fun DoFunc, opts ...Option) error {

	conf := newConf()

	for _, opt := range opts {
		opt(conf)
	}
	if conf.step == 0 {
		conf.step = 100
	}

	if conf.attempts == 0 {
		conf.attempts = 5
	}

	var n uint
	var err error
	// 全部的 error 信息
	fullError := make(Error, 0)

	for n < conf.attempts {
		errChan := make(chan error, 1)
		go func() {
			errChan <- fun()
		}()

		select {
		case <-time.After(conf.timeout):
			err = TimeoutError
		case err = <-errChan:

		}

		if err == nil {
			return nil
		}
		fmt.Println(fmt.Sprintf("retry get error, times: %v, error detail: %v", n+1, err.Error()))
		fullError = append(fullError, err)
		n += 1
		time.After(conf.delay + time.Duration(conf.step*int64(n)))
	}
	return err
}
