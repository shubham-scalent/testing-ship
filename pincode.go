package shiprocket

func (s *ShiprocketClient) GetLocalityDetails(request GetLocalityDetailsRequest) (*GetLocalityDetailsResponse, *ErrorResponse) {
	var response GetLocalityDetailsResponse
	requestBody := map[string]interface{}{"postcode": request.Postcode}
	errResp := SendRequestWithParams("GET", "/v1/external/open/postcode/details", BASE_URL, s.tokenConfig.Token, requestBody, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
