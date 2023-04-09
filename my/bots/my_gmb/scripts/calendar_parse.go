package scripts

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var year = 2024

var calendarEndpoint = fmt.Sprintf("https://my-calend.ru/holidays/%d", year)

const (
	Января = 1 + iota
	Февраля
	Марта
	Апреля
	Мая
	Июня
	Июля
	Августа
	Сентября
	Октября
	Ноября
	Декабря
)

var monthToIntMap = map[string]int{
	"января":   1,
	"февраля":  2,
	"марта":    3,
	"апреля":   4,
	"мая":      5,
	"июня":     6,
	"июля":     7,
	"августа":  8,
	"сентября": 9,
	"октября":  10,
	"ноября":   11,
	"декабря":  12,
}

type Holiday struct {
	Name string `json:"holiday"`
	Url  string `json:"url"`
}
type Calendar struct {
	Date     string     `json:"date"`
	Holidays []*Holiday `json:"holidays"`
}

type CalendarForTheYear struct {
}

func NewCalendarForTheYear() *CalendarForTheYear {
	return &CalendarForTheYear{}
}

func (c *CalendarForTheYear) Parse() ([]*Calendar, error) {
	res, err := http.Get("https://my-calend.ru/holidays/2023")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	calendars := make([]*Calendar, 0, 365)
	doc.Find(".holidays-year .holidays-month-items").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(ix int, tr *goquery.Selection) {
			calendar := Calendar{}
			tr.Find("td").Each(func(ids int, td *goquery.Selection) {
				_, b := td.Attr("nowrap")
				if b {
					td.Find("span").Not(".holidays-weekday").Each(func(_ int, span *goquery.Selection) {
						calendar.Date = c.date(span.Text())
					})
				}

				td.Find("div").Each(func(_ int, div *goquery.Selection) {
					url, _ := div.Find("a").Attr("href")
					r := regexp.MustCompile("\\s+")
					name := r.ReplaceAllString(strings.TrimSpace(div.Text()), " ")
					if name != "" {
						calendar.Holidays = append(calendar.Holidays, &Holiday{
							Name: name,
							Url:  url,
						})
					}
				})
			})
			calendars = append(calendars, &calendar)
		})
	})

	//fmt.Printf("%#v %#v", calendars[350].Date, calendars[350].Holidays[0])

	return calendars, nil
}
func (c *CalendarForTheYear) SaveToFile(filename string, calendars []*Calendar) error {
	file, _ := json.MarshalIndent(calendars, "", " ")

	_ = ioutil.WriteFile(filename, file, 0777)
	return nil
}

func (c *CalendarForTheYear) date(str string) string {
	date := make([]int, 0, 3)
	for _, d := range strings.Split(str, " ") {
		if d != "" {
			var dn int
			if res, ok := monthToIntMap[d]; ok {
				dn = res
			} else {
				dn, _ = strconv.Atoi(d)
			}
			date = append(date, dn)
		}
	}
	date = append(date, year)
	return fmt.Sprintf("%02d-%02d-%d", date[0], date[1], date[2])
}
