package client

func getBrasilApiClient(cep string) *CepClient {

	url := "https://brasilapi.com.br/api/cep/v1/" + cep

	return &CepClient{
		apiEndpoint: url,
	}
}
