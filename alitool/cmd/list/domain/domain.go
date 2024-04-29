package domain

import (
	"alitool-v2/internal/ali/domain"
	"github.com/spf13/cobra"
)

var (
	accountName string
	regionId    string
	domainName  string
	isReverse   bool
	isExpire    bool
	expireDay   int
)

func domainAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if isExpire == false {
			if accountName != "" && domainName == "" && isReverse == false {
				domain.ListDomainsByAccountName(accountName)
				return
			}
			if accountName == "" && domainName == "" && isReverse == false {
				domain.ListALLDomains()
				return
			}
			if domainName != "" && isReverse == true {
				domain.FindDomainOwner(domainName)
				return
			}
		}
		if isExpire {
			domain.FindExpiredDomainsWithinDays(expireDay)
			return
		}
		cmd.Help()

	}
}

// domainCmd represents the domain command
var DomainCmd = &cobra.Command{
	Use:                   "domain",
	Short:                 "list all domains",
	DisableFlagsInUseLine: true,
	Example: `  # List all domains
  	alitool list domain`,
	Run: domainAction(),
}

func init() {

	DomainCmd.Flags().StringVarP(&domainName, "domain", "i", "", "specific domain name")
	DomainCmd.Flags().StringVarP(&accountName, "account", "a", "", "specific account name")
	DomainCmd.Flags().StringVarP(&regionId, "regionId", "z", "", "specific region id")
	DomainCmd.Flags().BoolVarP(&isReverse, "isReverse", "R", false, "isReverse the domain in account")
	DomainCmd.Flags().BoolVarP(&isExpire, "isExpire", "E", false, "isExpire the domain in account")
	DomainCmd.Flags().IntVarP(&expireDay, "expireDay", "e", 30, "domain expireDay")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
