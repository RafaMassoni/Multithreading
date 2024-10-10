package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CepClient struct {
	apiEndpoint string
}

func getAdress(client CepClient, ch chan<- string) {

	req, err := http.NewRequest("GET", client.apiEndpoint, nil)

	if err != nil {
		fmt.Sprintf("\nERROR -> %s  %s", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Sprintf("\nERROR -> %s  %s", err)
	}
	defer res.Body.Close()

	json, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Sprintf("\nERROR -> %s  %s", err)
	}

	if res.StatusCode != 200 {
		fmt.Sprintf("\nERROR -> %s  %s", res.Status, json)
	}

	adress := string(json)

	ch <- adress
}

func GetFirstAdress(cep string) (string, error) {

	brasilApiClient := getBrasilApiClient(cep)
	viaCepApiClient := getViaCepApiClient(cep)

	chBrasilApi := make(chan string)
	chViaCepApi := make(chan string)

	go getAdress(*brasilApiClient, chBrasilApi)
	go getAdress(*viaCepApiClient, chViaCepApi)

	for {
		select {
		case response := <-chBrasilApi: // rabbitmq
			fmt.Printf("Received from brasilApi")
			return response, nil

		case response := <-chViaCepApi: // kafka
			fmt.Printf("Received from viaCepApI")
			return response, nil

		case <-time.After(time.Second * 1):
			return "", errors.New("Timeout")
		}
	}

}
