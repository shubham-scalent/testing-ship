package shiprocket

func (s *ShiprocketClient) GetLocalityDetails(request GetLocalityDetailsRequest) (*GetLocalityDetailsResponse, *ErrorResponse) {
	var response GetLocalityDetailsResponse
	errResp := SendRequest("GET", "/v1/external/open/postcode/details", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
