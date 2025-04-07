/*
Copyright © 2025 DanWlker danielhee2@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/DanWlker/asnbcli/internal"
	"github.com/spf13/cobra"
)

const (
	usernameF = "username"
	passwordF = "password"
	repeatF   = "repeat"
	offsetF   = "offset"
	fundsF    = "funds"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asnbcli",
	Short: "A cli app to simplify buying asnb funds so you (hopefully) don't have to wake up at 2 am ",
	Long:  `A cli app that returns a payment link, simplifies the process of buying asnb funds. It also has repeat functionality with offset for each loop`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString(usernameF)
		if err != nil || username == "" {
			username, err = internal.InputHelper("Username: ", false)
			if err != nil {
				panic(fmt.Errorf("unable to get username: %v", err))
			}
		}

		password, err := cmd.Flags().GetString(passwordF)
		if err != nil || password == "" {
			password, err = internal.InputHelper("Password: ", true)
			if err != nil {
				panic(fmt.Errorf("unable to get password: %v", err))
			}
		}

		funds, err := cmd.Flags().GetStringSlice(fundsF)
		if err != nil {
			panic(fmt.Errorf("unable to get funds: %v", err))
		}

		if len(funds) == 0 {
			fmt.Println("No funds specified")
			return
		}

		if err := internal.StartExecution(internal.NewEntryParams(
			internal.WithUsername(username),
			internal.WithPassword(password),
			internal.WithFunds(funds),
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.asnb_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP(usernameF, "u", "", "Username for your account")
	rootCmd.Flags().StringP(passwordF, "p", "", "Password for your account")
	rootCmd.Flags().IntP(repeatF, "r", 0, "Amount of times to repeat if fail")
	rootCmd.Flags().IntP(offsetF, "o", 5, "Offset time to wait before repeating in seconds")
	rootCmd.Flags().StringSliceP(fundsF, "f", []string{internal.ASM1, internal.ASM2, internal.ASM3}, "The funds to try, defaults to ASM1, ASM2 and ASM3")
}
