package model_paypal

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
type CaptureOrderResponse struct {
	Status        string `json:"status"`
	PurchaseUnits []struct {
		Payments struct {
			Captures []struct {
				Links []struct {
					Href   string `json:"href"`
					Rel    string `json:"rel"`
					Method string `json:"method"`
				} `json:"links"`
			} `json:"captures"`
		} `json:"payments"`
	} `json:"purchase_units"`
}
