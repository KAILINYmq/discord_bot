package syncer

import (
	"DiscordRolesBot/global"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"io/ioutil"
	"net/http"
	"time"
)

func getMint(tokenId string) *MintInfo {
	url := global.Config.Syncer.URL + "/search/mint"

	data := map[string]interface{}{
		"token_id": tokenId,
	}

	response, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:           data,
		RequestTimeout: time.Second * 5,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil
	}

	if response.StatusCode != http.StatusOK {
		return nil
	}

	var resp MintInfo

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil
	}
	return &resp
}

func getLevel(tokenId string) *LevelData {
	url := global.Config.Syncer.URL + fmt.Sprintf("/level/%v", tokenId)
	fmt.Println("url", url)

	response, err := grequests.Get(url, &grequests.RequestOptions{
		RequestTimeout: time.Second * 5,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil
	}
	if response.StatusCode != http.StatusOK {
		return nil
	}

	var resp LevelData

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil
	}
	return &resp
}

func GetNftAddress(contract string, tokenId string) (string, error) {
	nftInfo := getMint(tokenId)
	if nftInfo == nil {
		return "", fmt.Errorf("get mint info error")
	}
	if nftInfo.Data.Nft.CreatorAddress != contract {
		return "", fmt.Errorf("invalid contract")
	}
	return nftInfo.Data.Nft.MintAddress, nil
}

func GetNftLevel(tokenId string) (int, error) {
	nftInfo := getLevel(tokenId)
	if nftInfo == nil {
		return 0, fmt.Errorf("get nft level error")
	}
	return nftInfo.Level, nil
}

func TransactionMiddleware(url string, data map[string]interface{}) (map[string]interface{}, error) {
	response, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:           data,
		RequestTimeout: time.Second * 40,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}
	var resp map[string]interface{}
	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	resp["statusCode"] = response.StatusCode
	return resp, err
}

func GetCaredNftStatus(tokenIds []string) []CardMetadata {
	url := global.Config.Syncer.URL + "/launchpad/card-draw-tokens-info"

	data := map[string]interface{}{
		"token_ids": tokenIds,
	}

	response, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:           data,
		RequestTimeout: time.Second * 5,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil
	}

	if response.StatusCode != http.StatusOK {
		return nil
	}

	var resp CardMetaInfos

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil
	}
	return resp.Tokens
}

func DoHttpPost(url string, headers map[string]string, data map[string]interface{}, respData interface{}) error {
	headers["Content-Type"] = "application/json"
	response, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:           data,
		RequestTimeout: time.Second * 20,
		Headers:        headers,
	})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println(url)
	fmt.Println(string(body))
	err = json.Unmarshal(body, respData)
	return err
}
