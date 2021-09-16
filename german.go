package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text string `json:text`
}

type DeepLResponse struct {
	Translations[] Translation
}

func translate (apiKey string, toTranslate string) (DeepLResponse, error) {
	endpoint := "https://api-free.deepl.com/v2/translate"
	data := url.Values{}
	var deepLResponse DeepLResponse
	data.Set("auth_key", apiKey)
	data.Set("target_lang", "DE")
	data.Set("text", toTranslate)
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return deepLResponse, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return deepLResponse, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return deepLResponse, err
	}
	err = json.Unmarshal(body, &deepLResponse)

	if err != nil {
		return deepLResponse, err
	}

	return deepLResponse, nil
}

func main() {
	apiKey := os.Getenv("DEEPL_API_KEY")
	argsWithoutProg := os.Args[1:]
	word := strings.Join(argsWithoutProg," ")

	response, err := translate(apiKey, word)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", "In " + response.Translations[0].DetectedSourceLanguage + ": " + word + "\n")
	fmt.Printf("%s", "In German: " + response.Translations[0].Text)
	fmt.Printf("%s", "\n")
}
