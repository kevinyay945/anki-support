package helper

import "github.com/spf13/viper"

var Config Configer

//go:generate mockgen -destination=config.mock.go -package=helper -self_package=2023_asset_management/helper . Configer
type Configer interface {
	GoogleApiToken() string
	OpenAIToken() string
}

type config struct {
}

func (m *config) OpenAIToken() string {
	return viper.GetString("OPEN_AI_TOKEN")
}

func (m *config) GoogleApiToken() string {
	return viper.GetString("GOOGLE_API_TOKEN")
}

func newConfig() Configer {
	return &config{}
}

func init() {
	Config = newConfig()
}
