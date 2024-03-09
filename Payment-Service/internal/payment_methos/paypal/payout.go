package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal/model_paypal"
	"io/ioutil"
	"net/http"
	"time"
)

func (p *Paypal) Payout(itemsPayout []model_paypal.ItemsPayout) error {

	payoutRequest := model_paypal.Payout{
		SenderBatchHeader: model_paypal.SenderBatchHeader{
			SenderBatchID: "Payouts" + time.Now().Format("20060102150405"),
			EmailSubject:  "You have a payout!",
			EmailMessage:  "You have received a payout! Thanks for using our service!",
		},
		Items: itemsPayout,
	}

	reqBody, err := json.Marshal(payoutRequest)
	if err != nil {
		return err
	}
	payoutUrl := p.cf.Paypal.BaseUrl + p.cf.Paypal.PayoutApi
	req, err := http.NewRequest("POST", payoutUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return err

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", p.token.TokenType+" "+p.token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err

	}

	if resp.StatusCode != http.StatusCreated {
		return err

	}
	fmt.Println(string(body))

	fmt.Println("payout captured successfully")
	return nil
}
