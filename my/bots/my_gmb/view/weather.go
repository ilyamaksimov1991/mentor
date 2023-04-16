package view

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

//go:generate mockgen -source=weather.go -destination=weather_mock_test.go -package=view

type city int
type Coord struct {
	Lon float64
	Lat float64
}

const (
	ufa      city = 0
	dmitrov  city = 1
	tashkent city = 2
)

var cityToCoordsMap = []Coord{
	ufa:      {Lat: 54.809469, Lon: 56.113848}, // долгота longitude
	dmitrov:  {Lat: 56.375150, Lon: 37.531541}, // широта latitude
	tashkent: {Lat: 41.264650, Lon: 69.216270},
}

type Weatherer interface {
	Get(coord Coord) (string, error)
}

type Weather struct {
	yandex         Weatherer
	openweathermap Weatherer
	logger         *zap.Logger
}

func NewWeather(
	yandex Weatherer,
	openweathermap Weatherer,
	logger *zap.Logger,
) *Weather {
	return &Weather{
		yandex:         yandex,
		openweathermap: openweathermap,
		logger:         logger,
	}
}

func (w *Weather) View() (string, error) {
	result := make([]string, 0, 3)
	for _, coord := range cityToCoordsMap {
		yandexWeather, err := w.yandex.Get(coord)
		if err != nil {
			w.logger.Error("", zap.Error(err))
			continue
		}
		result = append(result, yandexWeather)
	}

	if len(result) == 0 {
		for _, coord := range cityToCoordsMap {
			owmWeather, err := w.openweathermap.Get(coord)
			if err != nil {
				w.logger.Error("", zap.Error(err))
				continue
			}
			result = append(result, owmWeather)
		}
	}

	if len(result) == 0 {
		return "", errors.New("failed to get weather data")
	}

	return fmt.Sprintf("*Погода:*\nПо данным Яндекс Погоды\n%s", strings.Join(result, "\n")), nil
}

func (w *Weather) Get() (string, error) {
	return w.View()
}
