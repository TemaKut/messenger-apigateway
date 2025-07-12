package config

type LoggerLevel int

const (
	LoggerLevelInvalid LoggerLevel = iota
	LoggerLevelDebug
	LoggerLevelInfo
)
