package dcdn

import (
	"alitool-v2/internal/ali/dcdn"
	"github.com/spf13/cobra"
)

var (
	accountName string
	certName    string
	domainName  string
)

func dcdnAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if accountName != "" && certName != "" {
			dcdn.UpdateDcdnSSLCertificate(accountName, domainName, certName)
			return
		}
		cmd.Help()
	}
}

// account represents the list command
var DcdnCmd = &cobra.Command{
	Use:   "dcdn",
	Short: "List Alibaba Cloud account information",
	Long: `The "account" command displays information about your Alibaba Cloud account.
This includes details such as subaccount, and subscription status.`,
	Run: dcdnAction(),
}

func init() {
	// Here you will define your flags and configuration settings.
	DcdnCmd.Flags().StringVarP(&accountName, "account", "a", "", "specified account name")
	DcdnCmd.Flags().StringVarP(&certName, "certName", "i", "", "certificate id")
	DcdnCmd.Flags().StringVarP(&domainName, "domainName", "d", "", "domain name")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dcdnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dcdnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
