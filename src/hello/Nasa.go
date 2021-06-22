package hello

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type APOD struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

type NasaPath string

const (
	nasaBaseURL NasaPath = "https://api.nasa.gov"
)

var (
	apodPath NasaPath = "/planetary/apod"
)

func (m NasaPath) fullPath() string {
	return fmt.Sprint(nasaBaseURL + m)
}

type NasaClient struct {
	ApiKey string
}

func (nasaClient *NasaClient) GetAPOD() (*APOD, error) {
	url := apodPath.fullPath()

	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"api_key": nasaClient.ApiKey,
		}).
		Get(url)

	if err != nil {
		return nil, err
	}

	var apod APOD
	if err := json.Unmarshal([]byte(resp.Body()), &apod); err != nil {
		return nil, err
	}

	return &apod, nil
}
