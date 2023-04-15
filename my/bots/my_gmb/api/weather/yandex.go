package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const yandexEndpoint = "'https://api.weather.yandex.ru/v2/informers?lat=55.75396&lon=37.620393"

//const yandexEndpoint = "'https://api.weather.yandex.ru/v2/informers?lat=55.75396&lon=37.620393' --header 'X-Yandex-API-Key: e840c646-9622-4275-8082-623fb5680d5c'"

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
		ObsTime      int     `json:"obs_time"`
		Uptime       int     `json:"uptime"`
		Temp         int     `json:"temp"`
		FeelsLike    int     `json:"feels_like"`
		Icon         string  `json:"icon"`
		Condition    string  `json:"condition"`
		Cloudness    int     `json:"cloudness"`
		PrecType     int     `json:"prec_type"`
		PrecProb     int     `json:"prec_prob"`
		PrecStrength int     `json:"prec_strength"`
		IsThunder    bool    `json:"is_thunder"`
		WindSpeed    int     `json:"wind_speed"`
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
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"night_short"`
			Night struct {
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
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"night"`
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
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"day_short"`
			Day struct {
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
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"day"`
			Morning struct {
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
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"morning"`
			Evening struct {
				Source       string  `json:"_source"`
				TempMin      int     `json:"temp_min"`
				TempAvg      int     `json:"temp_avg"`
				TempMax      int     `json:"temp_max"`
				WindSpeed    int     `json:"wind_speed"`
				WindGust     float64 `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				SoilTemp     int     `json:"soil_temp"`
				SoilMoisture float64 `json:"soil_moisture"`
				PrecMm       int     `json:"prec_mm"`
				PrecProb     int     `json:"prec_prob"`
				PrecPeriod   int     `json:"prec_period"`
				Cloudness    int     `json:"cloudness"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				UvIndex      int     `json:"uv_index"`
				FeelsLike    int     `json:"feels_like"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				FreshSnowMm  int     `json:"fresh_snow_mm"`
			} `json:"evening"`
		} `json:"parts"`
		Hours []struct {
			Hour         string  `json:"hour"`
			HourTs       int     `json:"hour_ts"`
			Temp         int     `json:"temp"`
			FeelsLike    int     `json:"feels_like"`
			Icon         string  `json:"icon"`
			Condition    string  `json:"condition"`
			Cloudness    int     `json:"cloudness"`
			PrecType     int     `json:"prec_type"`
			PrecStrength int     `json:"prec_strength"`
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
			PrecMm       int     `json:"prec_mm"`
			PrecPeriod   int     `json:"prec_period"`
			PrecProb     int     `json:"prec_prob"`
		} `json:"hours"`
		Biomet struct {
			Index     int    `json:"index"`
			Condition string `json:"condition"`
		} `json:"biomet,omitempty"`
	} `json:"forecasts"`
}

type Yandex struct {
}

func NewYandex() *Yandex {
	return &Yandex{}
}

func (w *Yandex) Get() (string, error) {
	return fmt.Sprint(" Утро: +9☁️ День:+11🌨 Ночь +12☀️\n"), nil
}
func (w *Yandex) Get2() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", yandexEndpoint, nil)
	req.Header.Set("X-Yandex-API-Key", "e840c646-9622-4275-8082-623fb5680d5c")

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	curs := &YandexResponse{}
	derr := json.NewDecoder(res.Body).Decode(curs)
	if derr != nil {
		return "", fmt.Errorf("crypto decoding error: %w", err)
	}
	//
	//return fmt.Sprintf(" USD %.2f \n EUR %.2f \n UZS %.2f \n",
	//	curs.Quotes.Usdrub, curs.Quotes.Usdrub/curs.Quotes.Usdeur, curs.Quotes.Usduzs), nil

	return "", nil
}
