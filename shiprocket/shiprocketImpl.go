package shiprocket

func (s *ShiprocketClient) GetAvailableCouriers(request CourierAvailabityRequest) (*CourierAvailabityResponse, *ErrorResponse) {

	var response CourierAvailabityResponse
	err := SendRequest("GET", "/v1/external/courier/serviceability/", s.Config.BaseURL, s.TokenConfig.Token, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ShiprocketClient) CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, *ErrorResponse) {

	var response CreateOrderResponse
	errResp := SendRequest("POST", "/v1/external/orders/create/adhoc", s.Config.BaseURL, s.TokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
