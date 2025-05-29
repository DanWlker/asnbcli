package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/DanWlker/asnbcli/internal/helpers"
)

type fpxData struct {
	FpxBuyerAccNo      string `json:"fpx_buyerAccNo"`
	FpxBuyerBankBranch string `json:"fpx_buyerBankBranch"`
	FpxBuyerBankId     string `json:"fpx_buyerBankId"`
	FpxBuyerEmail      string `json:"fpx_buyerEmail"`
	FpxBuyerIban       string `json:"fpx_buyerIban"`
	FpxBuyerId         string `json:"fpx_buyerId"`
	FpxBuyerName       string `json:"fpx_buyerName"`
	FpxCheckSum        string `json:"fpx_checkSum"`
	FpxMakerName       string `json:"fpx_makerName"`
	FpxMsgToken        string `json:"fpx_msgToken"`
	FpxMsgType         string `json:"fpx_msgType"`
	FpxProductDesc     string `json:"fpx_productDesc"`
	FpxSellerBankCode  string `json:"fpx_sellerBankCode"`
	FpxSellerExId      string `json:"fpx_sellerExId"`
	FpxSellerExOrderNo string `json:"fpx_sellerExOrderNo"`
	FpxSellerId        string `json:"fpx_sellerId"`
	FpxSellerOrderNo   string `json:"fpx_sellerOrderNo"`
	FpxSellerTxnTime   string `json:"fpx_sellerTxnTime"`
	FpxTxnAmount       string `json:"fpx_txnAmount"`
	FpxTxnCurrency     string `json:"fpx_txnCurrency"`
	FpxVersion         string `json:"fpx_version"`
}

type fundData struct {
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
	FpxUrl  string  `json:"FPX_URL"`
	FpxData fpxData `json:"FPX_DATA"`

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

type buyFundResponse struct {
	Data fundData `json:"data"`
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

func BuyFundWithFpx(authorization, amount, fund, unitHolderId, fpxBankId string) (string, error) {
	resp, err := buyFund(authorization, amount, fund, unitHolderId, "", fpxBankId)
	if err != nil {
		return "", fmt.Errorf("BuyFundWithFpx: %w", err)
	}

	if err := checkBuyFundError(resp); err != nil {
		return "", fmt.Errorf("BuyFundWithFpx: %w", err)
	}

	if resp.Data.FpxUrl == "" {
		return "", fmt.Errorf("BuyFundWithFpx: fpx url is empty")
	}

	if resp.Data.FpxData == (fpxData{}) {
		return "", fmt.Errorf("BuyFundWithFpx: fpx data is empty")
	}

	// let L = encodeURI(
	//     "fpx_buyerAccNo=" +
	//       y +
	//       "&fpx_buyerBankId=" +
	//       j +
	//       "&fpx_buyerEmail=" +
	//       S +
	//       "&fpx_msgToken=" +
	//       T +
	//       "&fpx_msgType=" +
	//       k +
	//       "&fpx_productDesc=" +
	//       R +
	//       "&fpx_sellerBankCode=" +
	//       O +
	//       "&fpx_sellerExId=" +
	//       M +
	//       "&fpx_sellerExOrderNo=" +
	//       I +
	//       "&fpx_sellerId=" +
	//       P +
	//       "&fpx_sellerOrderNo=" +
	//       F +
	//       "&fpx_sellerTxnTime=" +
	//       U +
	//       "&fpx_txnAmount=" +
	//       X +
	//       "&fpx_txnCurrency=" +
	//       D +
	//       "&fpx_version=" +
	//       B +
	//       "&fpx_checkSum=" +
	//       _,
	//   ),
	queryParams := url.Values{}
	queryParams.Add("fpx_buyerAccNo", resp.Data.FpxData.FpxBuyerAccNo)
	queryParams.Add("fpx_buyerBankId", resp.Data.FpxData.FpxBuyerBankId)
	queryParams.Add("fpx_buyerEmail", resp.Data.FpxData.FpxBuyerEmail)
	queryParams.Add("fpx_msgToken", resp.Data.FpxData.FpxMsgToken)
	queryParams.Add("fpx_msgType", resp.Data.FpxData.FpxMsgType)
	queryParams.Add("fpx_productDesc", resp.Data.FpxData.FpxProductDesc)
	queryParams.Add("fpx_sellerBankCode", resp.Data.FpxData.FpxSellerBankCode)
	queryParams.Add("fpx_sellerExId", resp.Data.FpxData.FpxSellerExId)
	queryParams.Add("fpx_sellerExOrderNo", resp.Data.FpxData.FpxSellerExOrderNo)
	queryParams.Add("fpx_sellerId", resp.Data.FpxData.FpxSellerId)
	queryParams.Add("fpx_sellerOrderNo", resp.Data.FpxData.FpxSellerOrderNo)
	queryParams.Add("fpx_sellerTxnTime", resp.Data.FpxData.FpxSellerTxnTime)
	queryParams.Add("fpx_txnAmount", resp.Data.FpxData.FpxTxnAmount)
	queryParams.Add("fpx_txnCurrency", resp.Data.FpxData.FpxTxnCurrency)
	queryParams.Add("fpx_version", resp.Data.FpxData.FpxVersion)
	queryParams.Add("fpx_checkSum", resp.Data.FpxData.FpxCheckSum)

	return resp.Data.FpxUrl + "?" + queryParams.Encode(), nil
}

func BuyFundWithTng(authorization, amount, fund, unitHolderId string) (string, error) {
	resp, err := buyFund(authorization, amount, fund, unitHolderId, "TNGD", "")
	if err != nil {
		return "", fmt.Errorf("BuyFundWithTng: %w", err)
	}

	if err := checkBuyFundError(resp); err != nil {
		return "", fmt.Errorf("BuyFundWithTng: %w", err)
	}

	if resp.Data.TngdUrl.TngdBody.TngdResponse.Response.Body.CheckoutUrl == "" {
		return "", fmt.Errorf("BuyFundWithTng: tngd checkout url is empty")
	}

	return resp.Data.TngdUrl.TngdBody.TngdResponse.Response.Body.CheckoutUrl, nil
}

func BuyFundWithBoost(authorization, amount, fund, unitHolderId string) (string, error) {
	resp, err := buyFund(authorization, amount, fund, unitHolderId, "boost", "")
	if err != nil {
		return "", fmt.Errorf("BuyFundWithBoost: %w", err)
	}

	if err := checkBuyFundError(resp); err != nil {
		return "", fmt.Errorf("BuyFundWithBoost: %w", err)
	}

	if resp.Data.Boost.BoostQrResponse.CheckoutUri == "" {
		return "", fmt.Errorf("BuyFundWithBoost: boost checkout uri is empty")
	}

	return resp.Data.Boost.BoostQrResponse.CheckoutUri, nil
}

func buyFund(authorization, amount, fund, unitHolderId, paymentProcessor, fpxBankId string) (buyFundResponse, error) {
	res := buyFundResponse{}

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
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Type", "application/json")
	helpers.PrintRequestHelper(req)

	defer helpers.HttpClient.CloseIdleConnections()
	resp, err := helpers.HttpClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("Do: %w", err)
	}
	defer resp.Body.Close()
	helpers.PrintResponseHelper(resp)

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}

func checkBuyFundError(resp buyFundResponse) error {
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
