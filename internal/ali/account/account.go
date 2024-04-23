package account

type AliAccount struct {
	// 账号名称
	Name string
	// Access Key
	AccessKey string
	// Secret key
	SecretKey string
	// SubAccount name
	SubAccountName string
}

func NewAliAccount(name, accessKey, secretKey, subAccountName string) *AliAccount {
	return &AliAccount{
		Name:           name,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		SubAccountName: subAccountName,
	}
}
