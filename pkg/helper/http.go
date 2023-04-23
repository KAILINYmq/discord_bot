package helper

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"io/ioutil"
	"time"
)

func DoGet(url string, header, param map[string]string, resp interface{}) (int, error) {
	fmt.Println(url)
	response, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: time.Second * 10,
		Headers:        header,
		Params:         param,
	})
	if err != nil {
		return 500, err
	}
	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()
	fmt.Println(string(body))
	err = json.Unmarshal(body, resp)
	if err != nil {
		return 500, err
	}
	return response.StatusCode, nil
}

func DoPut(url string, header map[string]string, data map[string]interface{}, resp interface{}) (int, error) {
	fmt.Println(url)
	response, err := grequests.Put(url, &grequests.RequestOptions{
		RequestTimeout: time.Second * 10,
		Headers:        header,
		JSON:           data,
	})
	if err != nil {
		return 500, err
	}
	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()
	if resp != nil {
		err = json.Unmarshal(body, resp)
	}
	if err != nil {
		return 500, err
	}

	return response.StatusCode, nil
}
