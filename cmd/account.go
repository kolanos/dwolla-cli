package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account Management",
}

var accountPaymentCmd = &cobra.Command{
	Use:   "mass-payment",
	Short: "Account Mass Payment Management",
}

var accountPaymentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Mass Payments",
	Run:   accountPaymentList,
}

var accountRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve Account Information",
	Run:   accountRetrieve,
}

var accountSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Account Funding Source Management",
}

var accountSourceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Account Funding Source",
	Run:   accountSourceCreate,
}

var accountSourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Funding Sources",
	Run:   accountSourceList,
}

var accountTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Account Transfer Management",
}

var accountTransferListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Transfers",
	Run:   accountTransferList,
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(accountPaymentCmd)
	accountCmd.AddCommand(accountRetrieveCmd)
	accountCmd.AddCommand(accountSourceCmd)
	accountCmd.AddCommand(accountTransferCmd)

	accountPaymentCmd.AddCommand(accountPaymentListCmd)

	accountSourceCmd.AddCommand(accountSourceCreateCmd)
	accountSourceCmd.AddCommand(accountSourceListCmd)

	accountTransferCmd.AddCommand(accountTransferListCmd)
}

func accountPaymentList(cmd *cobra.Command, args []string) {}

func accountRetrieve(cmd *cobra.Command, args []string) {
	initClient()

	res, err := client.Account.Retrieve()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	renderResource(res)
}

func accountSourceCreate(cmd *cobra.Command, args []string) {}
func accountSourceList(cmd *cobra.Command, args []string)   {}

func accountTransferList(cmd *cobra.Command, args []string) {}
