package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/DanWlker/asnbcli/internal/helpers"
	"github.com/golang-jwt/jwt/v5"
)

type LoginResult struct {
	Token                string `json:"token"`
	Uhid                 string `json:"uhid"`
	FirstTimeLogin       int    `json:"first_time_login"`
	PreferredLanguage    string `json:"preferred_language"`
	InitialFund          bool   `json:"initial_fund"`
	StepCode             any    `json:"stepCode"`
	IsBankdetailsUpdated int    `json:"is_bankdetails_updated"`
	SecureImage          bool   `json:"Secure_Image"`
	UserID               int    `json:"user_id"`
	IsRegisterDevice     bool   `json:"isRegisterDevice"`
}

type loginRequest struct {
	Jwt string `json:"jwt"`
}

func Login(username, password string) (LoginResult, error) {
	key := "IVObMPuj2rnhBGaOa0YN5TopZsLnGdPqOffDxXJVOSKFqonbIGE0a7xGgGUdq2_TSzpOBGqe9hMI5nV0AtGMV5ieBo_uwIFWzgL19LI16khI_xdvrMvsBN_i4Ay91qd1zt3lCXdp9-Df16mxIeqVbIn6E1hzQxc_QOWwwS-SkEI"
	result := LoginResult{}

	currDate := time.Now()
	currDateUnix := currDate.Unix()
	twentyMonthsLater := currDate.AddDate(1, 8, 0)
	twentyMonthsLaterUnix := twentyMonthsLater.Unix()

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"applicationID": "MYASNB",
			"aud":           "www.myasnb.com.my",
			"deviceId":      "",
			"exp":           twentyMonthsLaterUnix,
			"iat":           currDateUnix,
			"password":      password,
			"sub":           username,
		},
	)
	signedString, err := t.SignedString([]byte(key))
	if err != nil {
		return result, fmt.Errorf("t.SignedString: %w", err)
	}

	reqBody := loginRequest{Jwt: signedString}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return result, fmt.Errorf("json.Marshal: %w", err)
	}

	defer helpers.HttpClient.CloseIdleConnections()
	resp, err := helpers.HttpClient.Post(
		"https://myasnb-api-v4.myasnb.com.my/v2/login",
		"application/json",
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return result, fmt.Errorf("http.Post: %w", err)
	}
	defer resp.Body.Close()
	helpers.PrintResponseHelper(resp)

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}
