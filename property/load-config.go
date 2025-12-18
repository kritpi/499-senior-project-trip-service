package property

import (
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path, env string) (*Property, error) {
	v := viper.New()

	// config/{env}.yaml
	v.SetConfigName(env)
	v.SetConfigType("json")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Property
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
