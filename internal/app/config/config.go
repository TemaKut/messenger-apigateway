package config

type Config struct {
	Logger struct {
		Level LoggerLevel
	}
	Server struct {
		Http struct {
			Addr string
		}
	}
	Clients struct {
		UserApi struct {
			Addr string
		}
	}
}

func NewConfig() *Config { // TODO default + parse from env
	return &defaultConfig
}
