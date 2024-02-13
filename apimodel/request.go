package apimodel

type Config struct {
	BaseURL  string
	Email    string
	Password string
}

type CourierAvailabityRequest struct {
	PickupPostcode   int     `json:"pickupPostcode" binding:"omitempty"`
	DeliveryPostcode int     `json:"deliveryPostcode" binding:"omitempty"`
	OrderID          int     `json:"orderID"`
	Cod              int     `json:"cod" binding:"omitempty"`
	Weight           float64 `json:"weight" binding:"omitempty"`
}
