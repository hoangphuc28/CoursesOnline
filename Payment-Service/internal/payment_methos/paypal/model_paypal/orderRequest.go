package model_paypal

type OrderRequest struct {
	ID            string          `json:"id"`
	Intent        string          `json:"intent"`
	Status        string          `json:"status"`
	PurchaseUnits []PurchaseUnits `json:"purchase_units"`
}
type PurchaseUnits struct {
	ReferenceID string `json:"reference_id"`
	Amount      Amount `json:"amount"`
	Items       []Item `json:"items"`
}
type Item struct {
	Name        string     `json:"name"`
	UnitAmount  UnitAmount `json:"unit_amount"`
	Quantity    string     `json:"quantity"`
	Description string     `json:"description"`
}
type UnitAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}
