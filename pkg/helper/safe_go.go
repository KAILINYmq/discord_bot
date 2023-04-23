package helper

import "fmt"

func GoSafe(f func()) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("safe go recover panic:", err)
		}
	}()
	go func() {
		f()
	}()
}
