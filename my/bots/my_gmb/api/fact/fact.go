package fact

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://randstuff.ru/fact/generate/"

type factResponse struct {
	Fact struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"fact"`
}

type Fact struct {
}

func NewFact() *Fact {
	return &Fact{}
}

func (f *Fact) Get() (string, error) {
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create fact request: %w", err)
	}

	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
	r.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", fmt.Errorf("failed to complete the fact request: %w", err)
	}

	defer res.Body.Close()

	response := &factResponse{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		return "", fmt.Errorf("day fact decoding error: %w", err)
	}

	return response.Fact.Text, nil
}
