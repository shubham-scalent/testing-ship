package apimodel

type Config struct {
	BaseURL  string
	Email    string
	Password string
}

type CourierAvailabityRequest struct {
	PickupPostcode   int     `json:"pickup_Postcode" binding:"omitempty"`
	DeliveryPostcode int     `json:"delivery_Postcode" binding:"omitempty"`
	OrderID          int     `json:"order_id"`
	Cod              int     `json:"cod" binding:"omitempty"`
	Weight           float64 `json:"weight" binding:"omitempty"`
}
