package helper

import (
	"fmt"
	"github.com/hedzr/assert"
	"testing"
	"time"
)

func TestGetCurrentEthPrice(t *testing.T) {
	ok, p := GetCurrentEthPrice()
	fmt.Println(p)
	assert.Equal(t, true, ok)
}

func TestParseTimeString(t *testing.T) {
	//data := ParseTimeString("2022-01-20 09:36:23")
	//fmt.Println(data)
	data := ParseTimeString(time.Now().Add(-time.Hour * 10))
	fmt.Println(data)
}

func TestSlice(t *testing.T) {
	s1 := []int64{1, 2, 3, 4}
	s2 := []int64{1, 2, 3, 4}

	SliceDifference(s1, s2)
}
