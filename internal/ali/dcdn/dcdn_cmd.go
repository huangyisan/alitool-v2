package dcdn

import (
	"alitool-v2/internal/ali/account"
	"fmt"
)

func ListALLDomainsInfo() {
	ats := account.GetAccountMap()
	for _, at := range ats {
		ListDomainInfoByAccountName(at.AccountName)
	}
}

func ListDomainInfoByAccountName(accountName string) {
	ats := account.GetAccountMap()
	at := ats[accountName]
	fmt.Printf("Account Name: %s\n", at.AccountName)
	client := NewDCDNClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
	domainsInfo, err := listDCDNDomainsResponse(client)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range domainsInfo {
			for _, y := range v.Domains.PageData {
				tmpSource := make([]string, 0)
				for _, source := range y.Sources.Source {
					tmpSource = append(tmpSource, fmt.Sprintf("%s:%d", source.Content, source.Port))
				}
				fmt.Printf("%v, %v\n", y.DomainName, tmpSource)
			}
		}
	}
	fmt.Printf("\n")
}
