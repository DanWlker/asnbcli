package internal

import (
	"fmt"
	"strconv"

	"github.com/DanWlker/asnbcli/internal/private"
)

type entryParams struct {
	username string
	password string
	funds    []string
}

type Option func(*entryParams)

func NewEntryParams(options ...Option) entryParams {
	result := &entryParams{}
	for _, o := range options {
		o(result)
	}
	return *result
}

func WithUsername(username string) Option {
	return func(e *entryParams) {
		e.username = username
	}
}

func WithPassword(password string) Option {
	return func(e *entryParams) {
		e.password = password
	}
}

func WithFunds(funds []string) Option {
	return func(e *entryParams) {
		e.funds = funds
	}
}

func StartExecution(params entryParams) error {
	fmt.Println("=============")
	fmt.Println("Logging in...")
	fmt.Println("=============")
	loginResult, err := private.Login(params.username, params.password)
	if err != nil {
		return fmt.Errorf("private.Login: %w", err)
	}

	fmt.Println("=============")
	fmt.Println("Getting all fpx banks...")
	fmt.Println("=============")
	fpxBanks, err := private.GetAllFpxBanks(fmt.Sprintf("Bearer %v", loginResult.Token))
	if err != nil {
		return fmt.Errorf("private.GetAllFpxBanks: %w", err)
	}

	// TODO: Allow tng, boost as well
	fmt.Println("=============")
	fmt.Println("Select bank to use...")
	fmt.Println("=============")
	var selectedBank private.FpxBanks
	for range 3 {
		for i, fpxBank := range fpxBanks {
			fmt.Printf("%v: %v\n", i, fpxBank.FullName)
		}

		selectedId, err := InputHelper("Enter number (ex. 1): ", false)
		if err != nil {
			fmt.Println(fmt.Errorf("select bank: InputHelper: %w", err))
			continue
		}
		selectedIdInt, err := strconv.ParseInt(selectedId, 10, 64)
		if err != nil {
			fmt.Println(fmt.Errorf("select bank: strconv.ParseInt: %w", err))
			continue
		}
		if selectedIdInt >= int64(len(fpxBanks)) || selectedIdInt < 0 {
			fmt.Println(fmt.Errorf("invalid range, must be between 0 and %v: %v", len(fpxBanks), selectedId))
			continue
		}

		selectedBank = fpxBanks[int(selectedIdInt)]
		break
	}

	fmt.Println("=============")
	fmt.Println("Fund buy details...")
	fmt.Println("=============")
	amount, err := InputHelper("Amount (ex. 500): ", false)
	if err != nil {
		return fmt.Errorf("InputHelper: %w", err)
	}

	for _, fund := range params.funds {
		err = private.BuyFundWithFpx(fmt.Sprintf("Bearer %v", loginResult.Token), amount, fund, loginResult.Uhid, selectedBank.FpxBankCode)
		if err != nil {
			return fmt.Errorf("BuyFund: %w", err)
		}
	}

	return nil
}
