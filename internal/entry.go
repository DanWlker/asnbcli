package internal

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/DanWlker/asnbcli/internal/helpers"
	"github.com/DanWlker/asnbcli/internal/private"
)

type entryParams struct {
	username      string
	password      string
	funds         []string
	amount        string
	paymentMethod string
	fpxBank       string
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

func WithAmount(amount string) Option {
	return func(e *entryParams) {
		e.amount = amount
	}
}

func WithBoost() Option {
	return func(e *entryParams) {
		e.paymentMethod = Boost
	}
}

func WithTngd() Option {
	return func(e *entryParams) {
		e.paymentMethod = Tngd
	}
}

func WithFpx(fpxBank string) Option {
	return func(e *entryParams) {
		e.paymentMethod = Fpx
		e.fpxBank = fpxBank
	}
}

func WithFunds(funds []string) Option {
	return func(e *entryParams) {
		e.funds = funds
	}
}

func StartExecution(params entryParams) error {
	// Login
	helpers.VerboseLogger.Println("Logging in...")
	loginResult, err := private.Login(params.username, params.password)
	if err != nil {
		return fmt.Errorf("private.Login: %w", err)
	}
	if loginResult.Token == "" {
		return errors.New("token is empty, exiting")
	}
	formattedToken := "Bearer " + loginResult.Token

	// Start buying
	helpers.VerboseLogger.Println("Start buying...")
	var (
		errLists     []error
		paymentLinks []string
	)
	switch params.paymentMethod {
	case Tngd:
		for _, fund := range params.funds {
			link, err := private.BuyFundWithTng(
				formattedToken,
				params.amount,
				fund,
				loginResult.Uhid,
			)
			if err != nil {
				errLists = append(errLists, err)
				continue
			}

			paymentLinks = append(paymentLinks, link)
		}
	case Boost:
		for _, fund := range params.funds {
			link, err := private.BuyFundWithBoost(
				formattedToken,
				params.amount,
				fund,
				loginResult.Uhid,
			)
			if err != nil {
				errLists = append(errLists, err)
				continue
			}

			paymentLinks = append(paymentLinks, link)
		}
	case Fpx:
		if params.fpxBank == "" {
			helpers.VerboseLogger.Println("Getting all fpx banks...")
			fpxBanks, err := private.GetAllFpxBanks(formattedToken)
			if err != nil {
				return fmt.Errorf("private.GetAllFpxBanks: %w", err)
			}

			helpers.StdErrLogger.Println("Select bank to use...")
			for i, fpxBank := range fpxBanks {
				helpers.StdErrLogger.Printf("%v: %v\n", i, fpxBank.FullName)
			}

			selectedId, err := helpers.InputHelper("Enter number (ex. 1): ", false)
			if err != nil {
				return fmt.Errorf("select bank: InputHelper: %w", err)
			}
			selectedIdInt, err := strconv.ParseInt(selectedId, 10, 64)
			if err != nil {
				return fmt.Errorf("select bank: strconv.ParseInt: %w", err)
			}
			if selectedIdInt >= int64(len(fpxBanks)) || selectedIdInt < 0 {
				return fmt.Errorf("invalid range, must be between 0 and %v: %v", len(fpxBanks), selectedId)
			}

			params.fpxBank = fpxBanks[int(selectedIdInt)].FpxBankCode
		}

		for _, fund := range params.funds {
			link, err := private.BuyFundWithFpx(
				formattedToken,
				params.amount,
				fund,
				loginResult.Uhid,
				params.fpxBank,
			)
			if err != nil {
				errLists = append(errLists, err)
				continue
			}

			paymentLinks = append(paymentLinks, link)
		}
	default:
		panic(fmt.Errorf("unknown payment method: %v", params.paymentMethod))
	}

	// Print results
	for _, err := range errLists {
		helpers.StdErrLogger.Println(err)
	}
	for _, link := range paymentLinks {
		fmt.Println(link)
	}

	return nil
}
