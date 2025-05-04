package private

import (
	"fmt"
	"strings"

	"github.com/DanWlker/asnbcli/internal/helpers"
)

func Logout(authorization string) error {
	defer helpers.HttpClient.CloseIdleConnections()
	resp, err := helpers.HttpClient.Post(
		"https://myasnb-api-v4.myasnb.com.my/v2/logout",
		"application/json",
		strings.NewReader(""),
	)
	if err != nil {
		return fmt.Errorf("Logout: %w", err)
	}
	helpers.PrintResponseHelper(resp)

	return nil
}

