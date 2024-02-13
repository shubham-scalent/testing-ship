package pkg

import (
	"net/http"

	"github.com/shubham-scalent/testing-ship/apimodel"
)

type ShiprockertService interface {
	GetToken(config apimodel.Config) (string, error)
	GetAvailableCouriers(request apimodel.CourierAvailabityRequest, config apimodel.Config, token string) (*apimodel.CourierAvailabityResponse, error)
	GetAvailableCouriers2(request apimodel.CourierAvailabityRequest, config apimodel.Config, token string) (*http.Response, error)
}
