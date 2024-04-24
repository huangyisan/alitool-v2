package account

import (
	"fmt"
	"github.com/spf13/viper"
)

var accounts AliAccounts
var accountMap map[string]*AliAccount

type AliAccounts struct {
	AliAccount []AliAccount `mapstructure:"ali_account"`
}

type AliAccount struct {
	// 账号名称
	AccountName string `mapstructure:"accountName"`
	// Access Key
	AccessKeyId string `mapstructure:"accessKeyId"`
	// Secret key
	AccessKeySecret string `mapstructure:"secretKeyId"`
	// SubAccount name
	SubAccount string `mapstructure:"subAccount"`
}

func NewAliAccount(name, accessKey, secretKey, subAccountName string) *AliAccount {
	return &AliAccount{
		AccountName:     name,
		AccessKeyId:     accessKey,
		AccessKeySecret: secretKey,
		SubAccount:      subAccountName,
	}
}

func NewAccountMap(config *viper.Viper) {

}

func GetAccountMap(config *viper.Viper) map[string]*AliAccount {
	err := config.Unmarshal(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	accountMap = make(map[string]*AliAccount)
	for _, account := range accounts.AliAccount {
		accountMap[account.AccountName] = &account
	}
	return accountMap
}

//func GetAccessKeyId(accountName string) string {
//	return AccountMap[accountName].AccessKeyId
//}

func GetAccessKeySecret(account *AliAccount) string {
	return account.AccessKeySecret
}

//func init() {
//	NewAccountMap(common.Config)
//}
