package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"my/bots/my_gmb/view"
	"net/http"
	"strconv"
	"time"
)

//const yandexEndpoint = "https://api.weather.yandex.ru/v2/forecast?lat=56.375150&lon=37.531541&extra=true"
const yandexEndpoint = "https://api.weather.yandex.ru/v2/forecast"

type ConditionId string

type Condition struct {
	Id          ConditionId
	Icon        string
	Description string
}

/*
TODO поправить/добавить
clear — ясно.
partly-cloudy — малооблачно.
cloudy — облачно с прояснениями.
overcast — пасмурно.
drizzle — морось.
light-rain — небольшой дождь.
rain — дождь.
moderate-rain — умеренно сильный дождь.
heavy-rain — сильный дождь.
continuous-heavy-rain — длительный сильный дождь.
showers — ливень.
wet-snow — дождь со снегом.
light-snow — небольшой снег.
snow — снег.
snow-showers — снегопад.
hail — град.
thunderstorm — гроза.
thunderstorm-with-rain — дождь с грозой.
thunderstorm-with-hail — гроза с градом.
*/
var conditionIdToConditionMap = map[ConditionId]Condition{
	"clear": {
		Id:          "clear",
		Description: "ясно",
		Icon:        "☀️",
	},
	"partly-cloudy": {
		Id:          "partly-cloudy",
		Description: "малооблачно",
		Icon:        "🌤",
	},
	"cloudy": {
		Id:          "cloudy",
		Description: "облачно с прояснениями",
		Icon:        "🌥",
	},
	"overcast": {
		Id:          "overcast",
		Description: "пасмурно",
		Icon:        "☁️",
	},
	"partly-cloudy-and-light-rain": {
		Id:          "partly-cloudy-and-light-rain",
		Description: "малооблачно, небольшой дождь",
		Icon:        "🌦",
	},
	"light-rain": {
		Id:          "light-rain",
		Description: "небольшой дождь",
		Icon:        "🌦",
	},
	"partly-cloudy-and-rain": {
		Id:          "partly-cloudy-and-rain",
		Description: "малооблачно, дождь",
		Icon:        "🌦",
	},
	"overcast-and-rain": {
		Id:          "overcast-and-rain",
		Description: "значительная облачность, сильный дождь",
	},
	"overcast-thunderstorms-with-rain": {
		Id:          "overcast-thunderstorms-with-rain",
		Description: "сильный дождь с грозой",
		Icon:        "⛈",
	},
	"cloudy-and-light-rain": {
		Id:          "cloudy-and-light-rain",
		Description: "облачно, небольшой дождь",
		Icon:        "🌧", // ?
	},
	"overcast-and-light-rain": {
		Id:          "overcast-and-light-rain",
		Description: "значительная облачность, небольшой дождь",
		Icon:        "🌧", // ?
	},
	"cloudy-and-rain": {
		Id:          "cloudy-and-rain",
		Description: "облачно, дождь",
		Icon:        "🌧",
	},
	"rain": {
		Id:          "rain",
		Description: "дождь",
		Icon:        "🌧",
	},
	"overcast-and-wet-snow": {
		Id:          "overcast-and-wet-snow",
		Description: "дождь со снегом",
		Icon:        "🌨",
	},
	"partly-cloudy-and-light-snow": {
		Id:          "partly-cloudy-and-light-snow",
		Description: "небольшой снег",
	},
	"partly-cloudy-and-snow": {
		Id:          "partly-cloudy-and-snow",
		Description: "малооблачно, снег",
	},
	"overcast-and-snow": {
		Id:          "overcast-and-snow",
		Description: "снегопад",
	},
	"cloudy-and-light-snow": {
		Id:          "cloudy-and-light-snow",
		Description: "облачно, небольшой снег",
	},
	"overcast-and-light-snow": {
		Id:          "overcast-and-light-snow",
		Description: "значительная облачность, небольшой снег",
	},
	"cloudy-and-snow": {
		Id:          "cloudy-and-snow",
		Description: "облачно, снег",
	},
}

