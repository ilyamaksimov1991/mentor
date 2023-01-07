package quote

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	quoteEndpoint = "https://api.forismatic.com/api/1.0/"
)

type Quoter struct {
}

func NewQuoter() *Quoter {
	return &Quoter{}
}

func (g *Quoter) Get() (string, error) {
	req, err := http.NewRequest("GET", quoteEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create quote request: %w", err)
	}

	q := req.URL.Query()
	q.Add("method", "getQuote")
	q.Add("format", "text")
	q.Add("lang", "ru")
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return "", fmt.Errorf("failed to complete the quote request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("can't to read body: %w", err)
	}

	return string(body), nil
}
