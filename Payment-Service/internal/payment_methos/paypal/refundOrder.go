package paypal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal/model_paypal"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"io/ioutil"
	"net/http"
)

func (p *Paypal) Refund(linkRefund string) error {
	p.GetAccessToken("", "")
	req, err := http.NewRequest("POST", linkRefund, nil)
	if err != nil {
		fmt.Println(err)
		return common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", p.token.TokenType+" "+p.token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return common.NewCustomError(errors.New("Failed to refund"), "Failed to refund order")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return common.NewCustomError(errors.New("Failed to refund"), "Failed to refund order")

	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println(err)
		return common.NewCustomError(errors.New("Failed to refund"), "Failed to refund order")
	}

	var refundModel model_paypal.Refund
	err = json.Unmarshal(body, &refundModel)
	if err != nil {
		fmt.Println(err)
		return common.NewCustomError(errors.New("Failed to refund"), "Failed to refund order")
	}
	if refundModel.Status != "COMPLETED" {
		return common.NewCustomError(errors.New("Failed to refund"), "Failed to refund order")
	}

	return nil
}
