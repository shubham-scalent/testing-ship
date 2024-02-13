package pkg

import "github.com/shubham-scalent/testing-ship/apimodel"

type ShiprockertService interface {
	GetToken(config apimodel.Config) (string, error)
	GetAvailableCouriers(request apimodel.CourierAvailabityRequest, config apimodel.Config, token string) (string, error)
}
