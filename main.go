package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/openai/openai-go/v1"
)

type AppSettings struct {
	OpenAI struct {
		ApiKey string `json:"ApiKey"`
	} `json:"OpenAI"`
}

func main() {
	// Read the OpenAI API key from the appsettings.json file
	appSettings := readAppSettings()
	apiKey := appSettings.OpenAI.ApiKey

	// Set up the OpenAI API client
	client, err := openai.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	// Read the contents of the image file
	imageData, err := ioutil.ReadFile("/path/to/image.jpg")
	if err != nil {
		panic(err)
	}

	// Convert the image data to a base64-encoded string
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	// Send a request to the OpenAI API to extract text from the image
	response, err := client.Image.Extract(imageBase64, "")
	if err != nil {
		panic(err)
	}

	// Print the extracted text
	fmt.Println(response.Text)
}

func readAppSettings() AppSettings {
	// Read the appsettings.json file
	file, err := os.Open("appsettings.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the JSON data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var appSettings AppSettings
	err = json.Unmarshal(data, &appSettings)
	if err != nil {
		panic(err)
	}

	return appSettings
}
