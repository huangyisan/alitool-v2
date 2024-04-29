package common

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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

// DomainSuffix will return domain suffix, such as www.baidu.com will return baidu.com
func DomainSuffix(domainName string) string {
	dn := strings.Split(domainName, ".")
	if len(dn) == 1 {
		return strings.Join(dn, "")
	}
	return strings.Join(dn[len(dn)-2:], ".")
}
