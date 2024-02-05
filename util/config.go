package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver      string `mapstructure:"db_driver"`
	DBSource      string `mapstructure:"db_source"`
	ServerAddress string `mapstructure:"server_address"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml...

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
