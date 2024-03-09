package model_paypal

type Identify struct {
	UserID          string `json:"user_id"`
	Sub             string `json:"sub"`
	Email           string `json:"email"`
	Verified        string `json:"verified"`
	PayerID         string `json:"payer_id"`
	VerifiedAccount string `json:"verified_account"`
	EmailVerified   bool   `json:"email_verified"`
}
