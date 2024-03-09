package model_paypal

type OrderResponse struct {
	ID     string `json:"id"`
	Intent string `json:"intent"`
	Status string `json:"status"`
	Links  []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}
