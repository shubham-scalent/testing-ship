package shiprocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SendRequest(method, path string, BaseURL string, Token string, requestBody interface{}, responseBody interface{}) *ErrorResponse {

	var errResp ErrorResponse

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	url := BaseURL + path

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	if resp.StatusCode != http.StatusOK {
		var apiResponse ErrorResponse
		err = json.Unmarshal(respBody, &apiResponse)
		if err != nil {
			errResp.Message = err.Error()
			return &errResp
		}
		return &apiResponse
	}

	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	return nil
}

func SendRequestWithParams(method, path string, BaseURL string, Token string, requestBody map[string]interface{}, responseBody interface{}) *ErrorResponse {

	var errResp ErrorResponse

	params := url.Values{}
	for key, value := range requestBody {
		params.Add(key, fmt.Sprintf("%v", value))
	}

	url := BaseURL + path
	fullURL := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	if resp.StatusCode != http.StatusOK {
		var apiResponse ErrorResponse
		err = json.Unmarshal(respBody, &apiResponse)
		if err != nil {
			errResp.Message = err.Error()
			return &errResp
		}
		return &apiResponse
	}

	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	return nil
}

func SendRequestNil(method, path string, BaseURL string, Token string, requestBody interface{}, responseBody interface{}) *ErrorResponse {

	var errResp ErrorResponse

	url := BaseURL + path

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	if resp.StatusCode != http.StatusOK {
		var apiResponse ErrorResponse
		err = json.Unmarshal(respBody, &apiResponse)
		if err != nil {
			errResp.Message = err.Error()
			return &errResp
		}
		return &apiResponse
	}

	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		errResp.Message = err.Error()
		return &errResp
	}

	return nil
}
