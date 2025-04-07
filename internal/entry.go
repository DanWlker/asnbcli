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
	loginResult, err := private.Login(params.username, params.password)
	if err != nil {
		return err
	}

	fmt.Println(loginResult)

	return nil
}
