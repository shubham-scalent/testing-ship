package shiprocket

import (
	"fmt"
)

func (s *ShiprocketClient) TrackingThroughOrderID(request TrackingThroughOrderIDRequest) (*TrackingThroughOrderIDResponse, *ErrorResponse) {
	var response TrackingThroughOrderIDResponse
	errResp := SendRequest("GET", "/v1/external/courier/track", BASE_URL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}

func (s *ShiprocketClient) TrackingThroughShipmentID(shipmentID int) (*TrackingThroughShipmentIDResponse, *ErrorResponse) {
	var response TrackingThroughShipmentIDResponse

	endpoint := fmt.Sprintf("/v1/external/courier/track/shipment/%d", shipmentID)
	errResp := SendRequest("GET", endpoint, BASE_URL, s.tokenConfig.Token, nil, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
