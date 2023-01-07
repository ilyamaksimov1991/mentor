package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	cron "github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"my/bots/my_gmb/api/fact"
	"my/bots/my_gmb/api/money"
	"my/bots/my_gmb/api/quote"
	"my/bots/my_gmb/api/weather"
	"my/bots/my_gmb/config"
	"my/bots/my_gmb/sender"
	"my/bots/my_gmb/view"
	"os"
	"os/signal"
	"syscall"
)

const (
	token  = "5788134511:AAES_-z7AAaQc22huJjXLdJepLjjSKWtR5M"
	chatId = 335693490
	//chatId = -762680933
)

func main() {
	var cfg config.Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal("Failed to parse config.", err)
	}

	logger, err := initLogger(cfg.LogLevel, cfg.LogJSON)
	if err != nil {
		log.Fatal("Failed to init logger.", err)
	}

	logger.Info("config", zap.Any("logger", cfg))

	scheduler := cron.New()
	defer scheduler.Stop()

	s, err := sender.TgSender(token, chatId)
	if err != nil {
		logger.Fatal("tg sender creation error", zap.Error(err))
	}

	viewer := view.ViewGoodmorning(
		quote.NewQuoter(),
		fact.NewFact(),
		money.NewMoney(),
		money.NewCrypto(),
		weather.NewWeather(),
	)

	_, err = scheduler.AddFunc(cfg.GoodmorningBotSchedule, func() {
		logger.Info("start")

		template := ""
		retry := 0
		for retry <= cfg.Retry {
			template, err = viewer.View()
			if err != nil {
				retry++
				logger.Error("error getting view", zap.Error(err))
			} else {
				break
			}
		}

		if template == "" {
			logger.Error("template empty", zap.Error(err))
			return
		}

		err = s.SendVideo()
		if err != nil {
			logger.Error("template sending video error", zap.Error(err))
		}
		err = s.Send(template)
		if err != nil {
			logger.Error("template sending error", zap.Error(err))
		}
	})
	if err != nil {
		logger.Fatal("error starting scheduler", zap.Error(err))
	}

	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

// initLogger создает и настраивает новый экземпляр логгера
func initLogger(logLevel string, isLogJson bool) (*zap.Logger, error) {
	lvl := zap.InfoLevel
	err := lvl.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal log-level: %w", err)
	}
	opts := zap.NewProductionConfig()
	opts.Level = zap.NewAtomicLevelAt(lvl)
	opts.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if opts.InitialFields == nil {
		opts.InitialFields = map[string]interface{}{}
	}

	if !isLogJson {
		opts.Encoding = "console"
		opts.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return opts.Build()
}
