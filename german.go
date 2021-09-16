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
func main() {
	apiKey := os.Getenv("DEEPL_API_KEY")
	endpoint := "https://api-free.deepl.com/v2/translate"
	argsWithoutProg := os.Args[1:]
	word := strings.Join(argsWithoutProg," ")
	data := url.Values{}
	data.Set("auth_key", apiKey)
	data.Set("target_lang", "DE")
	data.Set("text", word)
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	var deepLResponse DeepLResponse
	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &deepLResponse)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", "In " + deepLResponse.Translations[0].DetectedSourceLanguage + ": " + word + "\n")
	fmt.Printf("%s", "In German: " + deepLResponse.Translations[0].Text)
	fmt.Printf("%s", "\n")
}
