package viperconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadViperConfig(dir string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper config not found", err)
	}
}