type YandexResponse struct {
	Now   int       `json:"now"`
	NowDt time.Time `json:"now_dt"`
	Info  struct {
		N      bool    `json:"n"`
		Geoid  int     `json:"geoid"`
		URL    string  `json:"url"`
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Tzinfo struct {
			Name   string `json:"name"`
			Abbr   string `json:"abbr"`
			Dst    bool   `json:"dst"`
			Offset int    `json:"offset"`
		} `json:"tzinfo"`
		DefPressureMm int    `json:"def_pressure_mm"`
		DefPressurePa int    `json:"def_pressure_pa"`
		Slug          string `json:"slug"`
		Zoom          int    `json:"zoom"`
		Nr            bool   `json:"nr"`
		Ns            bool   `json:"ns"`
		Nsr           bool   `json:"nsr"`
		P             bool   `json:"p"`
		F             bool   `json:"f"`
		H             bool   `json:"_h"`
	} `json:"info"`
	GeoObject struct {
		District struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"district"`
		Locality struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"locality"`
		Province struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"province"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"geo_object"`
	Yesterday struct {
		Temp int `json:"temp"`
	} `json:"yesterday"`
	Fact struct {
		ObsTime   int     `json:"obs_time"`
		Uptime    int     `json:"uptime"`
		Temp      int     `json:"temp"`
		FeelsLike int     `json:"feels_like"`
		Icon      string  `json:"icon"`
		Condition string  `json:"condition"`
		Cloudness float64 `json:"cloudness"`
		PrecType  int     `json:"prec_type"`
		PrecProb  int     `json:"prec_prob"`
		//PrecStrength int     `json:"prec_strength"`
		IsThunder    bool    `json:"is_thunder"`
		WindSpeed    float64 `json:"wind_speed"`
		WindDir      string  `json:"wind_dir"`
		PressureMm   int     `json:"pressure_mm"`
		PressurePa   int     `json:"pressure_pa"`
		Humidity     int     `json:"humidity"`
		Daytime      string  `json:"daytime"`
		Polar        bool    `json:"polar"`
		Season       string  `json:"season"`
		Source       string  `json:"source"`
		SoilMoisture float64 `json:"soil_moisture"`
		SoilTemp     int     `json:"soil_temp"`
		UvIndex      int     `json:"uv_index"`
		WindGust     float64 `json:"wind_gust"`
	} `json:"fact"`
	Forecasts []struct {
		Date      string `json:"date"`
		DateTs    int    `json:"date_ts"`
		Week      int    `json:"week"`
		Sunrise   string `json:"sunrise"`
		Sunset    string `json:"sunset"`
		RiseBegin string `json:"rise_begin"`
		SetEnd    string `json:"set_end"`
		MoonCode  int    `json:"moon_code"`
		MoonText  string `json:"moon_text"`
		Parts     struct {
			NightShort struct {
				Source       string  `json:"_source"`
				Temp         int     `json:"temp"`
				WindSpeed    float64 `json:"wind_speed"`
				WindGust     float64 `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				SoilTemp     int     `json:"soil_temp"`
				SoilMoisture float64 `json:"soil_moisture"`
				//PrecMm       int     `json:"prec_mm"`
				PrecProb   int     `json:"prec_prob"`
				PrecPeriod int     `json:"prec_period"`
				Cloudness  float64 `json:"cloudness"`
				PrecType   int     `json:"prec_type"`
				//PrecStrength int     `json:"prec_strength"`
				Icon        string `json:"icon"`
				Condition   string `json:"condition"`
				UvIndex     int    `json:"uv_index"`
				FeelsLike   int    `json:"feels_like"`
				Daytime     string `json:"daytime"`
				Polar       bool   `json:"polar"`
				FreshSnowMm int    `json:"fresh_snow_mm"`
			} `json:"night_short"`
			Night    Part `json:"night"`
			DayShort struct {
				Source       string  `json:"_source"`
				Temp         int     `json:"temp"`
				TempMin      int     `json:"temp_min"`
				WindSpeed    float64 `json:"wind_speed"`
				WindGust     float64 `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				SoilTemp     int     `json:"soil_temp"`
				SoilMoisture float64 `json:"soil_moisture"`
				//PrecMm       int     `json:"prec_mm"`
				PrecProb   int     `json:"prec_prob"`
				PrecPeriod int     `json:"prec_period"`
				Cloudness  float64 `json:"cloudness"`
				PrecType   int     `json:"prec_type"`
				//PrecStrength int     `json:"prec_strength"`
				Icon        string `json:"icon"`
				Condition   string `json:"condition"`
				UvIndex     int    `json:"uv_index"`
				FeelsLike   int    `json:"feels_like"`
				Daytime     string `json:"daytime"`
				Polar       bool   `json:"polar"`
				FreshSnowMm int    `json:"fresh_snow_mm"`
			} `json:"day_short"`
			Day     Part `json:"day"`
			Morning Part `json:"morning"`
			Evening Part `json:"evening"`
		} `json:"parts"`
		Hours []struct {
			Hour      string  `json:"hour"`
			HourTs    int     `json:"hour_ts"`
			Temp      int     `json:"temp"`
			FeelsLike int     `json:"feels_like"`
			Icon      string  `json:"icon"`
			Condition string  `json:"condition"`
			Cloudness float64 `json:"cloudness"`
			PrecType  int     `json:"prec_type"`
			//PrecStrength int     `json:"prec_strength"`
			IsThunder    bool    `json:"is_thunder"`
			WindDir      string  `json:"wind_dir"`
			WindSpeed    float64 `json:"wind_speed"`
			WindGust     float64 `json:"wind_gust"`
			PressureMm   int     `json:"pressure_mm"`
			PressurePa   int     `json:"pressure_pa"`
			Humidity     int     `json:"humidity"`
			UvIndex      int     `json:"uv_index"`
			SoilTemp     int     `json:"soil_temp"`
			SoilMoisture float64 `json:"soil_moisture"`
			//PrecMm       int     `json:"prec_mm"`
			PrecPeriod int `json:"prec_period"`
			PrecProb   int `json:"prec_prob"`
		} `json:"hours"`
		Biomet struct {
			Index     int    `json:"index"`
			Condition string `json:"condition"`
		} `json:"biomet,omitempty"`
	} `json:"forecasts"`
}

type Part struct {
	Source       string  `json:"_source"`
	TempMin      int     `json:"temp_min"`
	TempAvg      int     `json:"temp_avg"`
	TempMax      int     `json:"temp_max"`
	WindSpeed    float64 `json:"wind_speed"`
	WindGust     float64 `json:"wind_gust"`
	WindDir      string  `json:"wind_dir"`
	PressureMm   int     `json:"pressure_mm"`
	PressurePa   int     `json:"pressure_pa"`
	Humidity     int     `json:"humidity"`
	SoilTemp     int     `json:"soil_temp"`
	SoilMoisture float64 `json:"soil_moisture"`
	//PrecMm       int         `json:"prec_mm"`
	PrecProb   int     `json:"prec_prob"`
	PrecPeriod int     `json:"prec_period"`
	Cloudness  float64 `json:"cloudness"`
	PrecType   int     `json:"prec_type"`
	//PrecStrength int         `json:"prec_strength"`
	Icon        string      `json:"icon"`
	Condition   ConditionId `json:"condition"`
	UvIndex     int         `json:"uv_index"`
	FeelsLike   int         `json:"feels_like"`
	Daytime     string      `json:"daytime"`
	Polar       bool        `json:"polar"`
	FreshSnowMm int         `json:"fresh_snow_mm"`
}

type Yandex struct {
}

func NewYandex() *Yandex {
	return &Yandex{}
}

func (w *Yandex) Get(coord view.Coord) (string, error) {
	req, err := http.NewRequest("GET", yandexEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new get request: %w", err)
	}
	req.Header.Set("X-Yandex-API-Key", "de011669-83cc-425b-b32f-e467e0cdd46b")

	q := req.URL.Query()
	q.Add("lon", strconv.FormatFloat(coord.Lon, 'f', 6, 64))
	q.Add("lat", strconv.FormatFloat(coord.Lat, 'f', 6, 64))
	//q.Add("limit", "3")
	q.Add("extra", "true")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("get request failed: %w", err)
	}
	defer res.Body.Close()

	weather := &YandexResponse{}
	err = json.NewDecoder(res.Body).Decode(weather)
	if err != nil {
		return "", fmt.Errorf("yandex weather decoding error: %w", err)
	}

	if len(weather.Forecasts) == 0 {
		return "", errors.New("missing forecast")
	}
	day := 0
	tempMinMorning := weather.Forecasts[day].Parts.Morning.TempMin
	tempMinDay := weather.Forecasts[day].Parts.Day.TempMin
	tempMinEvening := weather.Forecasts[day].Parts.Evening.TempMin
	tempMinNight := weather.Forecasts[day+1].Parts.Night.TempMin

	tempMaxMorning := weather.Forecasts[day].Parts.Morning.TempMax
	tempMaxDay := weather.Forecasts[day].Parts.Day.TempMax
	tempMaxEvening := weather.Forecasts[day].Parts.Evening.TempMax
	tempMaxNight := weather.Forecasts[day+1].Parts.Night.TempMax

	iconMorning := conditionIdToConditionMap[weather.Forecasts[day].Parts.Morning.Condition].Icon
	iconDay := conditionIdToConditionMap[weather.Forecasts[day].Parts.Day.Condition].Icon
	iconEvening := conditionIdToConditionMap[weather.Forecasts[day].Parts.Evening.Condition].Icon
	iconNight := conditionIdToConditionMap[weather.Forecasts[day+1].Parts.Night.Condition].Icon

	conditionMorning := conditionIdToConditionMap[weather.Forecasts[day].Parts.Morning.Condition]
	conditionDay := conditionIdToConditionMap[weather.Forecasts[day].Parts.Day.Condition]
	conditionEvening := conditionIdToConditionMap[weather.Forecasts[day].Parts.Evening.Condition]
	conditionNight := conditionIdToConditionMap[weather.Forecasts[day+1].Parts.Night.Condition]

	return fmt.Sprintf(
		"[%s](%s): \n"+
			"_Утро_: %s %s (%s); \n"+
			"_День_: %s %s (%s); \n"+
			"_Вечер_: %s %s (%s); \n"+
			"_Ночь_: %s %s (%s);",
		weather.GeoObject.Locality.Name, weather.Info.URL,
		addRange(tempMinMorning, tempMaxMorning), iconMorning, conditionMorning.Description, //humidityMorning,
		addRange(tempMinDay, tempMaxDay), iconDay, conditionDay.Description, //humidityDay,
		addRange(tempMinEvening, tempMaxEvening), iconEvening, conditionEvening.Description, //humidityEvening,
		addRange(tempMinNight, tempMaxNight), iconNight, conditionNight.Description, //humidityNight,
	), nil
}

func addRange(from int, to int) string {
	var signFrom string
	var signTo string
	if from > 0 {
		signFrom = "+"
	}
	if to > 0 {
		signTo = "+"
	}
	return fmt.Sprintf("%s%d°...%s%d°", signFrom, from, signTo, to)
}
