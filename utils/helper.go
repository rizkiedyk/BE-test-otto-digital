package utils

import (
	"github.com/spf13/viper"
)

func GetEnv(key, defaultValue string) string {
	getEnv := viper.GetString(key)
	if len(getEnv) == 0 || getEnv == "" {
		return defaultValue
	}
	return getEnv

}
