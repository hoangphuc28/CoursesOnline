package paypal

import (
	"bytes"
	"encoding/json"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal/model_paypal"
	"io/ioutil"
	"net/http"
)

type Paypal struct {
	cf    *config.Config
	token *model_paypal.AuthResponse
}

func NewPayPalHandler(cf *config.Config) *Paypal {
	return &Paypal{cf: cf}
}
func (p *Paypal) GetAccessToken(grant_type string, token string) (*model_paypal.AuthResponse, error) {
	authURL := p.cf.Paypal.BaseUrl + p.cf.Paypal.GetAccessToken
	reqBody := bytes.NewBufferString("grant_type=client_credentials")
	if grant_type == "authorization_code" {
		reqBody = bytes.NewBufferString("grant_type=authorization_code&code=" + token)
	}

	req, err := http.NewRequest("POST", authURL, reqBody)
	if err != nil {
		return nil, err

	}

	req.SetBasicAuth(p.cf.Paypal.ClientId, p.cf.Paypal.SecretKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	if resp.StatusCode != http.StatusOK {
		return nil, err

	}

	if grant_type == "authorization_code" {
		var t *model_paypal.AuthResponse
		err = json.Unmarshal(body, &t)
		if err != nil {
			return nil, err

		}
		return t, nil
	}
	err = json.Unmarshal(body, &p.token)
	if err != nil {
		return nil, err
	}
	return p.token, nil
}
