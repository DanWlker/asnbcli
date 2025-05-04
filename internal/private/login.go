package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func Login(username, password string, debug bool) (*LoginResult, error) {
	key := "IVObMPuj2rnhBGaOa0YN5TopZsLnGdPqOffDxXJVOSKFqonbIGE0a7xGgGUdq2_TSzpOBGqe9hMI5nV0AtGMV5ieBo_uwIFWzgL19LI16khI_xdvrMvsBN_i4Ay91qd1zt3lCXdp9-Df16mxIeqVbIn6E1hzQxc_QOWwwS-SkEI"

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
		return nil, fmt.Errorf("t.SignedString: %w", err)
	}

	reqBody := loginRequest{Jwt: signedString}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := http.Post(
		"https://myasnb-api-v4.myasnb.com.my/v2/login",
		"application/json",
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, fmt.Errorf("http.Post: %w", err)
	}
	defer resp.Body.Close()
	// TODO: Remove this
	helpers.PrintResponseHelper(resp, debug)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("resp.Body.Read: %w", err)
	}

	result := LoginResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &result, nil
}
