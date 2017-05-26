package neverbounce

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	Expires          int    `json:"expires"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type VerifyEmailResponse struct {
	Success          bool   `json:"success"`
	Result           int    `json:"result"`
	ResultDetails    int    `json:"result_details"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorMsg         string `json:"error_msg"`
	Msg              string `json:"msg"`
}
