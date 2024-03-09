package paypal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal/model_paypal"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/pkg/common"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func ConvertCurrency(currency, value string) string {
	var conversionRate float64
	conversionRate = 1
	if currency == "VND" {
		conversionRate = 0.000043
	}
	valueParse, _ := strconv.ParseFloat(value, 64)
	return fmt.Sprintf("%f", math.Round(valueParse*conversionRate*100)/100)
}
func (p *Paypal) CreateOrder(cart *Cart.Cart) (string, error) {
	p.GetAccessToken("", "")
	var items []model_paypal.Item
	fmt.Println(cart.TotalPrice)

	cart.TotalPrice = ConvertCurrency("VND", cart.TotalPrice)[:5]
	for _, i := range cart.Courses {

		items = append(items, model_paypal.Item{
			Name: i.Title,
			UnitAmount: model_paypal.UnitAmount{
				CurrencyCode: "USD",
				Value:        ConvertCurrency(i.Currency, i.Price)[:5],
			},
		})
	}
	orderData := map[string]interface{}{
		"intent": "CAPTURE",
		"purchase_units": []map[string]interface{}{
			{
				"amount": map[string]interface{}{
					"currency_code": "USD",
					"value":         cart.TotalPrice,
				},
			},
		},
		"application_context": map[string]interface{}{
			"return_url": p.cf.ClientSide.Url + "/cart/paymentsuccess",
			"cancel_url": p.cf.ClientSide.Url + "/cart/paymentfail",
		},
	}

	reqBody, err := json.Marshal(orderData)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	orderCreateURL := p.cf.Paypal.BaseUrl + p.cf.Paypal.CreateOrderApi
	req, err := http.NewRequest("POST", orderCreateURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return "", err

	}
	fmt.Println("token", p.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", p.token.TokenType+" "+p.token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Failed to create order. Status: %d, Body: %s", resp.StatusCode, string(body))
		return "", common.NewCustomError(errors.New("Can not create order"), "Can not create order")
	}
	var orderResponse model_paypal.OrderResponse
	err = json.Unmarshal(body, &orderResponse)
	if err != nil {
		return "", common.NewCustomError(err, "Can not create order")
	}

	var linkApprove string
	// Redirect customer to PayPal payment page
	for _, link := range orderResponse.Links {
		if link.Rel == "approve" {
			linkApprove = link.Href
		}
	}

	return linkApprove, nil
}
