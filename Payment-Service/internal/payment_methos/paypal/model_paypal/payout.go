package model_paypal

type Payout struct {
	SenderBatchHeader SenderBatchHeader `json:"sender_batch_header"`
	Items             []ItemsPayout     `json:"items"`
}
type SenderBatchHeader struct {
	SenderBatchID string `json:"sender_batch_id"`
	EmailSubject  string `json:"email_subject"`
	EmailMessage  string `json:"email_message"`
}
type ItemsPayout struct {
	RecipientType        string `json:"recipient_type"`
	Amount               Amount `json:"amount"`
	Note                 string `json:"note"`
	SenderItemID         string `json:"sender_item_id"`
	Receiver             string `json:"receiver"`
	NotificationLanguage string `json:"notification_language,omitempty"`
}
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}
