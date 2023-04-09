package holiday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Holiday struct {
	Name string `json:"holiday"`
	Url  string `json:"url"`
}
type Calendar struct {
	Date     string     `json:"date"`
	Holidays []*Holiday `json:"holidays"`
}

type HolidaysToday struct {
	family map[string][]*Holiday
	world  map[string][]*Holiday
}

func NewHolidaysToday() *HolidaysToday {
	calendarsWorld, err := parseFile("2023.json")
	if err != nil {
		fmt.Printf("%s \n", err)
	}
	calendarsFamily, err := parseFile("family.json")
	if err != nil {
		fmt.Printf("%s \n", err)
	}
	dateToHolidaysMap := make(map[string][]*Holiday, 0)
	for _, calendar := range calendarsWorld {
		dateToHolidaysMap[calendar.Date] = calendar.Holidays
	}

	dateToHolidaysFamilyMap := make(map[string][]*Holiday, 0)
	for _, calendar := range calendarsFamily {
		dateToHolidaysFamilyMap[calendar.Date] = calendar.Holidays
	}

	return &HolidaysToday{
		family: dateToHolidaysFamilyMap,
		world:  dateToHolidaysMap,
	}
}

func (h *HolidaysToday) Get() (string, error) {
	t := time.Now()

	holidaysFamily := h.family[fmt.Sprintf("%02d-%02d", t.Day(), int(t.Month()))]
	holidaysWorld := h.world[fmt.Sprintf("%02d-%02d-%d", t.Day(), int(t.Month()), t.Year())]

	res := make([]string, 0, len(holidaysFamily)+len(holidaysWorld))
	for _, holiday := range holidaysFamily {
		res = append(res, fmt.Sprintf(" 🥳[%s](%s)", holiday.Name, holiday.Url))
	}

	for _, holiday := range holidaysWorld {
		res = append(res, fmt.Sprintf(" [%s](%s)", holiday.Name, holiday.Url))
	}
	return strings.Join(res, "\n"), nil
}

func parseFile(filename string) ([]*Calendar, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file open error: %w", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var calendars []*Calendar
	err = json.Unmarshal(byteValue, &calendars)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return calendars, nil
}
