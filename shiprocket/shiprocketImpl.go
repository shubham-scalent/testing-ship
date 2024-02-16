package shiprocket

func (s *ShiprocketClient) GetAvailableCouriers(request CourierAvailabityRequest) (*CourierAvailabityResponse, error) {
	// Create a new request
	resp, err := SendRequest("GET", "/v1/external/courier/serviceability/", s.Config.BaseURL, s.TokenConfig.Token, request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response CourierAvailabityResponse
	err = ReadResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ShiprocketClient) CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, error) {
	// Create a new request
	resp, err := SendRequest("GET", "/v1/external/orders/create/adhoc", s.Config.BaseURL, s.TokenConfig.Token, request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response CreateOrderResponse
	err = ReadResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
