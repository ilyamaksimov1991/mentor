package money

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const cryptoEndpoint = "http://api.coinlayer.com/api/live?access_key=2faaf66bafb4c2f38605a2219cd7b9d8&target=USD&symbols=BTC,ETH"

type CryptoResponse struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Target    string `json:"target"`
	Rates     struct {
		Btc float64 `json:"BTC"`
		Eth float64 `json:"ETH"`
	} `json:"rates"`
}

type Crypto struct {
}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (m *Crypto) Get() (string, error) {
	req, err := http.NewRequest("GET", cryptoEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create crypto request: %w", err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return "", fmt.Errorf("failed to complete the money request: %w", err)
	}
	defer resp.Body.Close()
	curs := &CryptoResponse{}
	derr := json.NewDecoder(resp.Body).Decode(curs)
	if derr != nil {
		return "", fmt.Errorf("crypto decoding error: %w", err)
	}

	return fmt.Sprintf(" BTC %.2f \n ETH %.2f \n", curs.Rates.Btc, curs.Rates.Eth), nil
}

// для сум к доллару
//curl --request GET 'https://api.apilayer.com/currency_data/live?source=USD&currencies=uzs,rub' \
//--header 'apikey: VWFCzZbdboGDomwzCPKvs7gtM9Q3r0dU'
//
//live?source={source}&currencies={currencies}
