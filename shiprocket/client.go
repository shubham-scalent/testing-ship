package shiprocket

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ShiprocketClient struct {
	Config ClientConfig
	Token  tokenConfig
}

func NewShiprocketClient(config ClientConfig) (*ShiprocketClient, error) {

	token, err := GetToken(config)
	if err != nil {
		return nil, err
	}

	return &ShiprocketClient{
		Config: config,
		Token: tokenConfig{
			Token:             token,
			CreatedOn:         time.Now(),
			NextRefreshOnTime: time.Now().Add(8 * time.Hour),
		},
	}, nil
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
