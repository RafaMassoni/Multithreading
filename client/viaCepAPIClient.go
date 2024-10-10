package client

func getViaCepApiClient(cep string) *CepClient {

	url := "http://viacep.com.br/ws/" + cep + "/json/"

	return &CepClient{
		apiEndpoint: url,
	}
}
