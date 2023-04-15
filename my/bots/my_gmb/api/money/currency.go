package money

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const currencyEndpoint = "https://api.apilayer.com/currency_data/live?source=USD&currencies=uzs,rub,eur"

type CurrencyResponse struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		Usduzs float64 `json:"USDUZS"`
		Usdrub float64 `json:"USDRUB"`
		Usdeur float64 `json:"USDEUR"`
	} `json:"quotes"`
}

type Currency struct {
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (m *Currency) Get() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", currencyEndpoint, nil)
	req.Header.Set("apikey", "VWFCzZbdboGDomwzCPKvs7gtM9Q3r0dU")

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	curs := &CurrencyResponse{}
	derr := json.NewDecoder(res.Body).Decode(curs)
	if derr != nil {
		return "", fmt.Errorf("crypto decoding error: %w", err)
	}

	return fmt.Sprintf(" USD %.2f \n EUR %.2f \n UZS %.2f",
		curs.Quotes.Usdrub, curs.Quotes.Usdrub/curs.Quotes.Usdeur, curs.Quotes.Usduzs), nil
}
