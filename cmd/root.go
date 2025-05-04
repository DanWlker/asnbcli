/*
Copyright © 2025 DanWlker danielhee2@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DanWlker/asnbcli/internal"
	"github.com/DanWlker/asnbcli/internal/helpers"
	"github.com/spf13/cobra"
)

const (
	usernameF      = "username"
	passwordF      = "password"
	repeatF        = "repeat"
	offsetF        = "offset"
	fundsF         = "funds"
	tokenF         = "token"
	writeTokenF    = "write-token"
	amountF        = "amount"
	paymentMethodF = "payment-method"
	fpxBankF       = "fpx-bank"
	debugF         = "debug"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asnbcli",
	Short: "A cli app to simplify buying asnb funds so you (hopefully) don't have to wake up at 2 am ",
	Long:  `A cli app that returns a payment link, simplifies the process of buying asnb funds. It also has repeat functionality with offset for each loop`,
	Run: func(cmd *cobra.Command, args []string) {
		// Username
		username, err := cmd.Flags().GetString(usernameF)
		if err != nil || username == "" {
			username, err = helpers.InputHelper("Username: ", false)
			if err != nil {
				panic(fmt.Errorf("unable to get username: %v", err))
			}
		}

		// Password
		password, err := cmd.Flags().GetString(passwordF)
		if err != nil || password == "" {
			password, err = helpers.InputHelper("Password: ", true)
			if err != nil {
				panic(fmt.Errorf("unable to get password: %v", err))
			}
		}

		// Amount
		amount, err := cmd.Flags().GetString(amountF)
		if err != nil || amount == "" {
			amount, err = helpers.InputHelper("Amount in RM (ex. 20): ", false)
			if err != nil {
				panic(fmt.Errorf("unable to get amount: %v", err))
			}

		}

		// Payment Method
		paymentMethod, err := cmd.Flags().GetString(paymentMethodF)
		if err != nil || paymentMethod == "" {
			for i, method := range internal.AllPaymentMethods {
				fmt.Printf("%v: %v\n", i, method)
			}

			selectedIdxStr, err := helpers.InputHelper("Select payment method (ex. 1): ", false)
			if err != nil {
				panic(fmt.Errorf("unable to get payment method: %v", err))
			}
			selectedIdx, err := strconv.ParseInt(selectedIdxStr, 10, 64)
			if err != nil {
				panic(fmt.Errorf("payment method: strconv.ParseInt: %w", err))
			}
			if selectedIdx >= int64(len(internal.AllPaymentMethods)) || selectedIdx < 0 {
				panic(fmt.Errorf("invalid range, must be between 0 and %v: %v", len(internal.AllPaymentMethods), selectedIdx))
			}

			paymentMethod = internal.AllPaymentMethods[int(selectedIdx)]
		}

		var withPaymentMethod internal.Option
		switch paymentMethod {
		case internal.Tngd:
			withPaymentMethod = internal.WithTngd()
		case internal.Boost:
			withPaymentMethod = internal.WithBoost()
		case internal.Fpx:
			fpxBank, err := cmd.Flags().GetString(fpxBankF)
			if err != nil {
				fmt.Println("error when getting fpx bank, will prompt again later")
			}
			if fpxBank == "" {
				fmt.Println("bank for fpx payment not specified, will prompt again later")
			}
			withPaymentMethod = internal.WithFpx(fpxBank)
		default:
			panic(fmt.Errorf("unknown payment method: %v", paymentMethod))
		}

		// Fund list
		funds, err := cmd.Flags().GetStringSlice(fundsF)
		if err != nil {
			panic(fmt.Errorf("unable to get funds: %v", err))
		}
		if len(funds) == 0 {
			fmt.Println("No funds specified")
			return
		}
		for i, fund := range funds {
			fundPostfix, ok := internal.FundToUrlPostfix[fund]
			if !ok {
				fmt.Printf("unknown fund %v, will still try to buy\n", fund)
				continue
			}
			funds[i] = fundPostfix
		}

		// Debug
		debug, err := cmd.Flags().GetBool(debugF)
		if err != nil {
			fmt.Printf("unable to get debug flag, not printing debug logs, please report this error: %v\n", err)
		}

		// Start Execution
		if err := internal.StartExecution(internal.NewEntryParams(
			internal.WithUsername(username),
			internal.WithPassword(password),
			internal.WithFunds(funds),
			internal.WithAmount(amount),
			internal.WithDebug(debug),
			withPaymentMethod,
		)); err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP(usernameF, "u", "", "Username for your account")
	rootCmd.Flags().StringP(passwordF, "p", "", "Password for your account")
	rootCmd.Flags().StringP(amountF, "a", "", "Amount to buy")
	rootCmd.Flags().StringP(paymentMethodF, "m", "", "Payment method to use, accepted values: "+strings.Join(internal.AllPaymentMethods, ","))
	rootCmd.Flags().String(fpxBankF, "", "Fpx bank to use (ex. HLB0224)")
	// rootCmd.Flags().IntP(repeatF, "r", 0, "Amount of times to repeat if fail")
	// rootCmd.Flags().IntP(offsetF, "o", 5, "Offset time to wait before repeating in seconds")
	rootCmd.Flags().StringSliceP(fundsF, "f", []string{internal.Asm1, internal.Asm2, internal.Asm3}, "The funds to try, if the fund provided is not in the list of accepted values, it will still try to buy the provided fund")
	// rootCmd.Flags().StringP(tokenF, "t", "", "The bearer token to use for requests")
	// rootCmd.Flags().Bool(writeTokenF, false, "Write the bearer token to a file to reuse in future calls")
	rootCmd.Flags().Bool(debugF, false, "Debug requests")
}
