package cmd

import (
	"fmt"

	"github.com/kolanos/dwolla-v2-go"
	"github.com/spf13/cobra"
)

var sourceUpdateAccountNumber string
var sourceUpdateBankAccountType string
var sourceUpdateName string
var sourceUpdateRoutingNumber string

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Funding source management",
}

var sourceRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove funding source(s)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		for _, arg := range args {
			fmt.Println("Removing funding source:", arg)

			if err := client.FundingSource.Remove(arg); err != nil {
				renderError(err)
			}
		}
	},
}

var sourceRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve funding source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		res, err := client.FundingSource.Retrieve(args[0])

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

var sourceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update funding source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		var bankAccountType dwolla.FundingSourceBankAccountType

		if sourceUpdateBankAccountType != "" {
			switch sourceUpdateBankAccountType {
			case "checking":
				bankAccountType = dwolla.FundingSourceBankAccountTypeChecking
			case "savings":
				bankAccountType = dwolla.FundingSourceBankAccountTypeSavings
			default:
				renderError(fmt.Errorf("Invalid bank account type: %s", accountSourceCreateBankAccountType))
			}
		}

		res, err := client.FundingSource.Update(args[0], &dwolla.FundingSourceRequest{
			RoutingNumber:   sourceUpdateRoutingNumber,
			AccountNumber:   sourceUpdateAccountNumber,
			BankAccountType: bankAccountType,
			Name:            sourceUpdateName,
		})

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)

	sourceCmd.AddCommand(sourceRemoveCmd)
	sourceCmd.AddCommand(sourceRetrieveCmd)
	sourceCmd.AddCommand(sourceUpdateCmd)

	sourceUpdateCmd.Flags().StringVar(&accountSourceCreateAccountNumber, "account-number", "", "bank account number")
	sourceUpdateCmd.Flags().StringVar(&accountSourceCreateBankAccountType, "account-type", "", "bank account type")
	sourceUpdateCmd.Flags().StringVar(&accountSourceCreateName, "name", "", "bank account nickname")
	sourceUpdateCmd.Flags().StringVar(&accountSourceCreateRoutingNumber, "routing-number", "", "bank routing number")
}
