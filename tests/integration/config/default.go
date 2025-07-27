package config

var defaultCfg TestConfig

func init() {
	defaultCfg.TestApiGatewayWsAddr = "ws://localhost:8000/ws"
}
