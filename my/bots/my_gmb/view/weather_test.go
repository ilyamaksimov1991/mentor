package view

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"testing"
)

func TestWeatherView(t *testing.T) {
	suite.Run(t, &WeatherView{})
}

type WeatherView struct {
	suite.Suite

	ctrl *gomock.Controller
	ctx  context.Context

	yandex         *MockWeatherer
	openweathermap *MockWeatherer
	logger         *zap.Logger
	view           *Weather
}

func (v *WeatherView) SetupTest() {
	v.ctrl = gomock.NewController(v.T())
	v.ctx = context.Background()

	v.yandex = NewMockWeatherer(v.ctrl)
	v.openweathermap = NewMockWeatherer(v.ctrl)
	v.logger = zap.L()

	v.view = NewWeather(
		v.yandex,
		v.openweathermap,
		v.logger,
	)
}

func (v *WeatherView) TestSuccess() {
	expected := "*Погода:*\nПо данным Яндекс Погоды\nufa\ndmitrov\ntashkent"
	v.yandex.EXPECT().Get(cityToCoordsMap[ufa]).Return("ufa", nil)
	v.yandex.EXPECT().Get(cityToCoordsMap[tashkent]).Return("tashkent", nil)
	v.yandex.EXPECT().Get(cityToCoordsMap[dmitrov]).Return("dmitrov", nil)

	res, err := v.view.View()
	v.Nil(err)
	v.Equal(expected, res)
}

func (v *WeatherView) TestError() {
	v.Run("error from api yandex", func() {
		expected := "*Погода:*\nПо данным Яндекс Погоды\nufa\ndmitrov\ntashkent"
		v.yandex.EXPECT().Get(cityToCoordsMap[ufa]).Return("", errors.New("error"))
		v.yandex.EXPECT().Get(cityToCoordsMap[tashkent]).Return("", errors.New("error"))
		v.yandex.EXPECT().Get(cityToCoordsMap[dmitrov]).Return("", errors.New("error"))

		v.openweathermap.EXPECT().Get(cityToCoordsMap[ufa]).Return("ufa", nil)
		v.openweathermap.EXPECT().Get(cityToCoordsMap[tashkent]).Return("tashkent", nil)
		v.openweathermap.EXPECT().Get(cityToCoordsMap[dmitrov]).Return("dmitrov", nil)

		res, err := v.view.View()
		v.Nil(err)
		v.Equal(expected, res)
	})

	v.Run("error from all api", func() {
		v.yandex.EXPECT().Get(cityToCoordsMap[ufa]).Return("", errors.New("error"))
		v.yandex.EXPECT().Get(cityToCoordsMap[tashkent]).Return("", errors.New("error"))
		v.yandex.EXPECT().Get(cityToCoordsMap[dmitrov]).Return("", errors.New("error"))

		v.openweathermap.EXPECT().Get(cityToCoordsMap[ufa]).Return("", errors.New("error"))
		v.openweathermap.EXPECT().Get(cityToCoordsMap[tashkent]).Return("", errors.New("error"))
		v.openweathermap.EXPECT().Get(cityToCoordsMap[dmitrov]).Return("", errors.New("error"))

		res, err := v.view.View()
		v.EqualError(err, "failed to get weather data")
		v.Equal("", res)
	})
}
