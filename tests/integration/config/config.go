package config

type TestConfig struct {
	TestApiGatewayWsAddr string
}

func NewTestConfig() *TestConfig { // TODO from envs
	return &defaultCfg
}
