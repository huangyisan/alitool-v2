package dcdn

import (
	"alitool-v2/internal/ali/account"
	"fmt"
	"github.com/spf13/cobra"
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
	domainsInfo, err := listDCDNDomains(client)

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

func ListDcdnSSLCertificatesWithIn30Days() {
	ats := account.GetAccountMap()
	for _, at := range ats {
		fmt.Printf("Account Name: %s\n", at.AccountName)
		client := NewDCDNClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
		res, err := getDCDNSSLCertificateList(client)
		if err != nil {
			cobra.CompError(err.Error())

		} else {
			for _, v := range res {
				for _, y := range v.CertInfos.CertInfo {
					if y.CertStatus == "expire_soon" {
						fmt.Printf("\t%s, %s\n", y.DomainName, y.CertExpireTime)
					}
				}
			}
		}
	}
}

func UpdateDcdnSSLCertificate(accountName, domainName, certName string) {
	ats := account.GetAccountMap()
	at := ats[accountName]
	client := NewDCDNClient("cn-shanghai", at.AccessKeyId, at.AccessKeySecret)
	res, err := updateDcdnSSLCertificate(client, domainName, certName)
	if err != nil {
		cobra.CompError(err.Error())
	} else {
		fmt.Printf("%#v\n", res.BaseResponse.GetHttpStatus())
	}
}
