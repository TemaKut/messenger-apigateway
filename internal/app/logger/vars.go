package logger

type LogLevel int

const (
	LogLevelInvalid LogLevel = iota
	LogLevelDebug
	LogLevelInfo
)
