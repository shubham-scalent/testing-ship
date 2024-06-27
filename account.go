package shiprocket

func (s *ShiprocketClient) GetWalletBalance() (*GetWalletBalanceResponse, *ErrorResponse) {
	var response GetWalletBalanceResponse
	errResp := SendRequest("GET", "/v1/external/account/details/wallet-balance", BASE_URL, s.tokenConfig.Token, nil, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
