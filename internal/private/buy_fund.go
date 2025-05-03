package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FundData struct {
	AgentCode                            string `json:"AGENTCODE"`
	AmountApplied                        int    `json:"AMOUNTAPPLIED"`
	BankTxnReferenceNumber               string `json:"BANKTXNREFERENCENUMBER"`
	BranchCode                           string `json:"BRANCHCODE"`
	ChannelType                          string `json:"CHANNELTYPE"`
	DeviceOwner                          string `json:"DEVICEOWNER"`
	FundId                               string `json:"FUNDID"`
	IdentificationNumber                 string `json:"IDENTIFICATIONNUMBER"`
	IdentificationType                   string `json:"IDENTIFICATIONTYPE"`
	InvestmentToleranceScore             string `json:"INVESTMENTTOLERANCESCORE"`
	PaymentType                          string `json:"PAYMENTTYPE"`
	RejectCode                           string `json:"REJECTCODE"`
	RejectReason                         string `json:"REJECTREASON"`
	RequestOrIdentification              string `json:"REQUESTORIDENTIFICATION"`
	RiskProfile                          string `json:"RISKPROFILE"`
	SourceRefNo                          string `json:"SOURCEREFNO"`
	SuitabilityAssessmentCode            string `json:"SUITABILITYASSESSMENTCODE"`
	SuitabilityAssessmentStatus          string `json:"SUITABILITYASSESSMENTSTATUS"`
	TransactionDate                      string `json:"TRANSACTIONDATE"`
	TransactionStatus                    string `json:"TRANSACTIONSTATUS"`
	TransactionTime                      string `json:"TRANSACTIONTIME"`
	UnitHolderId                         string `json:"UNITHOLDERID"`
	FirstName                            string `json:"FIRSTNAME"`
	CustomerIcNumber                     string `json:"CUSTOMERICNUMBER"`
	CustomerName                         string `json:"CUSTOMERNAME"`
	BankCustPhoneNumber                  string `json:"BANKCUSTPHONENUMBER"`
	BankAccountNumber                    string `json:"BANKACCOUNTNUMBER"`
	BankBranchCode                       string `json:"BANKBRANCHCODE"`
	ChequeNumber                         string `json:"CHEQUENUMBER"`
	ChequeDate                           string `json:"CHEQUEDATE"`
	GuardianId                           string `json:"GUARDIANID"`
	GuardianIcType                       string `json:"GUARDIANICTYPE"`
	GuardianIcNumber                     string `json:"GUARDIANICNUMBER"`
	PolicyNumber                         string `json:"POLICYNUMBER"`
	EpfNumber                            string `json:"EPFNUMBER"`
	SubPaymentType                       string `json:"SUBPAYMENTTYPE"`
	EwGateway                            string `json:"EWGATEWAY"`
	ThirdPartyInvestment                 string `json:"THIRDPARTYINVESTMENT"`
	ThirdPartyName                       string `json:"THIRDPARTYNAME"`
	ThirdPartyIcType                     string `json:"THIRDPARTYICTYPE"`
	ThirdPartyIcNumber                   string `json:"THIRDPARTYICNUMBER"`
	ThirdPartyRelationship               string `json:"THIRDPARTYRELATIONSHIP"`
	ReasonForTransfer                    string `json:"REASONFORTRANSFER"`
	SourceOfFund                         string `json:"SOURCEOFFUND"`
	OtherSourceOfFund                    string `json:"OTHERSOURCEOFFUND"`
	FunderName                           string `json:"FUNDERNAME"`
	SourceOfWealth                       string `json:"SOURCEOFWEALTH"`
	OtherSourceOfWealth                  string `json:"OTHERSOURCEOFWEALTH"`
	UnitsAlloted                         string `json:"UNITSALLOTED"`
	TransactionNumber                    string `json:"TRANSACTIONNUMBER"`
	FundPrice                            string `json:"FUNDPRICE"`
	FeePercentage                        string `json:"FEEPERCENTAGE"`
	SalesCharge                          string `json:"SALESCHARGE"`
	GstAmount                            string `json:"GSTAMOUNT"`
	TaxInvoiceNumber                     string `json:"TAXINVOICENUMBER"`
	InvestmentToleranceLevel             string `json:"INVESTMENTTOLERANCELEVEL"`
	ThirdPartyResidentialAddressLine1    string `json:"THIRDPARTYRESIDENTIALADDRESSLINE1"`
	ThirdPartyResidentialAddressLine2    string `json:"THIRDPARTYRESIDENTIALADDRESSLINE2"`
	ThirdPartyResidentialAddressCity     string `json:"THIRDPARTYRESIDENTIALADDRESSCITY"`
	ThirdPartyResidentialAddressState    string `json:"THIRDPARTYRESIDENTIALADDRESSSTATE"`
	ThirdPartyResidentialAddressPostCode string `json:"THIRDPARTYRESIDENTIALADDRESSPOSTCODE"`
	ThirdPartyResidentialAddressCountry  string `json:"THIRDPARTYRESIDENTIALADDRESSCOUNTRY"`
	ThirdPartyMobileNumber               string `json:"THIRDPARTYMOBILENUMBER"`
	FeeType                              string `json:"FEETYPE"`
	LeadGenerator                        string `json:"LEADGENERATOR"`
	FinancialExecutive                   string `json:"FINANCIALEXECUTIVE"`

	// FPX
	FpxUrl string `json:"FPX_URL"`

	// TNGD
	// TNGD_URL.TNGD_BODY.tngDResponse.response.body.checkoutUrl,
	TngdUrl struct {
		TngdBody struct {
			TngdResponse struct {
				Response struct {
					Body struct {
						CheckoutUrl string `json:"checkoutUrl"`
					} `json:"body"`
				} `json:"response"`
			} `json:"tngDResponse"`
		} `json:"TNGD_BODY"`
	} `json:"TNGD_URL"`

	// Boost
	// BOOST_URL.boostQRResponse.checkoutURI
	Boost struct {
		BoostQrResponse struct {
			CheckoutUri string `json:"checkoutURI"`
		} `json:"boostQRResponse"`
	} `json:"BOOST_URL"`
}

