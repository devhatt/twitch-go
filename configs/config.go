package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	TwTClientID        string `mapstructure:"TWITCH_CLIENT_ID"`
	TwTTokenURL        string `mapstructure:"TWITCH_TOKEN_URL"`
	DSCWebHook         string `mapstructure:"DISCORD_WEBHOOK"`
	DSCWebHookPetDex   string `mapstructure:"DISCORD_WEBHOOK_PETDEX"`
	DSCWebHookOctopost string `mapstructure:"DISCORD_WEBHOOK_OCTOPOST"`
}

func GetConfig() *config {
	return cfg
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
