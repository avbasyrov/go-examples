package main

// go-http-client

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient interface {
	Get(url string) (string, error)
}

type SimpleHttpClient struct{}

func (s *SimpleHttpClient) Get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyAsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	bodyAsText := string(bodyAsBytes)

	return bodyAsText, nil
}

type CbrRates struct {
	httpClient HttpClient
}

func (c *CbrRates) Get() (string, error) {
	return c.httpClient.Get("https://www.cbr.ru/scripts/XML_daily.asp")
}

func main() {
	// Получение котировок валют от Центробанка
	rates := CbrRates{httpClient: &SimpleHttpClient{}}
	log.Println(rates.Get())
}