type BuyFundResponse struct {
	Data FundData `json:"data"`
}

type buyFundRequest struct {
	Amount               string `json:"amount"`
	TxnCode              string `json:"TXN_CODE"`
	CriteriaConfirm      bool   `json:"CRITERIACONFIRM"`
	ProspectusAgree      bool   `json:"PROSPECTUSAGREE"`
	RiskSrcOfFund        string `json:"riskSrcOfFund"`
	RiskSrcOfWealth      string `json:"riskSrcOfWealth"`
	RiskOtherSrcOfFund   string `json:"riskOtherSrcOfFund"`
	RiskOtherSrcOfWealth string `json:"riskOtherSrcOfWealth"`
	ReferralCode         string `json:"referral_code"`

	// Different things that should be omitted if empty
	FpxBankID        string `json:"fpxBankId,omitempty"`
	PaymentProcessor string `json:"paymentProcessor,omitempty"`
}

func BuyFundWithFpx(authorization, amount, fund, unitHolderId, fpxBankId string) error {
	fmt.Println("buying with fpx")
	// return buyFund(authorization, amount, fund, unitHolderId, "", fpxBankId)
	return nil
}

func BuyFundWithTng(authorization, amount, fund, unitHolderId string) error {
	resp, err := buyFund(authorization, amount, fund, unitHolderId, "TNGD", "")
	if err != nil {
		return err
	}

	if err := checkBuyFundError(resp); err != nil {
		return err
	}

	if resp.Data.TngdUrl.TngdBody.TngdResponse.Response.Body.CheckoutUrl == "" {
		return fmt.Errorf("tngd checkout url is empty")
	}

	fmt.Println(resp.Data.TngdUrl.TngdBody.TngdResponse.Response.Body.CheckoutUrl)

	return nil
}

func BuyFundWithBoost(authorization, amount, fund, unitHolderId string) error {
	resp, err := buyFund(authorization, amount, fund, unitHolderId, "boost", "")
	if err != nil {
		return err
	}

	if err := checkBuyFundError(resp); err != nil {
		return err
	}

	if resp.Data.Boost.BoostQrResponse.CheckoutUri == "" {
		return fmt.Errorf("boost checkout uri is empty")
	}

	fmt.Println(resp.Data.Boost.BoostQrResponse.CheckoutUri)

	return nil
}

func buyFund(authorization, amount, fund, unitHolderId, paymentProcessor, fpxBankId string) (BuyFundResponse, error) {
	res := BuyFundResponse{}

	reqBody := buyFundRequest{
		Amount:               amount,
		TxnCode:              "I01",
		CriteriaConfirm:      true,
		ProspectusAgree:      true,
		RiskSrcOfFund:        "",
		RiskSrcOfWealth:      "",
		RiskOtherSrcOfFund:   "",
		RiskOtherSrcOfWealth: "",
		ReferralCode:         "",
		// Extras
		FpxBankID:        fpxBankId,
		PaymentProcessor: paymentProcessor,
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return res, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://myasnb-api-v4.myasnb.com.my/v2/subscription/provisional/%v/%v", unitHolderId, fund),
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return res, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// TODO: Remove this
	PrintRequestHelper(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer resp.Body.Close()

	// TODO: Remove this
	PrintResponseHelper(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("resp.Body.Read: %w", err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func checkBuyFundError(resp BuyFundResponse) error {
	if resp.Data.RejectCode != "" {
		return fmt.Errorf(
			"buy fund tng returned with reject code: %v, reject reason: %v, transaction status: %v",
			resp.Data.RejectCode,
			resp.Data.RejectReason,
			resp.Data.TransactionStatus,
		)
	}

	return nil
}
