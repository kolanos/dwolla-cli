package cmd

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/kolanos/dwolla-v2-go"
	"github.com/spf13/cobra"
)

var accountPaymentListCorrelationID string
var accountPaymentListLimit int
var accountPaymentListOffset int

var accountSourceCreateAccountNumber string
var accountSourceCreateBankAccountType string
var accountSourceCreateChannels string
var accountSourceCreateName string
var accountSourceCreateRoutingNumber string

var accountSourceListRemoved bool

var accountTransferListCorrelationID string
var accountTransferListEndAmount string
var accountTransferListEndDate string
var accountTransferListLimit int
var accountTransferListOffset int
var accountTransferListSearch string
var accountTransferListStartAmount string
var accountTransferListStartDate string
var accountTransferListStatus string

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account Management",
}

var accountPaymentCmd = &cobra.Command{
	Use:   "payment",
	Short: "Account Mass Payment Management",
}

var accountPaymentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Mass Payments",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		params := &url.Values{}

		if accountPaymentListCorrelationID != "" {
			params.Add("correlationId", accountPaymentListCorrelationID)
		}

		if accountPaymentListLimit != 0 {
			params.Add("limit", strconv.Itoa(accountPaymentListLimit))
		}

		if accountPaymentListOffset != 0 {
			params.Add("offset", strconv.Itoa(accountPaymentListOffset))
		}

		act, err := client.Account.Retrieve()
		if err != nil {
			renderError(err)
		}

		res, err := act.ListMassPayments(params)
		if err != nil {
			renderError(err)
		}

		data := make([][]string, len(res.Embedded["mass-payments"]))

		for i, d := range res.Embedded["mass-payments"] {
			data[i] = []string{d.ID, string(d.Status), d.Total.String(), d.TotalFees.String(), d.Created, d.CorrelationID}
		}

		header := []string{"ID", "Status", "Total Amount", "Total Fees Amount", "created", "Correlation ID"}
		footer := []string{"", "", "", "", "Total", strconv.Itoa(res.Total)}

		renderCollection(data, header, footer)
	},
}

var accountRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve Account Details",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		res, err := client.Account.Retrieve()
		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}
var accountSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Account Funding Source Management",
}

var accountSourceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Account Funding Source",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		act, err := client.Account.Retrieve()
		if err != nil {
			renderError(err)
		}

		var bankAccountType dwolla.FundingSourceBankAccountType
		switch accountSourceCreateBankAccountType {
		case "checking":
			bankAccountType = dwolla.FundingSourceBankAccountTypeChecking
		case "savings":
			bankAccountType = dwolla.FundingSourceBankAccountTypeSavings
		default:
			renderError(fmt.Errorf("Invalid bank account type: %s", accountSourceCreateBankAccountType))
		}

		var channels []string

		if accountSourceCreateChannels != "" {
			channels = strings.Split(accountSourceCreateChannels, ",")
		}

		res, err := act.CreateFundingSource(&dwolla.FundingSourceRequest{
			RoutingNumber:   accountSourceCreateRoutingNumber,
			AccountNumber:   accountSourceCreateAccountNumber,
			BankAccountType: bankAccountType,
			Name:            accountSourceCreateName,
			Channels:        channels,
		})

		if err != nil {
			renderError(err)
		}

		renderResource(res)
	},
}

var accountSourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Funding Sources",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		act, err := client.Account.Retrieve()
		if err != nil {
			renderError(err)
		}

		res, err := act.ListFundingSources(accountSourceListRemoved)
		if err != nil {
			renderError(err)
		}

		data := make([][]string, len(res.Embedded["funding-sources"]))

		for i, d := range res.Embedded["funding-sources"] {
			data[i] = []string{d.ID, string(d.Status), string(d.Type), string(d.BankAccountType), d.Name, d.BankName, strconv.FormatBool(d.Removed), d.Created}
		}

		header := []string{"ID", "Status", "Type", "Account Type", "Name", "Bank Name", "Created"}
		footer := []string{"", "", "", "", "", "Total", strconv.Itoa(res.Total)}

		renderCollection(data, header, footer)
	},
}

var accountTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Account Transfer Management",
}

var accountTransferListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Account Transfers",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		params := &url.Values{}

		if accountTransferListCorrelationID != "" {
			params.Add("correlationId", accountTransferListCorrelationID)
		}

		if accountTransferListEndAmount != "" {
			params.Add("endAmount", accountTransferListEndAmount)
		}

		if accountTransferListEndDate != "" {
			params.Add("endDate", accountTransferListEndDate)
		}

		if accountTransferListLimit != 0 {
			params.Add("limit", strconv.Itoa(accountTransferListLimit))
		}

		if accountTransferListOffset != 0 {
			params.Add("offset", strconv.Itoa(accountTransferListOffset))
		}

		if accountTransferListSearch != "" {
			params.Add("search", accountTransferListSearch)
		}

		if accountTransferListStartAmount != "" {
			params.Add("startAmount", accountTransferListStartAmount)
		}

		if accountTransferListStartDate != "" {
			params.Add("startDate", accountTransferListStartDate)
		}

		if accountTransferListStatus != "" {
			params.Add("status", accountTransferListStatus)
		}

		act, err := client.Account.Retrieve()
		if err != nil {
			renderError(err)
		}

		res, err := act.ListTransfers(params)
		if err != nil {
			renderError(err)
		}

		data := make([][]string, len(res.Embedded["transfers"]))

		for i, d := range res.Embedded["transfers"] {
			data[i] = []string{d.ID, string(d.Status), d.Amount.String(), d.Created, d.CorrelationID}
		}

		header := []string{"ID", "Status", "Amount", "Created", "Correlation ID"}
		footer := []string{"", "", "", "Total", strconv.Itoa(res.Total)}

		renderCollection(data, header, footer)
	},
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

	accountPaymentListCmd.Flags().StringVar(&accountPaymentListCorrelationID, "correlation-id", "", "filter by correlation id")
	accountPaymentListCmd.Flags().IntVar(&accountPaymentListLimit, "limit", 25, "number of results to return")
	accountPaymentListCmd.Flags().IntVar(&accountPaymentListOffset, "offset", 0, "number of results to skip")

	accountSourceCreateCmd.Flags().StringVar(&accountSourceCreateAccountNumber, "account-number", "", "bank account number (required)")
	accountSourceCreateCmd.MarkFlagRequired("account-number")
	accountSourceCreateCmd.Flags().StringVar(&accountSourceCreateBankAccountType, "account-type", "", "bank account type (required)")
	accountSourceCreateCmd.MarkFlagRequired("account-type")
	accountSourceCreateCmd.Flags().StringVar(&accountSourceCreateChannels, "channels", "", "bank account channels (comma separated)")
	accountSourceCreateCmd.Flags().StringVar(&accountSourceCreateName, "name", "", "bank account nickname (required)")
	accountSourceCreateCmd.MarkFlagRequired("name")
	accountSourceCreateCmd.Flags().StringVar(&accountSourceCreateRoutingNumber, "routing-number", "", "bank routing number (required)")
	accountSourceCreateCmd.MarkFlagRequired("routing-number")

	accountSourceListCmd.Flags().BoolVarP(&accountSourceListRemoved, "removed", "r", false, "include removed funding sources")

	accountTransferListCmd.Flags().StringVar(&accountTransferListCorrelationID, "correlation-id", "", "filter by correlation id")
	accountTransferListCmd.Flags().StringVar(&accountTransferListEndAmount, "end-amount", "", "filter by end amount")
	accountTransferListCmd.Flags().StringVar(&accountTransferListEndDate, "end-date", "", "filter by end date (YYYY-MM-DD)")
	accountTransferListCmd.Flags().IntVar(&accountTransferListLimit, "limit", 25, "number of results to return")
	accountTransferListCmd.Flags().IntVar(&accountTransferListOffset, "offset", 0, "number of results to skip")
	accountTransferListCmd.Flags().StringVar(&accountTransferListSearch, "search", "", "filter by search string")
	accountTransferListCmd.Flags().StringVar(&accountTransferListStartAmount, "start-amount", "", "filter by start amount")
	accountTransferListCmd.Flags().StringVar(&accountTransferListStartDate, "start-date", "", "filter by start date (YYYY-MM-DD)")
	accountTransferListCmd.Flags().StringVar(&accountTransferListStatus, "status", "", "filter by status")
}
