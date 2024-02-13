package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/shubham-scalent/testing-ship/apimodel"
)

type ShiprocketServiceImpl struct {
}

func NewShiprocketServiceImpl() (*ShiprocketServiceImpl, error) {
	return &ShiprocketServiceImpl{}, nil
}

func (s *ShiprocketServiceImpl) GetToken(config apimodel.Config) (string, error) {
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

	var response apimodel.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Token, err
}

func (s *ShiprocketServiceImpl) GetAvailableCouriers(request apimodel.CourierAvailabityRequest, config apimodel.Config, token string) (string, error) {
	// Create a new request
	resp, err := SendRequest("GET", "/v1/external/courier/courierListWithCounts", config.BaseURL, token, request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var abc string
	// var response ShippingAddressUpdateResponse
	err = ReadResponse(resp, &abc)
	if err != nil {
		return "", err
	}

	return abc, nil
}
