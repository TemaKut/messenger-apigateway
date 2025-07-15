package config

var defaultConfig Config

func init() {
	defaultConfig.Logger.Level = LoggerLevelDebug
	defaultConfig.Server.Http.Addr = ":8000"
}
