package translate

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type DeepLResponse struct {
	Translations []Translation
}

const endpoint = "https://api-free.deepl.com/v2/translate"

func New(apiKey string, toTranslate string) (DeepLResponse, error) {
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
	err = res.Body.Close()
	if err != nil {
		return deepLResponse, err
	}
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
