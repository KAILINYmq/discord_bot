package helper

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"time"
)

var solPrice float64

func GetUTTokenQuantity(ownerAddress string) (decimal.Decimal, error) {

	return decimal.Zero, nil
}

type solData struct {
	Quote struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

type Sol struct {
	SOL []solData `json:"SOL"`
}

type respData struct {
	Data Sol `json:"data"`
}

func GetSolPriceFromChain() (error, float64) {
	url := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?symbol=SOL"

	reqOption := grequests.RequestOptions{
		Headers: map[string]string{
			"X-CMC_PRO_API_KEY": "51a5341c-303a-479b-991d-707aa77914d7",
			"Content-Type":      "application/json",
		},
		DialTimeout:         10 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	response, err := grequests.DoRegularRequest(http.MethodGet, url, &reqOption)
	if err != nil {
		fmt.Println("do request error:", err.Error())
		return err, 0
	}

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()

	var result respData

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err, 0
	}
	if len(result.Data.SOL) == 0 {
		return err, 0
	}
	price := result.Data.SOL[0].Quote.USD.Price
	return nil, price
}

func GetSolPrice() float64 {
	return solPrice
}

func init() {
	go func() {
		for {
			err, price := GetSolPriceFromChain()
			if err != nil {
				fmt.Println("price error:", err.Error())
			}
			solPrice = price
			fmt.Println("set solana price")
			<-time.After(time.Minute * 30)
		}
	}()
}
