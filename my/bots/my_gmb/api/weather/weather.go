package weather

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//https://api.openweathermap.org/data/2.5/weather?lat=55.751244&lon=37.618423&appid=6b12d713c6675eeb686d1e76c3012dd3&lang=ru&type=like&units=metric&id=524901
//http://api.coinlayer.com/api/live?access_key=2faaf66bafb4c2f38605a2219cd7b9d8&target=USD&symbols=BTC,ETH
const (
	weatherEndpoint = "https://api.openweathermap.org/data/2.5/weather/"
)

type city string
type coord struct {
	lon float64
	lat float64
}

const (
	ufa     city = "ufa"
	dmitrov city = "dmitrov"
)

var cityToCoordsMap = map[city]coord{
	ufa:     {lon: 56.126128, lat: 54.815160},
	dmitrov: {lon: 37.531541, lat: 56.375150},
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
