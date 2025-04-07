package private

type LoginResult struct {
	Token                string `json:"token"` // TODO: This is technically a jwt token
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

func Login(username, password string) (LoginResult, error) {
	result := LoginResult{}

	return result, nil
}
