package route

// GetUserRequest is a struct that holds the request for GetUser.
type GetUserRequest struct {
	ID string `url:"id"`
}

// SetUserRequest is a struct that holds the request for SetUser.
type SetUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetCoinPriceRequest is a struct that holds the request for GetCoinPrice.
type GetCoinPriceRequest struct {
	Symbol string `url:"symbol"`
}

// SetCoinPriceRequest is a struct that holds the request for SetCoinPrice.
type SetCoinPriceRequest struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
