package view

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type Api interface {
	Get() (string, error)
}

type Goodmorning struct {
	holidays Api
	advice   Api
	quote    Api
	fact     Api
	//cbr           Api
	currency Api
	//crypto        Api
	weather       Api
	horoscope     Api
	maxCountRetry int
	logger        *zap.Logger
}

func ViewGoodmorning(
	holidays Api,
	advice Api,
	quote Api,
	fact Api,
	//cbr Api,
	currency Api,
	//crypto Api,
	weather Api,
	horoscope Api,
	maxCountRetry int,
	logger *zap.Logger,
) *Goodmorning {
	return &Goodmorning{
		holidays: holidays,
		advice:   advice,
		quote:    quote,
		fact:     fact,
		//cbr:           cbr,
		currency: currency,
		//crypto:        crypto,
		weather:       weather,
		horoscope:     horoscope,
		maxCountRetry: maxCountRetry,
		logger:        logger,
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

	exchangeRate, err := g.currency.Get()
	if err != nil {
		return "", fmt.Errorf("exchange rate getting error: %w", err)
	}
	result = append(result, exchangeRate)

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

// exchangeRate возвращает курс валют
//func (g *Goodmorning) exchangeRate() (string, error) {
//	wg := sync.WaitGroup{}
//
//	wg.Add(1)
//	var currencyCbr string
//	go func() {
//		defer wg.Done()
//
//		for retry := 0; retry <= g.maxCountRetry; retry++ {
//			currency, err := g.cbr.Get()
//			if err != nil {
//				g.logger.Error("failed to get the exchange rate from the central bank", zap.Int("retry", retry), zap.Error(err))
//				time.Sleep(time.Second * 20)
//			} else {
//				currencyCbr = currency
//				break
//			}
//		}
//	}()
//
//	wg.Add(1)
//	var currency string
//	go func() {
//		defer wg.Done()
//
//		for retry := 0; retry <= g.maxCountRetry; retry++ {
//			currency, err := g.currency.Get()
//			if err != nil {
//				g.logger.Error("failed to get the exchange rate", zap.Int("retry", retry), zap.Error(err))
//				time.Sleep(time.Second * 20)
//			} else {
//				currency = currency
//				break
//			}
//		}
//	}()
//
//	wg.Wait()
//
//	result := make([]string, 0, 3)
//	if currency != "" {
//		result = append(result, currency)
//	}
//
//	if currency == "" && currencyCbr != "" {
//		result = append(result, currencyCbr)
//	}
//
//	crypto, err := g.crypto.Get()
//	if err != nil {
//		g.logger.Error("failed to get the cryptocurrency exchange rate", zap.Error(err))
//	} else {
//		result = append(result, crypto)
//	}
//
//	return fmt.Sprintf("*Курс валют:* \n%s", result), nil
//}

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
