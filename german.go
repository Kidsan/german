package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kidsan/german/translate"
)

func main() {
	apiKey := os.Getenv("DEEPL_API_KEY")
	argsWithoutProg := os.Args[1:]
	word := strings.Join(argsWithoutProg, " ")

	response, err := translate.New(apiKey, word)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", "In "+response.Translations[0].DetectedSourceLanguage+": "+word+"\n")
	fmt.Printf("%s", "In German: "+response.Translations[0].Text)
	fmt.Printf("%s", "\n")
}
