package account

import (
	"fmt"
	"github.com/spf13/viper"
)

func ListAccount(config *viper.Viper) {
	err := config.Unmarshal(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	for _, account := range accounts.AliAccount {
		fmt.Printf("%s --> %s\n", account.AccountName, account.SubAccount)
	}
}
