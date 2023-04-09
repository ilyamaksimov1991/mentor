package weather

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	weatherEndpoint = "https://api.openweathermap.org/data/2.5/weather/"
)

type city int
type coord struct {
	lon float64
	lat float64
}

const (
	ufa      city = 1
	dmitrov  city = 2
	tashkent city = 3
)

var cityToCoordsMap = map[city]coord{
	ufa:      {lat: 54.809469, lon: 56.113848}, // долгота longitude
	dmitrov:  {lat: 56.375150, lon: 37.531541}, // широта latitude
	tashkent: {lat: 41.264650, lon: 69.216270},
}

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Weather struct {
}

func NewWeather() *Weather {
	return &Weather{}
}

func (w *Weather) Get() (string, error) {
	result := make([]string, 0, len(cityToCoordsMap))
	for _, coord := range cityToCoordsMap {
		res, err := w.weather(coord)
		if err != nil {
			return "", fmt.Errorf("failed to get weather: %w", err)
		}

		result = append(result, res)
	}

	return strings.Join(result, "\n"), nil
}

func (w *Weather) weather(coord coord) (string, error) {
	req, err := http.NewRequest("GET", weatherEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create weater request: %w", err)
	}

	q := req.URL.Query()
	q.Add("lon", strconv.FormatFloat(coord.lon, 'f', 6, 64))
	q.Add("lat", strconv.FormatFloat(coord.lat, 'f', 6, 64))
	q.Add("lang", "ru")
	q.Add("units", "metric")
	q.Add("appid", "6b12d713c6675eeb686d1e76c3012dd3")
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return "", fmt.Errorf("failed to complete the weater request: %w", err)
	}
	defer resp.Body.Close()

	weather := &WeatherResponse{}
	derr := json.NewDecoder(resp.Body).Decode(weather)
	if derr != nil {
		return "", fmt.Errorf("weather decoding error: %w", err)
	}

	return fmt.Sprintf("%v: %v°C %v, ощущается как %v°C", weather.Name, math.Round(weather.Main.Temp), weather.Weather[0].Description, math.Round(weather.Main.FeelsLike)), nil
}
