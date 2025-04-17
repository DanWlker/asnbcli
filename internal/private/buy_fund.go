package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type buyFundRequest struct {
	Amount               string `json:"amount"`
	FpxBankID            string `json:"fpxBankId"`
	TxnCode              string `json:"TXN_CODE"`
	CriteriaConfirm      bool   `json:"CRITERIACONFIRM"`
	ProspectusAgree      bool   `json:"PROSPECTUSAGREE"`
	RiskSrcOfFund        string `json:"riskSrcOfFund"`
	RiskSrcOfWealth      string `json:"riskSrcOfWealth"`
	RiskOtherSrcOfFund   string `json:"riskOtherSrcOfFund"`
	RiskOtherSrcOfWealth string `json:"riskOtherSrcOfWealth"`
	ReferralCode         string `json:"referral_code"`
}

func BuyFund(authorization, amount, fund, unitHolderId, fpxBankId string) error {
	reqBody := buyFundRequest{
		Amount:               amount,
		FpxBankID:            fpxBankId,
		TxnCode:              "I01",
		CriteriaConfirm:      true,
		ProspectusAgree:      true,
		RiskSrcOfFund:        "",
		RiskSrcOfWealth:      "",
		RiskOtherSrcOfFund:   "",
		RiskOtherSrcOfWealth: "",
		ReferralCode:         "",
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	url := fmt.Sprintf("https://myasnb-api-v4.myasnb.com.my/v2/subscription/provisional/%v/%v", unitHolderId, fund)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Add("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do: %w", err)
	}

	// TODO: Remove this
	PrintResponseHelper(resp)

	return nil
}
