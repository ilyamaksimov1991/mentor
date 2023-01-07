package view

import (
	"fmt"
	"strings"
)

type Api interface {
	Get() (string, error)
}

type Goodmorning struct {
	quote   Api
	fact    Api
	money   Api
	weather Api
}

func ViewGoodmorning(
	quote Api,
	fact Api,
	money Api,
	weather Api,
) *Goodmorning {
	return &Goodmorning{
		quote:   quote,
		fact:    fact,
		money:   money,
		weather: weather,
	}
}

func (g *Goodmorning) View() (string, error) {
	result := make([]string, 0, 3)
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
	res, err := g.money.Get()
	if err != nil {
		return "", fmt.Errorf("curs getting error: %w", err)
	}

	return fmt.Sprintf("*Курс валют:* \n%s", res), nil
}

//https://t.me/+SAp7bGMh_lVjMjky
