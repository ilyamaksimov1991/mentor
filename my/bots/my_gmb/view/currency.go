package view

//go:generate mockgen -source=currency.go -destination=mock/mock_currency.go

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

type Currencier interface {
	Get() (string, error)
}

type Currency struct {
	cbr                 Currencier
	currency            Currencier
	crypto              Currencier
	maxCountRetry       int
	timeoutBetweenRetry time.Duration
	logger              *zap.Logger
}

func NewCurrency(
	cbr Currencier,
	currency Currencier,
	crypto Currencier,
	maxCountRetry int,
	timeoutBetweenRetry time.Duration,
	logger *zap.Logger,
) *Currency {
	return &Currency{
		cbr:                 cbr,
		currency:            currency,
		crypto:              crypto,
		maxCountRetry:       maxCountRetry,
		timeoutBetweenRetry: timeoutBetweenRetry,
		logger:              logger,
	}
}

func (g *Currency) View() (string, error) {
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	var currencyCbr string
	var err error
	go func() {
		defer wg.Done()

		for retry := 1; retry <= g.maxCountRetry; retry++ {
			if ctx.Err() != nil {
				return
			}
			currencyCbr, err = g.cbr.Get()
			if err != nil {
				g.logger.Error("failed to get the exchange rate from the central bank", zap.Int("retry", retry), zap.Error(err))
				time.Sleep(g.timeoutBetweenRetry)
			} else {
				break
			}
		}
	}()

	wg.Add(1)
	var currency string
	go func() {
		defer wg.Done()

		for retry := 1; retry <= g.maxCountRetry; retry++ {
			currency, err = g.currency.Get()
			if err != nil {
				g.logger.Error("failed to get the exchange rate", zap.Int("retry", retry), zap.Error(err))
				time.Sleep(g.timeoutBetweenRetry)
			} else {
				cancel()
				break
			}
		}
	}()

	wg.Add(1)
	var cryptocurrencies string
	go func() {
		defer wg.Done()

		for retry := 1; retry <= g.maxCountRetry; retry++ {
			cryptocurrencies, err = g.crypto.Get()
			if err != nil {
				g.logger.Error("failed to get the exchange rate", zap.Int("retry", retry), zap.Error(err))
				time.Sleep(g.timeoutBetweenRetry)
			} else {
				break
			}
		}
	}()

	wg.Wait()

	result := make([]string, 0, 3)
	if currency != "" {
		result = append(result, currency)
	}

	if currency == "" && currencyCbr != "" {
		result = append(result, currencyCbr)
	}

	if cryptocurrencies != "" {
		result = append(result, cryptocurrencies)
	}

	if len(result) == 0 {
		return "", errors.New("failed to get exchange rate data")
	}

	return fmt.Sprintf("*Курс валют:* \n%s", strings.Join(result, "\n")), nil
}

func (g *Currency) Get() (string, error) {
	return g.View()
}
