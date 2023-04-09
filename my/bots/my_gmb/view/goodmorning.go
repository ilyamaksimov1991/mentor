package view

import (
	"fmt"
	"strings"
)

type Api interface {
	Get() (string, error)
}

type Goodmorning struct {
	holidays  Api
	advice    Api
	quote     Api
	fact      Api
	money     Api
	crypto    Api
	weather   Api
	horoscope Api
}

func ViewGoodmorning(
	holidays Api,
	advice Api,
	quote Api,
	fact Api,
	money Api,
	crypto Api,
	weather Api,
	horoscope Api,
) *Goodmorning {
	return &Goodmorning{
		holidays:  holidays,
		advice:    advice,
		quote:     quote,
		fact:      fact,
		money:     money,
		crypto:    crypto,
		weather:   weather,
		horoscope: horoscope,
	}
}

func (g *Goodmorning) View() (string, error) {
	result := make([]string, 0, 6)

	holidays2, err := g.holidaysAll()
	if err != nil {
		return "", fmt.Errorf("holidays getting error: %w", err)
	}
	result = append(result, holidays2)

	advice, err := g.adviceForDay()
	if err != nil {
		return "", fmt.Errorf("advice getting error: %w", err)
	}
	result = append(result, advice)

	fact, err := g.factOfTheDay()
	if err != nil {
		return "", fmt.Errorf("fact of the day getting error: %w", err)
	}
	result = append(result, fact)

	quote, err := g.quoter()
	if err != nil {
		return "", fmt.Errorf("quote getting error: %w", err)
	}
	result = append(result, quote)

	horoscope, err := g.horoscope2()
	if err != nil {
		return "", fmt.Errorf("horoscope getting error: %w", err)
	}
	result = append(result, horoscope)

	weather, err := g.weatherOfCities()
	if err != nil {
		return "", fmt.Errorf("weather getting error: %w", err)
	}
	result = append(result, weather)

	money, err := g.money2()
	if err != nil {
		return "", fmt.Errorf("money getting error: %w", err)
	}
	result = append(result, money)

	result = append(result, "Ваш доброе-утро бот ❤️")

	return strings.Join(result, "\n"), nil
}

func (g *Goodmorning) quoter() (string, error) {
	res, err := g.quote.Get()
	if err != nil {
		return "", fmt.Errorf("quote getting error: %w", err)
	}

	return fmt.Sprintf("*Цитата дня:* %s", res), nil
}

func (g *Goodmorning) factOfTheDay() (string, error) {
	res, err := g.fact.Get()
	if err != nil {
		return "", fmt.Errorf("fact getting error: %w", err)
	}

	return fmt.Sprintf("*Факт дня:* %s", res), nil
}

func (g *Goodmorning) weatherOfCities() (string, error) {
	res, err := g.weather.Get()
	if err != nil {
		return "", fmt.Errorf("weather getting error: %w", err)
	}

	return fmt.Sprintf("*Погода:* \n%s", res), nil
}

func (g *Goodmorning) money2() (string, error) {

	crypto, err := g.crypto.Get()
	if err != nil {
		return "", fmt.Errorf("curs getting error: %w", err)
	}

	res, err := g.money.Get()
	if err != nil {
		//return "", fmt.Errorf("curs getting error: %w", err)
		fmt.Printf("curs getting error: %w", err)
		return fmt.Sprintf("*Курс валют:* \n%s", crypto), nil
	}

	return fmt.Sprintf("*Курс валют:* \n%s%s", res, crypto), nil
}
func (g *Goodmorning) horoscope2() (string, error) {
	res, err := g.horoscope.Get()
	if err != nil {
		return "", fmt.Errorf("horoscope getting error: %w", err)
	}

	return fmt.Sprintf("*Гороскоп:* \n%s", res), nil
}

func (g *Goodmorning) holidaysAll() (string, error) {
	res, err := g.holidays.Get()
	if err != nil {
		return "", fmt.Errorf("holidays getting error: %w", err)
	}

	return fmt.Sprintf("*Праздники сегодня:* \n%s", res), nil
}

func (g *Goodmorning) adviceForDay() (string, error) {
	res, err := g.advice.Get()
	if err != nil {
		return "", fmt.Errorf("advice getting error: %w", err)
	}

	return fmt.Sprintf("*Совет дня:* %s", res), nil
}
