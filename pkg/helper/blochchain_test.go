package helper

import (
	"fmt"
	"testing"
)

func TestGetSolPrice(t *testing.T) {
	err, price := GetSolPriceFromChain()
	fmt.Println(err, price)
}
