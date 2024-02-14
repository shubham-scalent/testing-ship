package shiprocket

type CourierAvailabityRequest struct {
	PickupPostcode   int     `json:"pickup_postcode" validate:"omitempty"`
	DeliveryPostcode int     `json:"delivery_postcode" validate:"omitempty"`
	OrderID          int     `json:"order_id" validate:"omitempty"`
	Cod              int     `json:"cod" validate:"omitempty"`
	Weight           float64 `json:"weight" validate:"omitempty"`
}
