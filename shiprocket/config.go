package shiprocket

import "time"

type ClientConfig struct {
	BaseURL  string
	Email    string
	Password string
}

type tokenConfig struct {
	Token             string
	CreatedOn         time.Time
	NextRefreshOnTime time.Time
}
