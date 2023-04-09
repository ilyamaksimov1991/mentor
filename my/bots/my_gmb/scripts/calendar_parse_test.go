package scripts

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type Hol struct {
	Name string
	Url  string
}
type HAll struct {
	Day      string
	Holidays []*Hol
}

func TestName2(t *testing.T) {
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

	hall := make([]*HAll, 0, 365)
	doc.Find(".holidays-year .holidays-month-items").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(ix int, td1 *goquery.Selection) {
			//if i > 0 {
			//	return
			//}
			td1.Find("td").Each(func(ids int, td *goquery.Selection) {
				//if i > 0 {
				//	return
				//}

				//fmt.Println( td.Text(), td.Find("a href").Text())
				//s.Find("span").Each(func(idsz int, tdz2 *goquery.Selection) {
				//	//ss, b := tdz.Find("a").Attr("href")
				//	fmt.Println(idsz, tdz2.Text())
				//})
				//s.Find("nowrap").Each(func(idsz int, tdz2 *goquery.Selection) {
				//ss, b := tdz.Find("a").Attr("href")

				h := HAll{}
				_, b := td.Attr("nowrap")
				if b {
					td.Find("span").Not(".holidays-weekday").Each(func(idsz int, tdz2 *goquery.Selection) {
						//ss, b := tdz.Find("a").Attr("href")
						//fmt.Println(idsz, tdz2.Text())
						h.Day = tdz2.Text()
					})
				}
				//})
				s.Find("div").Each(func(idsz int, tdz *goquery.Selection) {
					ss, _ := tdz.Find("a").Attr("href")
					//fmt.Println(idsz, tdz.Text(), ss, b)
					h.Holidays = append(h.Holidays, &Hol{
						Name: tdz.Text(),
						Url:  ss,
					})
				})

				hall = append(hall, &h)
			})
		})
	})

	fmt.Printf("%#v %#v", hall[350].Day, hall[350].Holidays[0])
}

func TestName3(t *testing.T) {
	c := NewCalendarForTheYear()
	cal, err := c.Parse()

	fmt.Printf("%v \n", err)
	fmt.Printf("sdf %v %#v \n", cal[2].Date, cal[2].Holidays[0])
	fmt.Printf("\n\n\n sdf %#v %#v \n", len(cal[2].Holidays), cal[2])
	//assert.Nil(t, err)
	c.SaveToFile("2024_new.json", cal)
}
