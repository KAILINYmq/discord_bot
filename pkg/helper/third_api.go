package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
)

const (
	twitterApiKey            = "WvxBDjChlCAs1FFmwTeSoV4iV"
	twitterApiSecret         = "Q9JY7O36EBBQeJBKKENd53jg2MtTYVnMTTiyflqgj2NcdWWcZZ"
	twitterAccessToken       = "1427134094341677061-QQcxRFSQyOY6iGidaEilCqnarRz9XQ"
	twitterAccessTokenSecret = "vjrMOs7hC1NZX0a3xTewiIiLLJHRlT3XMgJhCtu0sxxM0"
	accountId                = "18ce55fl565"
	oauthId                  = "eGhKVnBWLVZ1S0xoTHlIaDJVbVM6MTpjaQ"
	oauthSecret              = "jrbwq6yuIoVRkQcRFzWafyMx-gKf_QF4EXrM5alDQUGWcobgGf"
)

// twitter api
func twitterClient() *http.Client {
	config := oauth1.NewConfig(twitterApiKey, twitterApiSecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	//config := &clientcredentials.Config{
	//	ClientID:     oauthId,
	//	ClientSecret: oauthSecret,
	//	TokenURL:     "https://api.twitter.com/oauth2/token",
	//}
	//httpClient := config.Client(oauth2.NoContext)

	return httpClient
}

type MediaUpload struct {
	MediaId       int    `json:"media_id"`
	MediaIdString string `json:"media_id_string"`
	MediaKey      string `json:"media_key"`
}

type MediaInfo struct {
	Data struct {
		MediaUrl string `json:"media_url"`
	}
}

func TwitterUploadMedia(fileName string, fileData string) (interface{}, error) {
	client := twitterClient()

	b := &bytes.Buffer{}
	form := multipart.NewWriter(b)

	fw, err := form.CreateFormFile("media", fileName)
	if err != nil {
		return "", err
	}

	//body, _ := file.Open()

	ok, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, fileData)
	if !ok {
		return "", fmt.Errorf("data is not base64 data")
	}

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	//allData := re.FindAllSubmatch([]byte(fileData), 2)
	base64Str := re.ReplaceAllString(fileData, "")

	imageData, _ := base64.StdEncoding.DecodeString(base64Str)
	body := bytes.NewReader(imageData)
	_, err = io.Copy(fw, body)
	if err != nil {
		return "", err
	}
	form.Close()
	resp, err := client.Post("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image",
		form.FormDataContentType(), bytes.NewReader(b.Bytes()))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()
	m := &MediaUpload{}
	_ = json.NewDecoder(resp.Body).Decode(m)
	var dd []byte
	resp.Body.Read(dd)
	fmt.Println("media:", m, "raw:", string(dd), resp.StatusCode)
	resp1, err := client.Post(fmt.Sprintf("https://ads-api.twitter.com/11/accounts/%v/media_library?media_key=%v", accountId, m.MediaKey), "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp1.Body.Close()
	var mm map[string]interface{}

	byteData, err := io.ReadAll(resp1.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(byteData, &mm)
	if err != nil {
		return "", err
	}
	fmt.Println("media image:", mm)
	return mm, nil
}
