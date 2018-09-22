package cmd

import (
	"fmt"

	"github.com/kolanos/dwolla-v2-go"
	"github.com/spf13/cobra"
)

var sourceDepositVerifyAmount1 string
var sourceDepositVerifyAmount2 string
var sourceDepositVerifyCurrency string

var sourceUpdateAccountNumber string
var sourceUpdateBankAccountType string
var sourceUpdateName string
var sourceUpdateRoutingNumber string

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Funding source management",
}

var sourceBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Balance management",
}

var sourceBalanceRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve balance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		src, err := client.FundingSource.Retrieve(args[0])

		if err != nil {
			renderError(err)
		}

		res, err := src.RetrieveBalance()

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

var sourceDepositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Micro deposit management",
}

var sourceDepositInitiateCmd = &cobra.Command{
	Use:   "initiate",
	Short: "Initiate micro deposits",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		src, err := client.FundingSource.Retrieve(args[0])

		if err != nil {
			renderError(err)
		}

		res, err := src.InitiateMicroDeposits()

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

var sourceDepositRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve micro deposits",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		src, err := client.FundingSource.Retrieve(args[0])

		if err != nil {
			renderError(err)
		}

		res, err := src.RetrieveMicroDeposits()

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

var sourceDepositVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify micro deposits",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		src, err := client.FundingSource.Retrieve(args[0])

		if err != nil {
			renderError(err)
		}

		var currency dwolla.Currency
		switch sourceDepositVerifyCurrency {
		case "USD":
			currency = dwolla.USD
		default:
			renderError(fmt.Errorf("Invalid currency: %s", sourceDepositVerifyCurrency))
		}

		err = src.VerifyMicroDeposits(&dwolla.MicroDepositRequest{
			Amount1: dwolla.Amount{
				Value:    sourceDepositVerifyAmount1,
				Currency: currency,
			},
			Amount2: dwolla.Amount{
				Value:    sourceDepositVerifyAmount2,
				Currency: currency,
			},
		})

		if err != nil {
			renderError(err)
		}
	},
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
				renderError(fmt.Errorf("Invalid bank account type: %s", sourceUpdateBankAccountType))
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

	sourceCmd.AddCommand(sourceBalanceCmd)
	sourceCmd.AddCommand(sourceDepositCmd)
	sourceCmd.AddCommand(sourceRemoveCmd)
	sourceCmd.AddCommand(sourceRetrieveCmd)
	sourceCmd.AddCommand(sourceUpdateCmd)

	sourceBalanceCmd.AddCommand(sourceBalanceRetrieveCmd)

	sourceDepositCmd.AddCommand(sourceDepositInitiateCmd)
	sourceDepositCmd.AddCommand(sourceDepositRetrieveCmd)
	sourceDepositCmd.AddCommand(sourceDepositVerifyCmd)

	sourceDepositVerifyCmd.Flags().StringVar(&sourceDepositVerifyAmount1, "amount1", "", "first micro deposit amount (required)")
	sourceDepositVerifyCmd.MarkFlagRequired("amount1")
	sourceDepositVerifyCmd.Flags().StringVar(&sourceDepositVerifyAmount2, "amount2", "", "second micro deposit amount (required)")
	sourceDepositVerifyCmd.MarkFlagRequired("amount2")
	sourceDepositVerifyCmd.Flags().StringVar(&sourceDepositVerifyCurrency, "currency", "USD", "micro deposit currency")

	sourceUpdateCmd.Flags().StringVar(&sourceUpdateAccountNumber, "account-number", "", "bank account number")
	sourceUpdateCmd.Flags().StringVar(&sourceUpdateBankAccountType, "account-type", "", "bank account type")
	sourceUpdateCmd.Flags().StringVar(&sourceUpdateName, "name", "", "bank account nickname")
	sourceUpdateCmd.Flags().StringVar(&sourceUpdateRoutingNumber, "routing-number", "", "bank routing number")
}
