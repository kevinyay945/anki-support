package helper

import "github.com/spf13/viper"

var Config Configer

//go:generate mockgen -destination=config.mock.go -typed=true -package=helper -self_package=2023_asset_management/helper . Configer
type Configer interface {
	GoogleApiToken() string
	OpenAIToken() string
	AssetPath() string
}

type config struct {
}

func (m *config) AssetPath() string {
	return viper.GetString("ASSET_PATH")
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
