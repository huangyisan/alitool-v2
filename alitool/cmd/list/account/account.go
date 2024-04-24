package account

import (
	"alitool-v2/internal/ali/account"
	"alitool-v2/internal/common"
	"github.com/spf13/cobra"
)

func accountAction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		config := common.GetConfig()
		account.ListAccount(config)
	}
}

// account represents the list command
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "List Alibaba Cloud account information",
	Long: `The "account" command displays information about your Alibaba Cloud account.
This includes details such as subaccount, and subscription status.`,
	Run: accountAction(),
}

// Here you will define your flags and configuration settings.

// Cobra supports Persistent Flags which will work for this command
// and all subcommands, e.g.:
// listCmd.PersistentFlags().String("foo", "", "A help for foo")

// Cobra supports local flags which will only run when this command
// is called directly, e.g.:
// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
