package dcdn

import (
	"alitool-v2/internal/ali/account"
	"alitool-v2/internal/common"
	"fmt"
)

func ListALLDomainsInfo() {
	ats := account.GetAccountMap(common.Config)
	fmt.Println("ListALLDomainsInfo")
	for _, at := range ats {
		//fmt.Printf("%#v", at)
		ListDomainInfoByAccountName(at.AccountName)
	}
}

func ListDomainInfoByAccountName(accountName string) {
	ats := account.GetAccountMap(common.Config)
	at := ats[accountName]
	client := NewDCDNClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
	domainsInfo, err := listDCDNDomainsResponse(client)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range domainsInfo {
			for _, y := range v.Domains.PageData {
				fmt.Printf("%#v\n", y)
			}
		}
	}
}
