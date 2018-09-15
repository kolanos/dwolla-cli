package cmd

import (
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Dwolla Account Management",
	Long:  "",
}

var accountRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve Dwolla Account Information",
	Long:  "",
	Run:   accountRetrieve,
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(accountRetrieveCmd)
}

func accountRetrieve(cmd *cobra.Command, args []string) {
	initClient()

	res, err := client.Account.Retrieve()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(res)
}
