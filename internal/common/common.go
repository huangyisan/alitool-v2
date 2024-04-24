package common

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var Config *viper.Viper

func NewConfig() {

	Config = viper.New()
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search Config in home directory with name ".alitool" (without extension).
	Config.AddConfigPath(home)
	Config.SetConfigType("yaml")
	Config.SetConfigName(".alitool")
	Config.AutomaticEnv()
	Config.SetEnvPrefix("ALITOOL")
	if err := Config.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	} else {
		fmt.Fprintln(os.Stderr, "Using config file:", Config.ConfigFileUsed())
	}
}

func GetConfig() *viper.Viper {
	return Config
}
