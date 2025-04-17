package internal

import (
	"fmt"

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

	fmt.Println("=============")
	fmt.Println("Select bank to use:")
	fmt.Println("=============")
	for i, fpxBank := range fpxBanks {
		fmt.Printf("%v: %v\n", i, fpxBank.FullName)
	}

	return nil
}
