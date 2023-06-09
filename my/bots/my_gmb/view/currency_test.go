package view

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	mock "my/bots/my_gmb/view/mock"
	"testing"
	"time"
)

func TestCurrencyView(t *testing.T) {
	suite.Run(t, &CurrencyView{})
}

type CurrencyView struct {
	suite.Suite

	ctrl *gomock.Controller
	ctx  context.Context

	cbr                 *mock.MockCurrencier
	crypto              *mock.MockCurrencier
	currency            *mock.MockCurrencier
	maxCountRetries     int
	timeoutBetweenRetry time.Duration
	logger              *zap.Logger
	view                *Currency
}

func (v *CurrencyView) SetupTest() {
	v.ctrl = gomock.NewController(v.T())
	v.ctx = context.Background()

	v.cbr = mock.NewMockCurrencier(v.ctrl)
	v.currency = mock.NewMockCurrencier(v.ctrl)
	v.crypto = mock.NewMockCurrencier(v.ctrl)

	v.maxCountRetries = 2
	v.timeoutBetweenRetry = time.Millisecond * 100
	v.logger = zap.L()

	v.view = NewCurrency(
		v.cbr,
		v.currency,
		v.crypto,
		v.maxCountRetries,
		v.timeoutBetweenRetry,
		v.logger,
	)
}

func (v *CurrencyView) TestSuccess() {
	expected := "*Курс валют:* \ncurrency\ncrypto"
	v.cbr.EXPECT().Get().Return("cbr", nil)
	v.currency.EXPECT().Get().Return("currency", nil)
	v.crypto.EXPECT().Get().Return("crypto", nil)

	res, err := v.view.View()
	v.Nil(err)
	v.Equal(expected, res)
}

func (v *CurrencyView) TestError() {
	v.Run("error from api currency", func() {
		expected := "*Курс валют:* \ncbr\ncrypto"
		v.cbr.EXPECT().Get().Return("cbr", nil)
		v.currency.EXPECT().Get().Return("", errors.New("error")).Times(v.maxCountRetries)
		v.crypto.EXPECT().Get().Return("crypto", nil)

		res, err := v.view.View()
		v.Nil(err)
		v.Equal(expected, res)
	})

	v.Run("error from api cbr", func() {
		expected := "*Курс валют:* \ncurrency\ncrypto"
		v.cbr.EXPECT().Get().Return("cbr", errors.New("error")).Times(1)
		v.currency.EXPECT().Get().Return("currency", nil)
		v.crypto.EXPECT().Get().Return("crypto", nil)

		res, err := v.view.View()
		v.Nil(err)
		v.Equal(expected, res)
	})

	v.Run("error from all APIs", func() {
		v.cbr.EXPECT().Get().Return("", errors.New("error")).Times(v.maxCountRetries)
		v.currency.EXPECT().Get().Return("", errors.New("error")).Times(v.maxCountRetries)
		v.crypto.EXPECT().Get().Return("", errors.New("error")).Times(v.maxCountRetries)

		res, err := v.view.View()
		v.EqualError(err, "failed to get exchange rate data")
		v.Equal("", res)
	})
}
