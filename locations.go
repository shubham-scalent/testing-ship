package shiprocket

func (s *ShiprocketClient) AddPickUpLocation(request AddPickUpLocationRequest) (*AddPickUpLocationResponse, *ErrorResponse) {
	var response AddPickUpLocationResponse
	errResp := SendRequest("POST", "/v1/external/settings/company/addpickup", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) GetAllPickUpLocations() (*GetAllPickUpLocationResponse, *ErrorResponse) {
	var response GetAllPickUpLocationResponse
	errResp := SendRequest("GET", "/v1/external/settings/company/pickup", BASE_URL, s.tokenConfig.Token, nil, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
