package shiprocket

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ShiprocketClient struct {
	config ClientConfig
	token  tokenConfig
}

func NewShiprocketClient(config ClientConfig) *ShiprocketClient {

	token, err := GetToken(config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &ShiprocketClient{
		config: config,
		token: tokenConfig{
			Token: token,
		},
	}
}

func GetToken(config ClientConfig) (string, error) {
	data := map[string]string{
		"email":    config.Email,
		"password": config.Password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(config.BaseURL+"/v1/external/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response ClientResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Token, err
}
