package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

type FpxBanks struct {
	Id            int    `json:"ID"`
	DateCreated   string `json:"DATE_CREATED"`
	FpxBankCode   string `json:"FPX_BANK_CODE"`
	FpxBankName   string `json:"FPX_BANK_NAME"`
	LastUpdated   string `json:"LAST_UPDATED"`
	FullName      string `json:"FULL_NAME"`
	IsActive      int    `json:"IS_ACTIVE"`
	Status        any    `json:"STATUS"`
	ImageUrl      string `json:"IMAGE_URL"`
	ImageUrlBanks string `json:"IMAGE_URL_BANKS"`
	BankCode      string `json:"BANK_CODE"`
}

type getAllFpxBanksResult struct {
	Data []FpxBanks `json:"data"`
}

func GetAllFpxBanks(authorization string) ([]FpxBanks, error) {
	req, err := http.NewRequest(http.MethodGet, "https://myasnb-api-v4.myasnb.com.my/v2/subscription/fpxbanks", nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Add("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer resp.Body.Close()

	respByte, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, fmt.Errorf("httputil.DumpResponse: %w", err)
	}
	fmt.Println("")
	fmt.Println(string(respByte))
	fmt.Println("")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	result := getAllFpxBanksResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result.Data, nil
}
