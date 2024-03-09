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

func (p *Paypal) CaptureOrder(orderID string) (string, error) {
	p.GetAccessToken("", "")

	orderCaptureURL := p.cf.Paypal.BaseUrl + p.cf.Paypal.CaptureOrderApi
	captureURL := fmt.Sprintf(orderCaptureURL, orderID)

	req, err := http.NewRequest("POST", captureURL, nil)
	if err != nil {
		fmt.Println(err)
		return "", common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", p.token.TokenType+" "+p.token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return "", common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")

	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println(err)
		return "", common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")
	}

	var captureResponse model_paypal.CaptureOrderResponse
	err = json.Unmarshal(body, &captureResponse)
	if err != nil {
		fmt.Println(err)

		return "", common.NewCustomError(errors.New("Failed to capture"), "Failed to capture order")
	}
	var linkRefund string

	for _, i := range captureResponse.PurchaseUnits {
		for _, j := range i.Payments.Captures {
			for _, link := range j.Links {
				if link.Rel == "refund" {
					linkRefund = link.Href
				}
			}
		}

	}

	return linkRefund, nil
}
