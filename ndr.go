package shiprocket

func (s *ShiprocketClient) GetAllNDRShipments(request GetAllNDRShipmentsRequest) (*GetAllNDRShipmentsResponse, *ErrorResponse) {
	var response GetAllNDRShipmentsResponse
	requestBody := map[string]interface{}{
		"page":     request.Page,
		"per_page": request.PerPage,
		"to":       request.To,
		"from":     request.From,
		"search":   request.SearchAWB}
	err := SendRequestWithParams("GET", "/v1/external/ndr/all", BASE_URL, s.tokenConfig.Token, requestBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ShiprocketClient) GetSpecificNDRShipmentDetails(request GetSpecificNDRShipmentDetailsRequest) (*GetSpecificNDRShipmentDetailsResponse, *ErrorResponse) {
	var response GetSpecificNDRShipmentDetailsResponse
	apiPath := "/v1/external/ndr/" + request.AWBNo
	errResp := SendRequestNil("GET", apiPath, BASE_URL, s.tokenConfig.Token, nil, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) NDRAction(request NDRActionRequest, awb string) (*NDRActionResponse, *ErrorResponse) {
	var response NDRActionResponse
	apiPath := "/v1/external/ndr/" + awb + "/action"
	errResp := SendRequest("POST", apiPath, BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
