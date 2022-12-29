package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	UserSrvUrl    string `mapstructure:"USER_SRV_URL"`
	MovieSrvUrl 	string `mapstructure:"MOVIE_SRV_URL"`
	WatchedSrvUrl string `mapstructure:"WATCHED_SRV_URL"`
	MaxUploadSize int		 `mapstructure:"MAX_UPLOAD_SIZE"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
