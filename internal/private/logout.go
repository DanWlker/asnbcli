package private

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DanWlker/asnbcli/internal/helpers"
)

func Logout(authorization string) error {
	defer helpers.HttpClient.CloseIdleConnections()

	req, err := http.NewRequest(
		http.MethodPost,
		"https://myasnb-api-v4.myasnb.com.my/v2/logout",
		strings.NewReader(""),
	)
	if err != nil {
		return fmt.Errorf("Logout: %w", err)
	}
	req.Header.Add("Authorization", authorization)
	helpers.PrintRequestHelper(req)

	resp, err := helpers.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Logout: %w", err)
	}
	helpers.PrintResponseHelper(resp)

	return nil
}
