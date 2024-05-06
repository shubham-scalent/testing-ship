package shiprocket

func (s *ShiprocketClient) CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, *ErrorResponse) {
	var response CreateOrderResponse
	errResp := SendRequest("POST", "/v1/external/orders/create/adhoc", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) CancelOrder(request CancelOrderRequest) *ErrorResponse {
	errResp := SendRequest("POST", "/v1/external/orders/cancel", BASE_URL, s.tokenConfig.Token, request, nil)
	if errResp != nil {
		return errResp
	}

	return nil
}
