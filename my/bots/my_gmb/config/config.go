package config

// Config конфигурация приложения
type Config struct {
	LogLevel               string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON                bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	GoodmorningBotSchedule string `long:"goodmorning-bot-schedule" description:"Goodmorning bot schedule" env:"CL_GOODMORNING_BOT_SCHEDULE" required:"true"`
	Retry                  int    `long:"retry" description:"retry" env:"CL_RETRY" required:"true"`
}
