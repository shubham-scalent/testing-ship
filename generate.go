package shiprocket

func (s *ShiprocketClient) GenerateInvoice(request GenerateInvoiceRequest) (*GenerateInvoiceResponse, *ErrorResponse) {
	var response GenerateInvoiceResponse
	errResp := SendRequest("POST", "/v1/external/orders/print/invoice", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) GenerateManifest(request GenerateManifestRequest) (*GenerateManifestResponse, *ErrorResponse) {
	var response GenerateManifestResponse
	errResp := SendRequest("POST", "/v1/external/manifests/generate", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) GenerateLabel(request GenerateManifestRequest) (*GenerateManifestResponse, *ErrorResponse) {
	var response GenerateManifestResponse
	errResp := SendRequest("POST", "/v1/external/courier/generate/label", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) GenerateAWB(request GenerateAWBRequest) (*GenerateAWBResponse, *ErrorResponse) {
	var response GenerateAWBResponse
	errResp := SendRequest("POST", "/v1/external/courier/assign/awb", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
