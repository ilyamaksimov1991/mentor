package api

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
)

const horoscopeEndpoint = "https://ignio.com/r/export/utf/xml/daily/com.xml"

type HoroscopeResponse struct {
	XMLName xml.Name `xml:"horo"`
	Text    string   `xml:",chardata"`
	Date    struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday,attr"`
		Today      string `xml:"today,attr"`
		Tomorrow   string `xml:"tomorrow,attr"`
		Tomorrow02 string `xml:"tomorrow02,attr"`
	} `xml:"date"`
	Aries struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"aries"`
	Taurus struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"taurus"`
	Gemini struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"gemini"`
	Cancer struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"cancer"`
	Leo struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"leo"`
	Virgo struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"virgo"`
	Libra struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"libra"`
	Scorpio struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"scorpio"`
	Sagittarius struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"sagittarius"`
	Capricorn struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"capricorn"`
	Aquarius struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"aquarius"`
	Pisces struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"pisces"`
}

type Horoscope struct {
}

func NewHoroscope() *Horoscope {
	return &Horoscope{}
}

func (h Horoscope) Get() (string, error) {
	req, err := http.NewRequest("GET", horoscopeEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create money request: %w", err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return "", fmt.Errorf("failed to complete the money request: %w", err)
	}
	defer resp.Body.Close()
	horo := &HoroscopeResponse{}
	derr := xml.NewDecoder(resp.Body)
	derr.CharsetReader = charset.NewReaderLabel
	err = derr.Decode(horo)
	if err != nil {
		return "", fmt.Errorf("curs decoding error: %w", err)
	}

	hr := []string{}
	hr = append(hr, fmt.Sprintf("  ♑️_Козерог_:  %s\n", strings.Trim(horo.Capricorn.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♒️_Водолей_:  %s\n", strings.Trim(horo.Aquarius.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♏️_Скорпион_:  %s\n", strings.Trim(horo.Scorpio.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♊️_Близнецы_:  %s\n", strings.Trim(horo.Gemini.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♋️_Рак_:  %s\n", strings.Trim(horo.Cancer.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♉️_Овен_: %s\n", strings.Trim(horo.Aries.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♉️_Телец_: %s\n", strings.Trim(horo.Taurus.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♌️_Лев_: %s\n", strings.Trim(horo.Leo.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♍️_Дева_: %s\n", strings.Trim(horo.Virgo.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♎️_Весы_: %s\n", strings.Trim(horo.Libra.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♐️_Стрелец_: %s\n", strings.Trim(horo.Sagittarius.Today, "\n")))
	hr = append(hr, fmt.Sprintf("  ♐️_Рыбы_: %s", strings.Trim(horo.Pisces.Today, "\n")))

	return strings.Join(hr, ""), nil
}
