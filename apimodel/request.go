package apimodel

type Config struct {
	BaseURL  string
	Email    string
	Password string
}

type CourierAvailabityRequest struct {
	PickupPostcode   int `json:"pickup_postcode" binding:"omitempty"`
	DeliveryPostcode int `json:"delivery_postcode" binding:"omitempty"`
	// OrderID          int     `json:"order_id" binding:"omitempty"`
	Cod    int     `json:"cod" binding:"omitempty"`
	Weight float64 `json:"weight" binding:"omitempty"`
}
