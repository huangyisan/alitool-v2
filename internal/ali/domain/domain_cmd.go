package domain

import (
	"alitool-v2/internal/ali/account"
	"alitool-v2/internal/common"
	"fmt"
	"github.com/spf13/cobra"
)

func ListDomainsByAccountName(accountName string) {
	ats := account.GetAccountMap()
	at := ats[accountName]
	client := NewDomainClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
	res, err := getDomainList(client)
	if err != nil {
		cobra.CheckErr(err)
	}
	for _, v := range res {
		fmt.Println(v.DomainName)
	}
}

func ListALLDomains() {
	ats := account.GetAccountMap()
	for _, at := range ats {
		fmt.Printf("AccountName: %s\n", at.AccountName)
		ListDomainsByAccountName(at.AccountName)
	}
}

func FindDomainOwner(domainName string) {
	ats := account.GetAccountMap()
	for _, at := range ats {
		client := NewDomainClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
		res, err := getDomainList(client)
		if err != nil {
			cobra.CheckErr(err)
		}
		for _, v := range res {
			if common.DomainSuffix(domainName) == v.DomainName {
				fmt.Printf("%s in account: %s", domainName, at.AccountName)
				return
			}
		}
	}
	fmt.Println("not found")
}

func FindExpiredDomainsWithinDays(expireDay int) {
	ats := account.GetAccountMap()
	for _, at := range ats {
		client := NewDomainClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
		res, err := getDomainList(client)
		if err != nil {
			cobra.CheckErr(err)
		}
		for _, v := range res {
			if v.ExpirationCurrDateDiff <= expireDay {
				fmt.Printf("%s in account: %s, will expire in %d days\n", v.DomainName, at.AccountName, v.ExpirationCurrDateDiff)
			}
		}
	}
}
