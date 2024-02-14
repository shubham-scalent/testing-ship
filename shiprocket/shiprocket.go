package shiprocket

type ShiprockertService interface {
	GetToken(config ClientConfig) (string, error)
	GetAvailableCouriers(request CourierAvailabityRequest, config ClientConfig, token string) (*CourierAvailabityResponse, error)
}
