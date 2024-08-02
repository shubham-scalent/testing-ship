package shiprocket

func (s *ShiprocketClient) GetDiscrepancyData() (*GetDiscrepancyDataResponse, *ErrorResponse) {
	var response GetDiscrepancyDataResponse

	errResp := SendRequestNil("GET", "/v1/external/billing/discrepancy", BASE_URL, s.tokenConfig.Token, nil, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
