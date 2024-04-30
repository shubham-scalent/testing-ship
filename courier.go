package shiprocket

func (s *ShiprocketClient) GetAvailableCouriers(request CourierAvailabityRequest) (*CourierAvailabityResponse, *ErrorResponse) {
	var response CourierAvailabityResponse
	err := SendRequest("GET", "/v1/external/courier/serviceability/", BASE_URL, s.tokenConfig.Token, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ShiprocketClient) CancelShipments(request CancelShipmentRequest) (*CancelShipmentResponse, *ErrorResponse) {
	var response CancelShipmentResponse
	err := SendRequest("POST", "/v1/external/orders/cancel/shipment/awbs", BASE_URL, s.tokenConfig.Token, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ShiprocketClient) ShipmentPickUp(request ShipmentPickUpRequest) (*ShipmentPickUpResponse, *ErrorResponse) {
	var response ShipmentPickUpResponse
	err := SendRequest("POST", "/v1/external/courier/generate/pickup", BASE_URL, s.tokenConfig.Token, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
