# GERMAN
A cli tool for converting text into German.

## Build Locally
```bash
$> go build
$> go install
```
## Dependencies
1. To execute successfully, a free tier DEEPL API key must be created. This can be done on the DEEPL website.
2. The application looks for an environment variable containing the API key: "DEEPL_API_KEY"

## Usage
```bash
$> german what time is it?
In EN: what time is it?
In German: Wie spät ist es?
$>
```

## Why German?
This tool could easily have been written to convert from any source language to any target language, thanks to
the way the deepl api can automatically detect source languages. However, this tool was written for my convenience to
be able to very quickly translate words from English to German without navigating around a browser. Although other 
such tools exist and are very capable, I took the opportunity to practice making API calls in golang and attempt to write
somewhat idiomatic golang.