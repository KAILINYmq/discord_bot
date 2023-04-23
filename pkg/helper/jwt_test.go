package helper

import (
	"fmt"
	"github.com/hedzr/assert"
	"testing"
)

func TestNewPayJwt(t *testing.T) {
	jwt := NewPayJwt("0x37", "1", "4", "1.2", "100001")

	sign, _ := jwt.GenSignature()
	fmt.Println("sign", sign)
	ok := jwt.VerifySignature(sign)

	assert.Equal(t, true, ok)
}
