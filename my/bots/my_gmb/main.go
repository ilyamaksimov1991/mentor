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

//https://api.openweathermap.org/data/3.0/onecall?lat=55.751244&lon=37.618423&exclude={part}&appid=ba1fb93d3ac2ffc61f4cada5c0e7d5c
//https://api.openweathermap.org/data/3.0/onecall?lat=55.751244&lon=37.618423&appid=ba1fb93d3ac2ffc61f4cada5c0e7d5c
//https://api.openweathermap.org/data/2.5/weather?lat=55.751244&lon=37.618423&appid=6b12d713c6675eeb686d1e76c3012dd3&lang=ru&type=like&units=metric&id=524901
const (
	token  = "5788134511:AAES_-z7AAaQc22huJjXLdJepLjjSKWtR5M"
	chatId = -762680933
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

func main1() {
	//fmt.Print(weather.Weather())
	//os.Exit(1)
	//bot, err := tgbotapi.NewBotAPI(token)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//
	//bot.Debug = true
	//
	//photo := tgbotapi.NewVideo(chatId, tgbotapi.FileURL(fmt.Sprintf("https://cataas.com/cat/gif?%v", time.Now())))
	//if _, err = bot.Send(photo); err != nil {
	//	log.Fatalln(err)
	//	panic(err)
	//}
	//
	//t := quote.NewQuoter()
	//f := fact.NewFact()
	//
	//msg1 := fmt.Sprint(t, f, weather.Weather())
	//msg := tgbotapi.NewMessage(chatId, msg1)
	//msg.ParseMode = "Markdown"
	//
	//_, err = bot.Send(msg)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Print(money.Money())
}
