package paypal

import (
	"encoding/json"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Payment-Service/internal/payment_methos/paypal/model_paypal"
	"io/ioutil"
	"log"
	"net/http"
)

func (p *Paypal) GetUserInfor(token *model_paypal.AuthResponse) (*model_paypal.Identify, error) {
	var response *model_paypal.Identify
	req, err := http.NewRequest("GET", p.cf.Paypal.BaseUrl+p.cf.Paypal.IdentifyApi, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get access token. Status: %d, Body: %s", resp.StatusCode, string(body))
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
	return response, nil
}
