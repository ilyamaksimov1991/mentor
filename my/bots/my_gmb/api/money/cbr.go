package money

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"strconv"
	"strings"
)

const (
	cbrEndpoint = "http://www.cbr.ru/scripts/XML_daily.asp"
)

var moneyCurrency = map[string]bool{"R01235": true, "R01239": true, "R01717": true}

//type AutoGenerated struct {
//	ValCurs struct {
//		Valute []struct {
//			NumCode  string `xml:"NumCode"`
//			CharCode string `xml:"CharCode"`
//			Nominal  string `xml:"Nominal"`
//			Name     string `xml:"Name"`
//			Value    string `xml:"Value"`
//			ID       string `xml:"_ID"`
//		} `xml:"Valute"`
//		Date string `xml:"_Date"`
//		Name string `xml:"_name"`
//	} `xml:"ValCurs"`
//}

type ValCurrency struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  int    `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

type Cbr struct {
}

func NewCbr() *Cbr {
	return &Cbr{}
}

func (m *Cbr) Get() (string, error) {
	req, err := http.NewRequest("GET", cbrEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create moneyCurrency request: %w", err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return "", fmt.Errorf("failed to complete the moneyCurrency request: %w", err)
	}
	defer resp.Body.Close()
	curs := &ValCurrency{}
	derr := xml.NewDecoder(resp.Body)
	derr.CharsetReader = charset.NewReaderLabel
	err = derr.Decode(curs)
	if err != nil {
		return "", fmt.Errorf("curs decoding error: %w", err)
	}

	str := ""
	for _, val := range curs.Valute {
		if _, ok := moneyCurrency[val.ID]; ok {
			// замена запятой на точку
			cost, err := strconv.ParseFloat(strings.Replace(val.Value, ",", ".", -1), 64)
			if err != nil {
				return "", fmt.Errorf("price conversion error: %w", err)
			}
			if val.Nominal > 1 {
				cost = float64(val.Nominal) / cost
			}
			str += fmt.Sprintf(" %v %.2f\n", val.CharCode, cost)
		}
	}

	return str, nil
}
